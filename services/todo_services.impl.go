package services

import (
	"net/http"
	"strconv"
	"todogorest/data/request"
	"todogorest/data/response"
	"todogorest/helpers"
	"todogorest/repositories"
	"todogorest/validations"
)

type TodoServicesImpl struct {
	TodoRepository repositories.TodoRepository
}

// Create implements TodoServices.
func (s *TodoServicesImpl) Create(todo request.TodoRequest) response.Response {
	validationErr := validations.ValidateRequest(todo)

	if validationErr != nil {
		return response.Response{StatusCode: http.StatusUnprocessableEntity, Code: helpers.UnprocessableEntity, Message: validationErr.Error()}
	}

	newTodo, err := s.TodoRepository.Create(todo)

	if err != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Code: helpers.BadRequest, Message: err.Error()}
	}

	return response.Response{StatusCode: http.StatusCreated, Message: "Successfully created todo", Code: helpers.Success, Data: newTodo}
}

// Delete implements TodoServices.
func (s *TodoServicesImpl) Delete(todoId string) response.Response {
	id, err := strconv.Atoi(todoId)

	if err != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Code: helpers.BadRequest, Message: err.Error()}
	}

	s.TodoRepository.Delete(id)

	return response.Response{StatusCode: http.StatusOK, Message: "Successfully deleted todo", Code: helpers.Success, Data: nil}

}

// FindAll implements TodoServices.
func (s *TodoServicesImpl) FindAll(pageReq request.PaginationRequest) response.Response {
	todos, err := s.TodoRepository.GetAll(pageReq)

	if err != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Message: err.Error(), Code: helpers.BadRequest}
	}

	return response.Response{StatusCode: http.StatusOK, Code: helpers.Success, Data: todos}
}

// FindById implements TodoServices.
func (s *TodoServicesImpl) FindById(todoId string) response.Response {
	id, err := strconv.Atoi(todoId)

	if err != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Message: err.Error(), Code: helpers.BadRequest}
	}

	todo, err := s.TodoRepository.GetById(id)

	if err != nil {
		return response.Response{StatusCode: http.StatusNotFound, Message: err.Error(), Code: helpers.NotFound}
	}

	return response.Response{StatusCode: http.StatusOK, Code: helpers.Success, Data: todo}
}

// Update implements TodoServices.
func (s *TodoServicesImpl) Update(todo request.TodoRequest, todoId string) response.Response {
	id, err := strconv.Atoi(todoId)

	if err != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Message: err.Error(), Code: helpers.BadRequest}
	}

	validationErr := validations.ValidateRequest(todo)

	if validationErr != nil {
		return response.Response{StatusCode: http.StatusUnprocessableEntity, Message: validationErr.Error(), Code: helpers.UnprocessableEntity}
	}

	updatedTodo, updateErr := s.TodoRepository.Update(todo, id)

	if updateErr != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Code: helpers.BadRequest, Message: updateErr.Error()}
	}

	return response.Response{StatusCode: http.StatusOK, Code: helpers.Success, Data: updatedTodo}

}

func NewTodoServicesImpl(todoRepository repositories.TodoRepository) TodoServices {
	return &TodoServicesImpl{
		TodoRepository: todoRepository,
	}
}
