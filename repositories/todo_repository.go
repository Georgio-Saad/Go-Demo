package repositories

import (
	"todogorest/data/request"
	"todogorest/data/response"
	"todogorest/models"
)

type TodoRepository interface {
	GetAll(request.PaginationRequest) (todos response.PaginationResponse, err error)
	GetById(todoId int) (models.Todo, error)
	Create(todo request.TodoRequest) (models.Todo, error)
	Update(todoDetails request.TodoRequest, todoId int) (models.Todo, error)
	Delete(todoId int) error
}
