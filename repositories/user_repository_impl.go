package repositories

import (
	"errors"
	"log"
	"todogorest/data/request"
	"todogorest/data/response"
	"todogorest/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

// RemoveProfilePicture implements UserRepository.
func (u *UserRepositoryImpl) RemoveProfilePicture(userId int) (models.User, error) {
	var user models.User

	result := u.Db.Model(&models.User{}).Where("id = ?", userId).First(&user)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	user.ProfilePicture = nil

	u.Db.Save(&user)

	return user, nil
}

// SetProfilePicture implements UserRepository.
func (u *UserRepositoryImpl) SetProfilePicture(userId int, profilePictureUrl string) (models.User, error) {
	var user models.User

	result := u.Db.Model(&models.User{}).Where("id = ?", userId).First(&user)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	user.ProfilePicture = &profilePictureUrl

	u.Db.Save(&user)

	return user, nil
}

// Save implements UserRepository.
func (u *UserRepositoryImpl) Save(user *models.User) {
	u.Db.Save(&user)
}

// FindByUsername implements UserRepository.
func (u *UserRepositoryImpl) FindUser(username string, password string) (models.User, error) {
	var user models.User

	result := u.Db.First(&user, models.User{Username: username})

	if result.Error != nil {
		return models.User{}, errors.New("Please check your login credentials user")
	}

	log.Default().Println(result.RowsAffected, user)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return models.User{}, errors.New("Please check your login credentials pass")
	}

	if !user.Verified {
		return models.User{}, errors.New("Please verify your account before signing in")
	}

	return user, nil
}

// Create implements UserRepository.
func (u UserRepositoryImpl) Create(userDetails request.CreateUserRequest) (models.User, error) {
	var userAlreadyExists models.User

	userResult := u.Db.Model(&models.User{}).Where("username = ?", userDetails.Username).Or("email = ?", userDetails.Email).First(&userAlreadyExists)

	if userResult.Error == nil {
		return models.User{}, errors.New("User already exists")
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(userDetails.Password), 12)
	var user = models.User{Username: userDetails.Username,
		Password:       string(password),
		Email:          userDetails.Email,
		DateOfBirth:    userDetails.DateOfBirth,
		CountryCode:    userDetails.CountryCode,
		PhoneNumber:    userDetails.PhoneNumber,
		Verified:       false,
		Role:           userDetails.Role,
		ProfilePicture: nil,
	}

	result := u.Db.Create(&user)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

// Delete implements UserRepository.
func (UserRepositoryImpl) Delete(userId int) error {
	panic("unimplemented")
}

// FindAll implements UserRepository.
func (UserRepositoryImpl) FindAll(request.PaginationRequest) (users response.PaginationResponse[models.User], err error) {
	panic("unimplemented")
}

// FindById implements UserRepository.
func (u UserRepositoryImpl) FindById(userId int) (models.User, error) {
	var user models.User

	userResult := u.Db.First(&user, userId)

	if userResult.Error != nil {
		return models.User{}, userResult.Error
	}

	return user, nil
}

// Update implements UserRepository.
func (UserRepositoryImpl) Update() (models.User, error) {
	panic("unimplemented")
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: db}
}
