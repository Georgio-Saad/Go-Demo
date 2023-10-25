package services

import (
	"todogorest/data/request"
	"todogorest/data/response"
)

type ProductServices interface {
	Create(prod request.CreateProductRequest) response.Response
	Update(prod request.UpdateProductRequest) response.Response
	Delete(prodId string) response.Response
	FindById(prodId string) response.Response
	FindAll(pageReq request.PaginationRequest) response.Response
}
