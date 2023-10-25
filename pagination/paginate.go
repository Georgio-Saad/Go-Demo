package pagination

import (
	"math"
	"todogorest/data/request"
	"todogorest/data/response"

	"gorm.io/gorm"
)

func Paginate[T interface{}](model interface{}, pageReq request.PaginationRequest, db *gorm.DB) (response.PaginationResponse[T], error) {
	var items []T

	offset := (pageReq.Page - 1) * pageReq.Size
	total := int64(0)

	result := db.Model(&model).Count(&total).Limit(pageReq.Size).Offset(offset).Find(&items)

	totalPages := math.Ceil(float64(total) / float64(int64(pageReq.Size)))

	hasNext := pageReq.Page < int(totalPages)

	hasPrev := pageReq.Page > 1

	if result.Error != nil {
		return response.PaginationResponse[T]{}, result.Error
	}

	return response.PaginationResponse[T]{
		PerPage:     pageReq.Size,
		CurrentPage: pageReq.Page,
		Items:       items, Total: int(total),
		FirstPage: 1,
		LastPage:  int(totalPages),
		HasNext:   hasNext,
		HasPrev:   hasPrev,
		Visible:   len(items),
	}, nil

}
