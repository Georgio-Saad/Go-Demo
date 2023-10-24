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

type VerificationCodeController struct {
	VerificationCodeServices services.VerificationCodeServices
}

func (controller *VerificationCodeController) GetVerificationCode(ctx *gin.Context) {
	verType := ctx.Query("get_by")
	id := ctx.Param("id")

	if verType == constants.VerificationCodeUserIdType {
		res := controller.VerificationCodeServices.FindByUserId(id)

		if res.StatusCode != http.StatusOK {
			ctx.JSON(res.StatusCode, response.ErrorResponse{StatusCode: res.StatusCode, Code: res.Code, Data: response.ErrorMessage{Message: res.Message}})
			return
		}

		ctx.JSON(res.StatusCode, gin.H{"verification_code": res.Data})
		return
	}

	res := controller.VerificationCodeServices.FindById(id)

	if res.StatusCode != http.StatusOK {
		ctx.JSON(res.StatusCode, response.ErrorResponse{StatusCode: res.StatusCode, Code: res.Code, Data: response.ErrorMessage{Message: res.Message}})
		return
	}

	ctx.JSON(res.StatusCode, gin.H{"verification_code": res.Data})
	return
}

func (controller *VerificationCodeController) CreateVerificationCode(ctx *gin.Context) {
	userId := ctx.Param("user_id")

	id, _ := strconv.Atoi(userId)

	createVerificationCodeRequest := request.VerificationCodeRequest{UserID: id}

	bindErr := ctx.ShouldBindJSON(&createVerificationCodeRequest)

	if bindErr != nil {
		ctx.JSON(http.StatusConflict, response.ErrorResponse{StatusCode: http.StatusConflict, Code: helpers.InvalidData, Data: response.ErrorMessage{Message: bindErr.Error()}})
		return
	}

	res := controller.VerificationCodeServices.Create(createVerificationCodeRequest)

	if res.StatusCode != http.StatusCreated {
		ctx.JSON(res.StatusCode, response.ErrorResponse{StatusCode: res.StatusCode, Code: res.Code, Data: response.ErrorMessage{Message: res.Message}})
		return
	}

	ctx.JSON(res.StatusCode, gin.H{"verification_code": res.Data})
}

func (controller *VerificationCodeController) UpdateVerificationCode(ctx *gin.Context) {
	userId := ctx.Param("user_id")

	id, _ := strconv.Atoi(userId)

	updateVerificationCodeRequest := request.VerificationCodeRequest{UserID: id}

	bindErr := ctx.ShouldBindJSON(&updateVerificationCodeRequest)

	if bindErr != nil {
		ctx.JSON(http.StatusConflict, response.ErrorResponse{StatusCode: http.StatusConflict, Code: helpers.InvalidData, Data: response.ErrorMessage{Message: bindErr.Error()}})
		return
	}

	res := controller.VerificationCodeServices.Update(updateVerificationCodeRequest)

	if res.StatusCode != http.StatusOK {
		ctx.JSON(res.StatusCode, response.ErrorResponse{StatusCode: res.StatusCode, Code: res.Code, Data: response.ErrorMessage{Message: res.Message}})
		return
	}

	ctx.JSON(res.StatusCode, gin.H{"verification_code": res.Data})
}

func (controller *VerificationCodeController) DeleteVerificationCode(ctx *gin.Context) {
	verType := ctx.Query("get_by")
	id := ctx.Param("id")

	if verType == constants.VerificationCodeUserIdType {
		res := controller.VerificationCodeServices.DeleteByUserId(id)

		if res.StatusCode != http.StatusOK {
			ctx.JSON(res.StatusCode, response.ErrorResponse{StatusCode: res.StatusCode, Code: res.Code, Data: response.ErrorMessage{Message: res.Message}})
			return
		}

		ctx.JSON(res.StatusCode, gin.H{"message": res.Message})
		return
	}

	res := controller.VerificationCodeServices.Delete(id)

	if res.StatusCode != http.StatusOK {
		ctx.JSON(res.StatusCode, response.ErrorResponse{StatusCode: res.StatusCode, Code: res.Code, Data: response.ErrorMessage{Message: res.Message}})
		return
	}

	ctx.JSON(res.StatusCode, gin.H{"message": res.Message})
	return
}

func NewVerificationCodeController(service services.VerificationCodeServices) *VerificationCodeController {
	return &VerificationCodeController{VerificationCodeServices: service}
}
