package controllers

import (
	"net/http"
	"todogorest/data/request"
	"todogorest/data/response"
	"todogorest/helpers"
	"todogorest/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userServices services.UserServices
}

func (controller *UserController) Signup(ctx *gin.Context) {
	createTodoRequest := request.CreateUserRequest{}

	err := ctx.ShouldBindJSON(&createTodoRequest)

	if err != nil {
		ctx.JSON(http.StatusConflict, response.ErrorResponse{StatusCode: http.StatusConflict, Code: helpers.InvalidData, Data: response.ErrorMessage{Message: err.Error()}})
		return
	}

	res := controller.userServices.Create(createTodoRequest)

	if res.StatusCode != http.StatusCreated {
		ctx.JSON(res.StatusCode, response.ErrorResponse{StatusCode: res.StatusCode, Code: res.Code, Data: response.ErrorMessage{Message: res.Message}})
		return
	}

	ctx.JSON(res.StatusCode, res.Data)
}

func NewUserController(service services.UserServices) *UserController {
	return &UserController{userServices: service}
}
