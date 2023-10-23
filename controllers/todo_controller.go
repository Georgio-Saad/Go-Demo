package controllers

import (
	"net/http"
	"strconv"
	"todogorest/constants"
	"todogorest/data/request"
	"todogorest/data/response"
	"todogorest/helpers"
	"todogorest/services"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	todoService services.TodoServices
}

func (controller *TodoController) CreateTodo(ctx *gin.Context) {
	createTodoRequest := request.TodoRequest{}

	err := ctx.ShouldBindJSON(&createTodoRequest)

	if err != nil {
		ctx.JSON(http.StatusConflict, response.ErrorResponse{StatusCode: http.StatusConflict, Code: helpers.InvalidData, Data: response.ErrorMessage{Message: err.Error()}})
		return
	}

	res := controller.todoService.Create(createTodoRequest)

	if res.StatusCode != http.StatusCreated {
		ctx.JSON(res.StatusCode, response.ErrorResponse{StatusCode: res.StatusCode, Code: res.Code, Data: response.ErrorMessage{Message: res.Message}})
		return
	}

	ctx.JSON(res.StatusCode, gin.H{"todo": res.Data})
}

func (controller *TodoController) GetAllTodos(ctx *gin.Context) {
	page, pageErr := strconv.Atoi(ctx.Query("page"))
	size, sizeErr := strconv.Atoi(ctx.Query("limit"))

	var res response.Response

	if sizeErr != nil && (pageErr != nil || page < 1) {
		res = controller.todoService.FindAll(request.PaginationRequest{Page: 1, Size: constants.PerPage})
	} else if sizeErr != nil {
		res = controller.todoService.FindAll(request.PaginationRequest{Page: page, Size: constants.PerPage})
	} else if pageErr != nil || page < 1 {
		res = controller.todoService.FindAll(request.PaginationRequest{Page: 1, Size: size})
	} else {
		res = controller.todoService.FindAll(request.PaginationRequest{Page: page, Size: size})
	}

	if res.StatusCode != http.StatusOK {
		ctx.JSON(res.StatusCode, response.ErrorResponse{StatusCode: res.StatusCode, Code: res.Code, Data: response.ErrorMessage{Message: res.Message}})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"todos": res.Data})
}

func (controller *TodoController) GetTodo(ctx *gin.Context) {
	todoId := ctx.Param("todo_id")

	res := controller.todoService.FindById(todoId)

	if res.StatusCode != http.StatusOK {
		ctx.JSON(res.StatusCode, response.ErrorResponse{StatusCode: res.StatusCode, Code: res.Code, Data: response.ErrorMessage{Message: res.Message}})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"todo": res.Data})
}

func (controller *TodoController) DeleteTodo(ctx *gin.Context) {
	todoId := ctx.Param("todo_id")

	res := controller.todoService.Delete(todoId)

	if res.StatusCode != http.StatusOK {
		ctx.JSON(res.StatusCode, response.ErrorResponse{StatusCode: res.StatusCode, Code: res.Code, Data: response.ErrorMessage{Message: res.Message}})
		return
	}

	ctx.Status(res.StatusCode)
}

func (controller *TodoController) UpdateTodo(ctx *gin.Context) {
	todoId := ctx.Param("todo_id")

	createTodoRequest := request.TodoRequest{}

	reqError := ctx.ShouldBindJSON(&createTodoRequest)

	if reqError != nil {
		ctx.JSON(http.StatusConflict, response.ErrorResponse{StatusCode: http.StatusConflict, Code: "INVALID_DATA"})
		return
	}

	res := controller.todoService.Update(createTodoRequest, todoId)

	if res.StatusCode != http.StatusOK {
		ctx.JSON(res.StatusCode, response.ErrorResponse{StatusCode: res.StatusCode, Code: res.Code, Data: response.ErrorMessage{Message: res.Message}})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"todo": res.Data})
}

func NewTodoController(service services.TodoServices) *TodoController {
	return &TodoController{todoService: service}
}
