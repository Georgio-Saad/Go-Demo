package repositories

import (
	"todogorest/data/request"
	"todogorest/data/response"
	"todogorest/models"
)

type UserRepository interface {
	Create(request.CreateUserRequest) (models.User, error)
	Update() (models.User, error)
	Delete(userId int) error
	FindById(userId int) (models.User, error)
	FindUser(username string, password string) (models.User, error)
	FindAll(request.PaginationRequest) (users response.PaginationResponse[models.User], err error)
	Save(user *models.User)
}
