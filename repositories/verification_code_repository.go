package repositories

import (
	"todogorest/data/request"
	"todogorest/models"
)

type VerificationCodeRepository interface {
	Create(request.VerificationCodeRequest) (models.VerificationCode, error)
	Update(request.VerificationCodeRequest) (models.VerificationCode, error)
	Delete(verificationCodeId int) error
	DeleteByUserId(userId int) error
	FindByUserId(userId int) (models.VerificationCode, error)
	FindById(verificationCodeId int) (models.VerificationCode, error)
}
