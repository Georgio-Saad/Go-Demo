package repositories

import "todogorest/models"

type VerificationCodeRepository interface {
	Create() (models.VerificationCode, error)
	Update() (models.VerificationCode, error)
	Delete() error
	FindByUserId() (models.VerificationCode, error)
}
