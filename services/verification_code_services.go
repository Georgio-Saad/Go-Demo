package services

import (
	"todogorest/data/request"
	"todogorest/data/response"
)

type VerificationCodeServices interface {
	Create(request.VerificationCodeRequest) response.Response
	Update(request.VerificationCodeRequest) response.Response
	FindByUserId(userId string) response.Response
	FindById(verificationCodeId string) response.Response
	DeleteByUserId(userId string) response.Response
	Delete(verificationCodeId string) response.Response
}
