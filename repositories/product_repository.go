package repositories

import (
	"todogorest/data/request"
	"todogorest/data/response"
	"todogorest/models"
)

type ProductRepository interface {
	Create(request.CreateProductRequest) (models.Product, error)
	Update(request.UpdateProductRequest) (models.Product, error)
	FindAll(request.PaginationRequest) (response.PaginationResponse[models.Product], error)
	Delete(productId int) error
	Save(product *models.Product)
}
