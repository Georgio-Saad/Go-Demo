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

func (controller *UserController) RefreshUser(ctx *gin.Context) {
	refreshUserRequest := request.RefreshUserRequest{}

	err := ctx.ShouldBindJSON(&refreshUserRequest)

	if err != nil {
		ctx.JSON(http.StatusConflict, response.ErrorResponse{StatusCode: http.StatusConflict, Code: helpers.InvalidData, Data: response.ErrorMessage{Message: err.Error()}})
		return
	}

	res := controller.userServices.Refresh(refreshUserRequest.RefreshToken)

	if res.StatusCode != http.StatusOK {
		ctx.JSON(res.StatusCode, response.ErrorResponse{StatusCode: res.StatusCode, Code: res.Code, Data: response.ErrorMessage{Message: res.Message}})
		return
	}

	ctx.JSON(res.StatusCode, gin.H{"user": res.Data})
}

func (controller *UserController) Signup(ctx *gin.Context) {
	createUserRequest := request.CreateUserRequest{}

	err := ctx.ShouldBindJSON(&createUserRequest)

	if err != nil {
		ctx.JSON(http.StatusConflict, response.ErrorResponse{StatusCode: http.StatusConflict, Code: helpers.InvalidData, Data: response.ErrorMessage{Message: err.Error()}})
		return
	}

	res := controller.userServices.Create(createUserRequest)

	if res.StatusCode != http.StatusCreated {
		ctx.JSON(res.StatusCode, response.ErrorResponse{StatusCode: res.StatusCode, Code: res.Code, Data: response.ErrorMessage{Message: res.Message}})
		return
	}

	ctx.JSON(res.StatusCode, gin.H{"user": res.Data})
}

func (controller *UserController) Signin(ctx *gin.Context) {
	signinUserRequest := request.SigninUserRequest{}

	err := ctx.ShouldBindJSON(&signinUserRequest)

	if err != nil {
		ctx.JSON(http.StatusConflict, response.ErrorResponse{StatusCode: http.StatusConflict, Code: helpers.InvalidData, Data: response.ErrorMessage{Message: err.Error()}})
		return
	}

	res := controller.userServices.FindUser(signinUserRequest)

	if res.StatusCode != http.StatusOK {
		ctx.JSON(res.StatusCode, response.ErrorResponse{StatusCode: res.StatusCode, Code: res.Code, Data: response.ErrorMessage{Message: res.Message}})
		return
	}

	ctx.JSON(res.StatusCode, gin.H{"user": res.Data})
}

func (controller *UserController) GetSignedInUser(ctx *gin.Context) {
	userId := helpers.GetAuthUserId(ctx)

	res := controller.userServices.FindById(userId)

	if res.StatusCode != http.StatusOK {
		ctx.JSON(res.StatusCode, response.ErrorResponse{StatusCode: res.StatusCode, Code: res.Code, Data: response.ErrorMessage{Message: res.Message}})
		return
	}

	ctx.JSON(res.StatusCode, gin.H{"user": res.Data})
}

func (controller *UserController) GetUserById(ctx *gin.Context) {
	userId := ctx.Param("user_id")

	res := controller.userServices.FindById(userId)

	if res.StatusCode != http.StatusOK {
		ctx.JSON(res.StatusCode, response.ErrorResponse{StatusCode: res.StatusCode, Code: res.Code, Data: response.ErrorMessage{Message: res.Message}})
		return
	}

	ctx.JSON(res.StatusCode, gin.H{"user": res.Data})

}

func NewUserController(service services.UserServices) *UserController {
	return &UserController{userServices: service}
}
