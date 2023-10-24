package services

import (
	"todogorest/data/request"
	"todogorest/data/response"
)

type UserServices interface {
	Create(user request.CreateUserRequest) response.Response
	Update() response.Response
	Delete(userId string) response.Response
	FindById(userId string) response.Response
	FindUser(request.SigninUserRequest) response.Response
	FindAll(request.PaginationRequest) response.Response
	Refresh(token string) response.Response
}
