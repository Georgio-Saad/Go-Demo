package controllers

import (
	"net/http"
	"strconv"
	"todogorest/data/request"
	"todogorest/data/response"
	"todogorest/helpers"
	"todogorest/services"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductServices services.ProductServices
}

func (controller *ProductController) CreateProduct(ctx *gin.Context) {
	createProductRequest := request.CreateProductRequest{}

	reqError := ctx.ShouldBindJSON(&createProductRequest)

	if reqError != nil {
		ctx.JSON(http.StatusConflict, response.ErrorResponse{StatusCode: http.StatusConflict, Code: helpers.InvalidData})
		return
	}

	res := controller.ProductServices.Create(createProductRequest)

	if res.StatusCode != http.StatusCreated {
		ctx.JSON(res.StatusCode, response.ErrorResponse{StatusCode: res.StatusCode, Code: res.Code, Data: response.ErrorMessage{Message: res.Message}})
		return
	}

	ctx.JSON(res.StatusCode, gin.H{"product": res.Data})
}

func (controller *ProductController) UpdateProduct(ctx *gin.Context) {
	prodId := ctx.Param("prod_id")

	id, idErr := strconv.Atoi(prodId)

	if idErr != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{StatusCode: http.StatusBadRequest, Code: helpers.BadRequest, Data: response.ErrorMessage{Message: idErr.Error()}})
		return
	}

	updateProductRequest := request.UpdateProductRequest{ProductID: id}

	reqError := ctx.ShouldBindJSON(&updateProductRequest)

	if reqError != nil {
		ctx.JSON(http.StatusConflict, response.ErrorResponse{StatusCode: http.StatusConflict, Code: helpers.InvalidData})
		return
	}

	res := controller.ProductServices.Update(updateProductRequest)

	if res.StatusCode != http.StatusOK {
		ctx.JSON(res.StatusCode, response.ErrorResponse{StatusCode: res.StatusCode, Code: res.Code, Data: response.ErrorMessage{Message: res.Message}})
		return
	}

	ctx.JSON(res.StatusCode, gin.H{"product": res.Data})
}

func (controller *ProductController) GetAllProducts(ctx *gin.Context) {
	page := ctx.Query("page")
	size := ctx.Query("limit")

	res := controller.ProductServices.FindAll(request.PaginationRequest{Page: page, Size: size})

	if res.StatusCode != http.StatusOK {
		ctx.JSON(res.StatusCode, response.ErrorResponse{StatusCode: res.StatusCode, Code: res.Code, Data: response.ErrorMessage{Message: res.Message}})
		return
	}

	ctx.JSON(res.StatusCode, gin.H{"products": res.Data})
}

func (controller *ProductController) GetProduct(ctx *gin.Context) {
	prodId := ctx.Param("prod_id")

	res := controller.ProductServices.FindById(prodId)

	if res.StatusCode != http.StatusOK {
		ctx.JSON(res.StatusCode, response.ErrorResponse{StatusCode: res.StatusCode, Code: res.Code, Data: response.ErrorMessage{Message: res.Message}})
		return
	}

	ctx.JSON(res.StatusCode, gin.H{"product": res.Data})
}

func (controller *ProductController) DeleteProduct(ctx *gin.Context) {
	prodId := ctx.Param("prod_id")

	res := controller.ProductServices.Delete(prodId)

	if res.StatusCode != http.StatusOK {
		ctx.JSON(res.StatusCode, response.ErrorResponse{StatusCode: res.StatusCode, Code: res.Code, Data: response.ErrorMessage{Message: res.Message}})
		return
	}

	ctx.Status(res.StatusCode)
}

func NewProductController(service services.ProductServices) *ProductController {
	return &ProductController{ProductServices: service}
}
