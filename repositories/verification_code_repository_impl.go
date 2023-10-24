package repositories

import (
	"errors"
	"todogorest/data/request"
	"todogorest/models"

	"gorm.io/gorm"
)

type VerificationCodeRepositoryImpl struct {
	Db *gorm.DB
}

// DeleteByUserId implements VerificationCodeRepository.
func (v *VerificationCodeRepositoryImpl) DeleteByUserId(userId int) error {
	var verificationCode models.VerificationCode

	result := v.Db.Model(&models.VerificationCode{}).Where("user_id = ?", userId).Delete(&verificationCode)

	if result != nil {
		return result.Error
	}

	return nil
}

// Create implements VerificationCodeRepository.
func (v *VerificationCodeRepositoryImpl) Create(verificationCodeDetails request.VerificationCodeRequest) (models.VerificationCode, error) {
	verificatinCode := models.VerificationCode{AlreadyUsed: verificationCodeDetails.AlreadyUsed, UserID: verificationCodeDetails.UserID, VerificationCode: verificationCodeDetails.VerificationCode}
	var verificationCodeAlreadyExists *models.VerificationCode

	var user models.User

	userResult := v.Db.Model(&models.User{}).Where("id = ?", verificationCodeDetails.UserID).Find(&user)

	alreadyExistsResult := v.Db.Model(&models.VerificationCode{}).Where("user_id = ?", verificationCodeDetails.UserID).First(&verificationCodeAlreadyExists)

	if userResult.RowsAffected == 0 {
		return models.VerificationCode{}, errors.New("User doesn't exist")
	}

	if alreadyExistsResult.RowsAffected > 0 {
		return models.VerificationCode{}, errors.New("A verification code already exists for this user")
	}

	result := v.Db.Create(&verificatinCode)

	if result.Error != nil {
		return models.VerificationCode{}, result.Error
	}

	return verificatinCode, nil
}

// Delete implements VerificationCodeRepository.
func (v *VerificationCodeRepositoryImpl) Delete(verificationCodeId int) error {
	var verificationCode models.VerificationCode

	result := v.Db.Delete(&verificationCode, verificationCodeId)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// FindById implements VerificationCodeRepository.
func (v *VerificationCodeRepositoryImpl) FindById(verificationCodeId int) (models.VerificationCode, error) {
	var verificationCode models.VerificationCode

	result := v.Db.Model(&models.VerificationCode{}).Where("id = ?", verificationCodeId).First(&verificationCode)

	if result.Error != nil {
		return models.VerificationCode{}, result.Error
	}

	return verificationCode, nil
}

// FindByUserId implements VerificationCodeRepository.
func (v *VerificationCodeRepositoryImpl) FindByUserId(userId int) (models.VerificationCode, error) {
	var verificationCode models.VerificationCode

	result := v.Db.Model(&models.VerificationCode{}).Where("user_id = ?", userId).First(&verificationCode)

	if result.Error != nil {
		return models.VerificationCode{}, result.Error
	}

	return verificationCode, nil
}

// Update implements VerificationCodeRepository.
func (v *VerificationCodeRepositoryImpl) Update(verificationCodeDetails request.VerificationCodeRequest) (models.VerificationCode, error) {
	var verificationCode models.VerificationCode

	result := v.Db.Model(&models.VerificationCode{}).Where("user_id = ?", verificationCodeDetails.UserID).Find(&verificationCode)

	if result.RowsAffected == 0 {
		return models.VerificationCode{}, errors.New("no record found")
	}

	verificationCode.AlreadyUsed = verificationCodeDetails.AlreadyUsed
	verificationCode.VerificationCode = verificationCodeDetails.VerificationCode

	v.Db.Save(&verificationCode)

	return verificationCode, nil
}

func NewVerificationCodeRepository(db *gorm.DB) VerificationCodeRepository {
	return &VerificationCodeRepositoryImpl{Db: db}
}
