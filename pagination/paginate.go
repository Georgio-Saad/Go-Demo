package pagination

import (
	"math"
	"strconv"
	"todogorest/constants"
	"todogorest/data/request"
	"todogorest/data/response"

	"gorm.io/gorm"
)

func Paginate[T interface{}](model interface{}, pageReq request.PaginationRequest, db *gorm.DB) (response.PaginationResponse[T], error) {
	var items []T

	var page, pageErr = strconv.Atoi(pageReq.Page)
	var size, sizeErr = strconv.Atoi(pageReq.Size)

	if pageErr != nil || page < 1 {
		page = 1
	}

	if sizeErr != nil || size < 1 {
		size = constants.PerPage
	}

	offset := (page - 1) * size
	total := int64(0)

	result := db.Model(&model).Count(&total).Limit(size).Offset(offset).Find(&items)

	totalPages := math.Ceil(float64(total) / float64(int64(size)))

	hasNext := page < int(totalPages)

	hasPrev := page > 1

	if result.Error != nil {
		return response.PaginationResponse[T]{}, result.Error
	}

	return response.PaginationResponse[T]{
		PerPage:     size,
		CurrentPage: page,
		Items:       items, Total: int(total),
		FirstPage: 1,
		LastPage:  int(totalPages),
		HasNext:   hasNext,
		HasPrev:   hasPrev,
		Visible:   len(items),
	}, nil

}
