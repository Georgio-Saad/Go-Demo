package services

import (
	"net/http"
	"strconv"
	"time"
	"todogorest/constants"
	"todogorest/data/request"
	"todogorest/data/response"
	"todogorest/helpers"
	"todogorest/repositories"
	"todogorest/validations"

	"github.com/kataras/jwt"
)

type UserServicesImpl struct {
	UserRepository repositories.UserRepository
}

// Refresh implements UserServices.
func (s *UserServicesImpl) Refresh(token string) response.Response {
	if len(token) == 0 {
		return response.Response{StatusCode: http.StatusUnprocessableEntity, Message: "Token is required", Code: helpers.UnprocessableEntity}
	}

	_, accessDec, err := jwt.GCM(constants.AccessEncKey, nil)

	if err != nil {
		return response.Response{StatusCode: http.StatusUnauthorized, Code: helpers.Unauthenticated, Message: err.Error()}
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

	userCreated, err := s.UserRepository.Create(user)

	if err != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Message: err.Error(), Code: helpers.BadRequest}
	}

	accessToken, jwtErr := helpers.GenerateAccessToken(userCreated)
	refreshToken, refErr := helpers.GenerateRefreshToken(userCreated)

	if jwtErr != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Message: jwtErr.Error(), Code: helpers.BadRequest}
	}

	if refErr != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Message: refErr.Error(), Code: helpers.BadRequest}
	}

	return response.Response{
		StatusCode: http.StatusCreated,
		Message:    "Successfully created user",
		Code:       helpers.Success,
		Data: response.AuthResponse{
			User:         userCreated,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			ExpiresAt:    time.Now().Add(15 * time.Minute).Local().String(),
		},
	}
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
	id, err := strconv.Atoi(userId)

	if err != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Message: err.Error(), Code: helpers.BadRequest}
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

	user, err := s.UserRepository.FindUser(credentials.Username, credentials.Password)

	if err != nil {
		return response.Response{StatusCode: http.StatusUnauthorized, Message: err.Error(), Code: helpers.Unauthorized}
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

func NewUserServicesImpl(userRepository repositories.UserRepository) UserServices {
	return &UserServicesImpl{UserRepository: userRepository}
}
