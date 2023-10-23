package services

import (
	"todogorest/data/request"
	"todogorest/data/response"
)

type TodoServices interface {
	Create(todo request.TodoRequest) response.Response
	Update(todo request.TodoRequest, todoId string) response.Response
	Delete(todoId string) response.Response
	FindById(todoId string) response.Response
	FindAll(request.PaginationRequest) response.Response
}
