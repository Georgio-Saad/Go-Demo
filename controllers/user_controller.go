package controllers

import (
	"log"
	"net/http"
	"todogorest/data/request"
	"todogorest/data/response"
	"todogorest/helpers"
	"todogorest/services"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserServices services.UserServices
}

func (controller *UserController) RefreshUser(ctx *gin.Context) {
	refreshUserRequest := request.RefreshUserRequest{}

	jsonErr := ctx.ShouldBindJSON(&refreshUserRequest)

	if jsonErr != nil {
		ctx.JSON(http.StatusConflict, response.ErrorResponse{StatusCode: http.StatusConflict, Code: helpers.InvalidData, Data: response.ErrorMessage{Message: jsonErr.Error()}})
		return
	}

	res := controller.UserServices.Refresh(refreshUserRequest.RefreshToken)

	if res.StatusCode != http.StatusOK {
		ctx.JSON(res.StatusCode, response.ErrorResponse{StatusCode: res.StatusCode, Code: res.Code, Data: response.ErrorMessage{Message: res.Message}})
		return
	}

	ctx.JSON(res.StatusCode, gin.H{"user": res.Data})
}

func (controller *UserController) Signup(ctx *gin.Context) {
	var role = "user"

	userRole := ctx.Query("user_role")

	if userRole == "admin" {
		role = userRole
	}
	log.Default().Println(role)

	createUserRequest := request.CreateUserRequest{Role: role}

	jsonErr := ctx.ShouldBindJSON(&createUserRequest)

	if jsonErr != nil {
		ctx.JSON(http.StatusConflict, response.ErrorResponse{StatusCode: http.StatusConflict, Code: helpers.InvalidData, Data: response.ErrorMessage{Message: jsonErr.Error()}})
		return
	}

	res := controller.UserServices.Create(createUserRequest)

	if res.StatusCode != http.StatusCreated {
		ctx.JSON(res.StatusCode, response.ErrorResponse{StatusCode: res.StatusCode, Code: res.Code, Data: response.ErrorMessage{Message: res.Message}})
		return
	}

	ctx.JSON(res.StatusCode, gin.H{"message": res.Message, "code": res.Code})
}

func (controller *UserController) Signin(ctx *gin.Context) {
	signinUserRequest := request.SigninUserRequest{}

	jsonErr := ctx.ShouldBindJSON(&signinUserRequest)

	if jsonErr != nil {
		ctx.JSON(http.StatusConflict, response.ErrorResponse{StatusCode: http.StatusConflict, Code: helpers.InvalidData, Data: response.ErrorMessage{Message: jsonErr.Error()}})
		return
	}

	res := controller.UserServices.FindUser(signinUserRequest)

	if res.StatusCode != http.StatusOK {
		ctx.JSON(res.StatusCode, response.ErrorResponse{StatusCode: res.StatusCode, Code: res.Code, Data: response.ErrorMessage{Message: res.Message}})
		return
	}

	ctx.JSON(res.StatusCode, gin.H{"user": res.Data})
}

func (controller *UserController) GetSignedInUser(ctx *gin.Context) {
	userId := helpers.GetAuthUserId(ctx)

	res := controller.UserServices.FindById(userId)

	if res.StatusCode != http.StatusOK {
		ctx.JSON(res.StatusCode, response.ErrorResponse{StatusCode: res.StatusCode, Code: res.Code, Data: response.ErrorMessage{Message: res.Message}})
		return
	}

	ctx.JSON(res.StatusCode, gin.H{"user": res.Data})
}

func (controller *UserController) GetUserById(ctx *gin.Context) {
	userId := ctx.Param("user_id")

	res := controller.UserServices.FindById(userId)

	if res.StatusCode != http.StatusOK {
		ctx.JSON(res.StatusCode, response.ErrorResponse{StatusCode: res.StatusCode, Code: res.Code, Data: response.ErrorMessage{Message: res.Message}})
		return
	}

	ctx.JSON(res.StatusCode, gin.H{"user": res.Data})

}

func (controller *UserController) VerifyUser(ctx *gin.Context) {
	userId := ctx.Param("user_id")
	login := ctx.Query("login")

	var isLogin bool = login == "true"

	verifyUserRequest := request.VerifyUserRequest{UserID: userId, Login: isLogin}

	jsonErr := ctx.ShouldBindJSON(&verifyUserRequest)

	if jsonErr != nil {
		ctx.JSON(http.StatusConflict, response.ErrorResponse{StatusCode: http.StatusConflict, Code: helpers.InvalidData, Data: response.ErrorMessage{Message: jsonErr.Error()}})
		return
	}

	res := controller.UserServices.Verify(verifyUserRequest)

	if res.StatusCode != http.StatusOK {
		ctx.JSON(res.StatusCode, response.ErrorResponse{StatusCode: res.StatusCode, Code: res.Code, Data: response.ErrorMessage{Message: res.Message}})
		return
	}

	ctx.JSON(res.StatusCode, gin.H{"user": res.Data})
}

func (controller *UserController) ResendVerification(ctx *gin.Context) {
	userId := ctx.Param("user_id")

	res := controller.UserServices.ResendVerification(userId)

	if res.StatusCode != http.StatusOK {
		ctx.JSON(res.StatusCode, response.ErrorResponse{StatusCode: res.StatusCode, Code: res.Code, Data: response.ErrorMessage{Message: res.Message}})
		return
	}

	ctx.JSON(res.StatusCode, gin.H{"message": res.Message, "code": res.Code})
}

func (controller *UserController) RemoveProfilePicture(ctx *gin.Context) {
	userId := helpers.GetAuthUserId(ctx)

	res := controller.UserServices.RemoveProfilePicture(userId)

	if res.StatusCode != http.StatusOK {
		ctx.JSON(res.StatusCode, response.ErrorResponse{StatusCode: res.StatusCode, Code: res.Code, Data: response.ErrorMessage{Message: res.Message}})
		return
	}

	ctx.JSON(res.StatusCode, gin.H{"user": res.Data})
}

func (controller *UserController) UploadProfilePicture(ctx *gin.Context) {
	sess := ctx.MustGet("sess").(*session.Session)
	userId := helpers.GetAuthUserId(ctx)

	profilePicture, profilePictureHeader, profilePictureErr := ctx.Request.FormFile("profile_picture")

	if profilePictureErr != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{StatusCode: http.StatusUnprocessableEntity, Code: helpers.UnprocessableEntity, Data: response.ErrorMessage{Message: profilePictureErr.Error()}})
		return
	}

	res := controller.UserServices.UploadProfilePicture(userId, profilePicture, profilePictureHeader, sess)

	if res.StatusCode != http.StatusOK {
		ctx.JSON(res.StatusCode, response.ErrorResponse{StatusCode: res.StatusCode, Code: res.Code, Data: response.ErrorMessage{Message: res.Message}})
		return
	}

	ctx.JSON(res.StatusCode, gin.H{"user": res.Data})
}

func NewUserController(service services.UserServices) *UserController {
	return &UserController{UserServices: service}
}
