package request

type CreateProductRequest struct {
	Product string `validate:"required"`
	Slug    string `validate:"required"`
}

type UpdateProductRequest struct {
	ProductID int    `validate:"required"`
	Product   string `validate:"required"`
}
