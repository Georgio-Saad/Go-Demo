package repositories

import (
	"errors"
	"todogorest/data/request"
	"todogorest/data/response"
	"todogorest/models"
	"todogorest/pagination"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	Db *gorm.DB
}

// FindById implements ProductRepository.
func (p *ProductRepositoryImpl) FindById(prodId int) (models.Product, error) {
	var product models.Product

	result := p.Db.Model(&models.Product{}).Where("id = ?", prodId).First(&product)

	if result.Error != nil {
		return models.Product{}, result.Error
	}

	return product, nil
}

// Create implements ProductRepository.
func (p *ProductRepositoryImpl) Create(prod request.CreateProductRequest) (models.Product, error) {
	var product models.Product
	var productAlreadyExists models.Product

	alreadyExistsResult := p.Db.Model(&models.Product{}).Where("product = ?", prod.Product).Or("slug = ?", prod.Slug).First(&productAlreadyExists)

	if alreadyExistsResult.Error == nil {
		return models.Product{}, errors.New("Product already exists")
	}

	product.Product = prod.Product
	product.Slug = prod.Slug

	result := p.Db.Create(&product)

	if result.Error != nil {
		return models.Product{}, result.Error
	}

	return product, nil
}

// Delete implements ProductRepository.
func (p *ProductRepositoryImpl) Delete(productId int) error {
	var product models.Product

	result := p.Db.Model(&models.Product{}).Where("id = ?", productId).Delete(&product, productId)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// FindAll implements ProductRepository.
func (p *ProductRepositoryImpl) FindAll(pageReq request.PaginationRequest) (response.PaginationResponse[models.Product], error) {
	res, resErr := pagination.Paginate[models.Product](models.Product{}, pageReq, p.Db)

	if resErr != nil {
		return response.PaginationResponse[models.Product]{}, resErr
	}

	return res, nil
}

// Save implements ProductRepository.
func (p *ProductRepositoryImpl) Save(product *models.Product) {
	p.Db.Save(&product)
}

// Update implements ProductRepository.
func (p *ProductRepositoryImpl) Update(prod request.UpdateProductRequest) (models.Product, error) {
	var product models.Product

	result := p.Db.Model(&models.Product{}).Where("id = ?", prod.ProductID).First(&product)

	if result.Error != nil {
		return models.Product{}, nil
	}

	product.Product = prod.Product

	p.Db.Save(&product)

	return product, nil
}

func NewProductRepositoryImpl(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{Db: db}
}
