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

type ProductServicesImpl struct {
	ProductRepository repositories.ProductRepository
}

// Create implements ProductServices.
func (p *ProductServicesImpl) Create(prod request.CreateProductRequest) response.Response {
	validationErr := validations.ValidateRequest(prod)

	if validationErr != nil {
		return response.Response{StatusCode: http.StatusUnprocessableEntity, Message: validationErr.Error(), Code: helpers.UnprocessableEntity}
	}

	res, resErr := p.ProductRepository.Create(prod)

	if resErr != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Message: resErr.Error(), Code: helpers.BadRequest}
	}

	return response.Response{StatusCode: http.StatusCreated, Message: "Successfully created product", Code: helpers.Success, Data: res}
}

// Delete implements ProductServices.
func (p *ProductServicesImpl) Delete(prodId string) response.Response {
	id, idErr := strconv.Atoi(prodId)

	if idErr != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Message: idErr.Error(), Code: helpers.BadRequest}
	}

	resErr := p.ProductRepository.Delete(id)

	if resErr != nil {
		return response.Response{StatusCode: http.StatusNotFound, Message: resErr.Error(), Code: helpers.NotFound}
	}

	return response.Response{StatusCode: http.StatusOK, Message: "Successfully deleted product", Code: helpers.Success}
}

// FindAll implements ProductServices.
func (p *ProductServicesImpl) FindAll(pageReq request.PaginationRequest) response.Response {
	prods, prodsErr := p.ProductRepository.FindAll(pageReq)

	if prodsErr != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Message: prodsErr.Error(), Code: helpers.BadRequest}
	}

	return response.Response{StatusCode: http.StatusOK, Message: "Successfully fetched products", Code: helpers.Success, Data: prods}
}

// FindById implements ProductServices.
func (p *ProductServicesImpl) FindById(prodId string) response.Response {
	id, idErr := strconv.Atoi(prodId)

	if idErr != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Message: idErr.Error(), Code: helpers.BadRequest}
	}

	prod, prodErr := p.ProductRepository.FindById(id)

	if prodErr != nil {
		return response.Response{StatusCode: http.StatusNotFound, Message: prodErr.Error(), Code: helpers.NotFound}
	}

	return response.Response{StatusCode: http.StatusOK, Message: "Successfully fetched product", Code: helpers.Success, Data: prod}
}

// Update implements ProductServices.
func (p *ProductServicesImpl) Update(prod request.UpdateProductRequest) response.Response {
	validationErr := validations.ValidateRequest(prod)

	if validationErr != nil {
		return response.Response{StatusCode: http.StatusUnprocessableEntity, Message: validationErr.Error(), Code: helpers.UnprocessableEntity}
	}

	prodUp, prodUpErr := p.ProductRepository.Update(prod)

	if prodUpErr != nil {
		return response.Response{StatusCode: http.StatusNotFound, Message: prodUpErr.Error(), Code: helpers.NotFound}
	}

	return response.Response{StatusCode: http.StatusOK, Message: "Successfully updated product", Code: helpers.Success, Data: prodUp}
}

func NewProductServicesImpl(productRepository repositories.ProductRepository) ProductServices {
	return &ProductServicesImpl{
		ProductRepository: productRepository,
	}
}
