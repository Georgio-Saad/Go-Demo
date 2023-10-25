package services

import (
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"
	"todogorest/constants"
	"todogorest/data/request"
	"todogorest/data/response"
	"todogorest/helpers"
	"todogorest/repositories"
	"todogorest/validations"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/kataras/jwt"
)

type UserServicesImpl struct {
	UserRepository             repositories.UserRepository
	VerificationCodeRepository repositories.VerificationCodeRepository
}

// RemoveProfilePicture implements UserServices.
func (s *UserServicesImpl) RemoveProfilePicture(userId string) response.Response {
	id, idErr := strconv.Atoi(userId)

	if idErr != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Message: idErr.Error(), Code: helpers.BadRequest}
	}

	user, userErr := s.UserRepository.RemoveProfilePicture(id)

	if userErr != nil {
		return response.Response{StatusCode: http.StatusNotFound, Message: userErr.Error(), Code: helpers.NotFound}
	}

	return response.Response{StatusCode: http.StatusOK, Message: "Successfully removed profile picture", Code: helpers.Success, Data: user}
}

// UploadProfilePicture implements UserServices.
func (s *UserServicesImpl) UploadProfilePicture(userId string, profilePictureFile multipart.File, profilePictureHeader *multipart.FileHeader, sess *session.Session) response.Response {
	id, idErr := strconv.Atoi(userId)

	if idErr != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Message: idErr.Error(), Code: helpers.BadRequest}
	}
	bucket := os.Getenv("AWS_S3_BUCKET_NAME")

	uploader := s3manager.NewUploader(sess)

	fileName := userId + "-" + profilePictureHeader.Filename

	upload, uploadErr := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		ACL:    aws.String("public-read"),
		Key:    &fileName,
		Body:   profilePictureFile,
	})

	if uploadErr != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Message: uploadErr.Error(), Code: helpers.BadRequest}
	}

	user, userErr := s.UserRepository.SetProfilePicture(id, upload.Location)

	if userErr != nil {
		return response.Response{StatusCode: http.StatusNotFound, Message: userErr.Error(), Code: helpers.NotFound}
	}

	return response.Response{StatusCode: http.StatusOK, Message: "Successfully uploaded profile picture", Code: helpers.Success, Data: user}
}

// ResendVerification implements UserServices.
func (s *UserServicesImpl) ResendVerification(userId string) response.Response {
	id, idErr := strconv.Atoi(userId)

	if idErr != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Message: idErr.Error(), Code: helpers.BadRequest}
	}

	user, userErr := s.UserRepository.FindById(id)

	if userErr != nil {
		return response.Response{StatusCode: http.StatusNotFound, Message: userErr.Error(), Code: helpers.NotFound}
	}

	verificationCode, verErr := s.VerificationCodeRepository.Update(request.VerificationCodeRequest{UserID: int(user.ID), VerificationCode: helpers.GenerateVerificationCode(), AlreadyUsed: false})

	if verErr != nil {
		return response.Response{StatusCode: http.StatusNotFound, Message: verErr.Error(), Code: helpers.NotFound}
	}
	log.Default().Println(verificationCode)

	return response.Response{StatusCode: http.StatusOK, Message: "Successfully resent verification code", Code: helpers.Success}
}

// Verify implements UserServices.
func (s *UserServicesImpl) Verify(verifyData request.VerifyUserRequest) response.Response {
	id, idErr := strconv.Atoi(verifyData.UserID)

	if idErr != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Message: idErr.Error(), Code: helpers.BadRequest}
	}

	user, userErr := s.UserRepository.FindById(id)

	if userErr != nil {
		return response.Response{StatusCode: http.StatusNotFound, Message: userErr.Error(), Code: helpers.NotFound}
	}

	if user.Verified && verifyData.Login {
		return response.Response{StatusCode: http.StatusForbidden, Message: "User already verified", Code: helpers.Forbidden}
	}

	verificationCode, verErr := s.VerificationCodeRepository.FindByUserId(int(user.ID))

	if verErr != nil {
		return response.Response{StatusCode: http.StatusNotFound, Message: verErr.Error(), Code: helpers.NotFound}
	}

	if verificationCode.AlreadyUsed {
		return response.Response{StatusCode: http.StatusConflict, Message: "Invalid verification code", Code: helpers.InvalidData}
	}

	isMatching := verifyData.VerificationCode == verificationCode.VerificationCode

	if isMatching {
		user.Verified = true
		verificationCode.AlreadyUsed = true

		s.UserRepository.Save(&user)
		s.VerificationCodeRepository.Save(&verificationCode)

		if verifyData.Login {
			accessToken, accErr := helpers.GenerateAccessToken(user)
			refreshToken, refErr := helpers.GenerateRefreshToken(user)

			if accErr != nil {
				return response.Response{StatusCode: http.StatusConflict, Message: accErr.Error(), Code: helpers.InvalidData}
			}

			if refErr != nil {
				return response.Response{StatusCode: http.StatusConflict, Message: refErr.Error(), Code: helpers.InvalidData}
			}

			return response.Response{
				StatusCode: http.StatusOK,
				Message:    "Successfully verified user",
				Code:       helpers.Success,
				Data: response.AuthResponse{
					User:         user,
					AccessToken:  accessToken,
					RefreshToken: refreshToken,
					ExpiresAt:    time.Now().Add(15 * time.Minute).Local().String(),
				},
			}
		}

		return response.Response{StatusCode: http.StatusOK, Message: "Successfully verified user", Code: helpers.Success, Data: user}
	}

	return response.Response{StatusCode: http.StatusConflict, Message: "Invalid verification code", Code: helpers.InvalidData}
}

