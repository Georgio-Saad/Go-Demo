package services

import (
	"net/http"
	"time"
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

// Create implements UserServices.
func (s *UserServicesImpl) Create(user request.CreateUserRequest) response.Response {
	validationErr := validations.ValidateRequest(user)

	if validationErr != nil {
		return response.Response{StatusCode: http.StatusUnprocessableEntity, Message: validationErr.Error(), Code: helpers.UnprocessableEntity}
	}

	userCreated, err := s.UserRepository.Create(user)

	if err != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Message: err.Error(), Code: helpers.BadRequest}
	}

	accessTokenPrivateKey, pkErr := jwt.LoadPrivateKeyRSA("./rsapss_private_key.pem")
	refreshTokenPrivateKey, rfPkErr := jwt.LoadPrivateKeyRSA("./rsa_private_key.pem")

	if pkErr != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Message: pkErr.Error(), Code: helpers.BadRequest}
	}

	if rfPkErr != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Message: rfPkErr.Error(), Code: helpers.BadRequest}
	}

	accessTokenClaims := &helpers.JWTClaims{
		User:   userCreated,
		Claims: jwt.Claims{Expiry: time.Now().Add(15 * time.Minute).Unix()},
	}

	refreshTokenClaims := &helpers.JWTClaims{
		User:   userCreated,
		Claims: jwt.Claims{Expiry: time.Now().Add(2 * 7 * 24 * time.Hour).Unix()},
	}

	accessToken, jwtErr := jwt.SignEncrypted(
		jwt.RS256,
		accessTokenPrivateKey,
		func(plainPayload []byte) ([]byte, error) {
			return []byte{}, nil
		},
		accessTokenClaims,
	)

	refreshToken, refErr := jwt.SignEncrypted(
		jwt.RS256,
		refreshTokenPrivateKey,
		func(plainPayload []byte) ([]byte, error) {
			return []byte{}, nil
		},
		refreshTokenClaims,
	)

	if jwtErr != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Message: jwtErr.Error(), Code: helpers.BadRequest}
	}

	if refErr != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Message: jwtErr.Error(), Code: helpers.BadRequest}
	}

	return response.Response{
		StatusCode: http.StatusCreated,
		Message:    "Successfully created user",
		Code:       helpers.Success,
		Data: response.AuthResponse{
			User:         userCreated,
			AccessToken:  string(accessToken),
			RefreshToken: string(refreshToken),
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
func (*UserServicesImpl) FindById(userId string) response.Response {
	panic("unimplemented")
}

// FindUser implements UserServices.
func (*UserServicesImpl) FindUser(request.SigninUserRequest) response.Response {
	panic("unimplemented")
}

// Update implements UserServices.
func (*UserServicesImpl) Update() response.Response {
	panic("unimplemented")
}

func NewUserServicesImpl(userRepository repositories.UserRepository) UserServices {
	return &UserServicesImpl{UserRepository: userRepository}
}
