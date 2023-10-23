package services

import (
	"net/http"
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

	privateKey, pkErr := jwt.LoadPrivateKeyRSA("../rsapss_private_key.pem")

	if pkErr != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Message: pkErr.Error(), Code: helpers.BadRequest}
	}

	accessToken, jwtErr := jwt.SignEncrypted(
		jwt.RS256,
		privateKey,
		func(plainPayload []byte) ([]byte, error) {
			return []byte{}, nil
		},
		userCreated,
	)

	refreshToken, refErr := jwt.SignEncrypted(
		jwt.RS256,
		privateKey,
		func(plainPayload []byte) ([]byte, error) {
			return []byte{}, nil
		},
		userCreated,
	)

	if jwtErr != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Message: jwtErr.Error(), Code: helpers.BadRequest}
	}

	if refErr != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Message: jwtErr.Error(), Code: helpers.BadRequest}
	}

	return response.Response{StatusCode: http.StatusCreated, Message: "Successfully created user", Code: helpers.Success, Data: response.AuthResponse{User: userCreated, AccessToken: string(accessToken), RefreshToken: string(refreshToken), ExpiresAt: ""}}
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
