package services

import (
	"net/http"
	"strconv"
	"todogorest/data/request"
	"todogorest/data/response"
	"todogorest/helpers"
	"todogorest/repositories"
	"todogorest/validations"
)

type VerificationCodeServicesImpl struct {
	VerificationCodeRepository repositories.VerificationCodeRepository
}

// Create implements VerificationCodeServices.
func (s *VerificationCodeServicesImpl) Create(verificationCodeRequest request.VerificationCodeRequest) response.Response {
	createVerificationCodeErr := validations.ValidateRequest(verificationCodeRequest)

	if createVerificationCodeErr != nil {
		return response.Response{StatusCode: http.StatusUnprocessableEntity, Message: createVerificationCodeErr.Error(), Code: helpers.UnprocessableEntity}
	}

	verificationCode, err := s.VerificationCodeRepository.Create(verificationCodeRequest)

	if err != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Message: err.Error(), Code: helpers.BadRequest}
	}

	return response.Response{StatusCode: http.StatusCreated, Message: "Success", Code: helpers.Success, Data: verificationCode}
}

// Delete implements VerificationCodeServices.
func (v *VerificationCodeServicesImpl) Delete(verificationCodeId string) response.Response {
	id, err := strconv.Atoi(verificationCodeId)

	if err != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Message: "Invalid verification code id", Code: helpers.BadRequest}
	}

	resErr := v.VerificationCodeRepository.Delete(id)

	if resErr != nil {
		return response.Response{StatusCode: http.StatusNotFound, Message: "Verification code not found", Code: helpers.NotFound}
	}

	return response.Response{StatusCode: http.StatusOK, Message: "Successfully deleted verification code", Code: helpers.Success}
}

// DeleteByUserId implements VerificationCodeServices.
func (v *VerificationCodeServicesImpl) DeleteByUserId(userId string) response.Response {
	id, err := strconv.Atoi(userId)

	if err != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Message: "Invalid user id", Code: helpers.BadRequest}
	}

	resErr := v.VerificationCodeRepository.DeleteByUserId(id)

	if resErr != nil {
		return response.Response{StatusCode: http.StatusNotFound, Message: "Verification code not found", Code: helpers.NotFound}
	}

	return response.Response{StatusCode: http.StatusOK, Message: "Successfully deleted verification code", Code: helpers.Success}
}

// FindById implements VerificationCodeServices.
func (v *VerificationCodeServicesImpl) FindById(verificationCodeId string) response.Response {
	id, err := strconv.Atoi(verificationCodeId)

	if err != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Message: "Invalid verification code id", Code: helpers.BadRequest}
	}

	verificationCode, resErr := v.VerificationCodeRepository.FindById(id)

	if resErr != nil {
		return response.Response{StatusCode: http.StatusNotFound, Message: "Verification code not found", Code: helpers.NotFound}
	}

	return response.Response{StatusCode: http.StatusOK, Message: "Successfully fetched verification code", Code: helpers.Success, Data: verificationCode}
}

// FindByUserId implements VerificationCodeServices.
func (v *VerificationCodeServicesImpl) FindByUserId(userId string) response.Response {
	id, err := strconv.Atoi(userId)

	if err != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Message: "Invalid verification code id", Code: helpers.BadRequest}
	}

	verificationCode, resErr := v.VerificationCodeRepository.FindByUserId(id)

	if resErr != nil {
		return response.Response{StatusCode: http.StatusNotFound, Message: "Verification code not found", Code: helpers.NotFound}
	}

	return response.Response{StatusCode: http.StatusOK, Message: "Successfully fetched verification code", Code: helpers.Success, Data: verificationCode}
}

// Update implements VerificationCodeServices.
func (v *VerificationCodeServicesImpl) Update(updateRequest request.VerificationCodeRequest) response.Response {
	updateVerificationCodeRequestErr := validations.ValidateRequest(updateRequest)

	if updateVerificationCodeRequestErr != nil {
		return response.Response{StatusCode: http.StatusUnprocessableEntity, Message: updateVerificationCodeRequestErr.Error(), Code: helpers.UnprocessableEntity}
	}

	verificationCode, resErr := v.VerificationCodeRepository.Update(updateRequest)

	if resErr != nil {
		return response.Response{StatusCode: http.StatusNotFound, Message: resErr.Error(), Code: helpers.NotFound}
	}

	return response.Response{StatusCode: http.StatusOK, Message: "Successfully updated verification code", Code: helpers.Success, Data: verificationCode}
}

func NewVerificationCodeServicesImpl(verificationCodeRepository repositories.VerificationCodeRepository) VerificationCodeServices {
	return &VerificationCodeServicesImpl{
		VerificationCodeRepository: verificationCodeRepository,
	}
}