// Refresh implements UserServices.
func (s *UserServicesImpl) Refresh(token string) response.Response {
	if len(token) == 0 {
		return response.Response{StatusCode: http.StatusUnprocessableEntity, Message: "Token is required", Code: helpers.UnprocessableEntity}
	}

	_, accessDec, decErr := jwt.GCM(constants.AccessEncKey, nil)

	if decErr != nil {
		return response.Response{StatusCode: http.StatusUnauthorized, Code: helpers.Unauthenticated, Message: decErr.Error()}
	}

	decToken, jwtErr := jwt.VerifyEncrypted(jwt.HS256, constants.RefreshSignKey, accessDec, []byte(token))

	if jwtErr != nil {
		return response.Response{StatusCode: http.StatusUnauthorized, Code: helpers.Unauthenticated, Message: jwtErr.Error()}
	}

	var claims helpers.JWTClaims

	claimsErr := decToken.Claims(&claims)

	if claimsErr != nil {
		return response.Response{StatusCode: http.StatusUnauthorized, Code: helpers.Unauthenticated, Message: jwtErr.Error()}
	}

	if claims.GrantType != constants.RefreshToken {
		return response.Response{StatusCode: http.StatusUnauthorized, Code: helpers.Unauthenticated, Message: "Unauthorized"}
	}

	user, resErr := s.UserRepository.FindById(int(claims.User.ID))

	if resErr != nil {
		return response.Response{StatusCode: http.StatusNotFound, Message: resErr.Error(), Code: helpers.NotFound}
	}

	accessToken, accErr := helpers.GenerateAccessToken(user)
	refreshToken, refErr := helpers.GenerateRefreshToken(user)

	if accErr != nil {
		return response.Response{StatusCode: http.StatusConflict, Message: accErr.Error(), Code: helpers.InvalidData}
	}

	if refErr != nil {
		return response.Response{StatusCode: http.StatusConflict, Message: refErr.Error(), Code: helpers.InvalidData}
	}

	return response.Response{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Code:       helpers.Success,
		Data: response.AuthResponse{
			User:         user,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			ExpiresAt:    time.Now().Add(15 * time.Minute).Local().String(),
		},
	}

}

// Create implements UserServices.
func (s *UserServicesImpl) Create(user request.CreateUserRequest) response.Response {
	validationErr := validations.ValidateRequest(user)

	if validationErr != nil {
		return response.Response{StatusCode: http.StatusUnprocessableEntity, Message: validationErr.Error(), Code: helpers.UnprocessableEntity, Data: validationErr}
	}

	userCreated, createErr := s.UserRepository.Create(user)

	if createErr != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Message: createErr.Error(), Code: helpers.BadRequest}
	}

	_, verErr := s.VerificationCodeRepository.Create(request.VerificationCodeRequest{AlreadyUsed: false, VerificationCode: helpers.GenerateVerificationCode(), UserID: int(userCreated.ID)})

	if verErr != nil {
		return response.Response{StatusCode: http.StatusUnprocessableEntity, Message: verErr.Error(), Code: helpers.UnprocessableEntity}
	}

	return response.Response{StatusCode: http.StatusCreated, Message: "Signed up successfully, please check your email to verify your account", Code: helpers.Success}
}

// Delete implements UserServices.
func (*UserServicesImpl) Delete(userId string) response.Response {
	panic("unimplemented")
}

// FindAll implements UserServices.
func (*UserServicesImpl) FindAll(request.PaginationRequest) response.Response {
	panic("unimplemented")
}

// FindById implements UserServices.
func (s *UserServicesImpl) FindById(userId string) response.Response {
	id, strConvErr := strconv.Atoi(userId)

	if strConvErr != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Message: strConvErr.Error(), Code: helpers.BadRequest}
	}

	user, resErr := s.UserRepository.FindById(id)

	if resErr != nil {
		return response.Response{StatusCode: http.StatusNotFound, Message: resErr.Error(), Code: helpers.NotFound}
	}

	return response.Response{StatusCode: http.StatusOK, Message: "Successfully fetched user", Code: helpers.Success, Data: user}
}

// FindUser implements UserServices.
func (s *UserServicesImpl) FindUser(credentials request.SigninUserRequest) response.Response {
	validationErr := validations.ValidateRequest(credentials)

	if validationErr != nil {
		return response.Response{StatusCode: http.StatusUnprocessableEntity, Message: validationErr.Error(), Code: helpers.UnprocessableEntity}
	}

	user, findErr := s.UserRepository.FindUser(credentials.Username, credentials.Password)

	if findErr != nil {
		return response.Response{StatusCode: http.StatusUnauthorized, Message: findErr.Error(), Code: helpers.Unauthorized}
	}

	accessToken, jwtErr := helpers.GenerateAccessToken(user)
	refreshToken, refErr := helpers.GenerateRefreshToken(user)

	if jwtErr != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Message: jwtErr.Error(), Code: helpers.BadRequest}
	}

	if refErr != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Message: refErr.Error(), Code: helpers.BadRequest}
	}

	return response.Response{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Code:       helpers.Success,
		Data: response.AuthResponse{
			User:         user,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			ExpiresAt:    time.Now().Add(15 * time.Minute).Local().String(),
		},
	}
}

// Update implements UserServices.
func (*UserServicesImpl) Update() response.Response {
	panic("unimplemented")
}

func NewUserServicesImpl(userRepository repositories.UserRepository, verificationCodeRepository repositories.VerificationCodeRepository) UserServices {
	return &UserServicesImpl{UserRepository: userRepository, VerificationCodeRepository: verificationCodeRepository}
}
