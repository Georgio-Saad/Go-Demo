package repositories

import (
	"errors"
	"todogorest/data/request"
	"todogorest/data/response"
	"todogorest/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

// FindByUsername implements UserRepository.
func (u *UserRepositoryImpl) FindUser(username string, password string) (models.User, error) {
	var user models.User

	result := u.Db.First(&user, username)

	if result.Error != nil {
		return models.User{}, errors.New("Please check your login credentials")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return models.User{}, errors.New("Please check your login credentials")
	}

	return user, nil
}

// Create implements UserRepository.
func (u UserRepositoryImpl) Create(userDetails request.CreateUserRequest) (models.User, error) {
	var user = models.User{}

	userAlreadyFound := u.Db.First(&user, userDetails.Username)

	if userAlreadyFound.Error == nil {
		return models.User{}, errors.New("A user with this username already exists")
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(userDetails.Password), 12)

	user.Username = userDetails.Username
	user.Locale = userDetails.Locale
	user.Password = string(password)
	user.Email = userDetails.Email
	user.DateOfBirth = userDetails.DateOfBirth
	user.CountryCode = userDetails.CountryCode
	user.PhoneNumber = userDetails.PhoneNumber

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
func (UserRepositoryImpl) FindById(userId int) (models.User, error) {
	panic("unimplemented")
}

// Update implements UserRepository.
func (UserRepositoryImpl) Update() (models.User, error) {
	panic("unimplemented")
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: db}
}
