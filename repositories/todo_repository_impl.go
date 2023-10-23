package repositories

import (
	"math"
	"todogorest/data/request"
	"todogorest/data/response"
	"todogorest/models"

	"gorm.io/gorm"
)

type TodoRepositoryImpl struct {
	Db *gorm.DB
}

// Create implements TodoRepository.
func (t *TodoRepositoryImpl) Create(todoDetails request.TodoRequest) (models.Todo, error) {
	todo := models.Todo{Item: todoDetails.Item, Completed: todoDetails.Completed}

	result := t.Db.Create(&todo)

	if result.Error != nil {
		return todo, result.Error
	}

	return todo, nil
}

// Delete implements TodoRepository.
func (t *TodoRepositoryImpl) Delete(todoId int) error {
	var todo models.Todo

	todoToDelete := t.Db.First(&todo, todoId)

	if todoToDelete.Error != nil {
		return todoToDelete.Error
	}

	result := t.Db.Delete(&todo, todoId)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetAll implements TodoRepository.
func (t *TodoRepositoryImpl) GetAll(pageReq request.PaginationRequest) (response.PaginationResponse[models.Todo], error) {
	var todos []models.Todo

	offset := (pageReq.Page - 1) * pageReq.Size
	total := int64(0)

	result := t.Db.Model(&models.Todo{}).Count(&total).Limit(pageReq.Size).Offset(offset).Find(&todos)

	totalPages := math.Ceil(float64(total) / float64(int64(pageReq.Size)))

	hasNext := pageReq.Page < int(totalPages)

	hasPrev := pageReq.Page > 1

	if result.Error != nil {
		return response.PaginationResponse[models.Todo]{}, result.Error
	}

	return response.PaginationResponse[models.Todo]{
		PerPage:     pageReq.Size,
		CurrentPage: pageReq.Page,
		Items:       todos, Total: int(total),
		FirstPage: 1,
		LastPage:  int(totalPages),
		HasNext:   hasNext,
		HasPrev:   hasPrev,
		Visible:   len(todos),
	}, nil
}

// GetById implements TodoRepository.
func (t *TodoRepositoryImpl) GetById(todoId int) (models.Todo, error) {
	var todo models.Todo

	result := t.Db.First(&todo, todoId)

	if result.Error != nil {
		return todo, result.Error
	}

	return todo, nil
}

// Update implements TodoRepository.
func (t *TodoRepositoryImpl) Update(todoDetails request.TodoRequest, todoId int) (models.Todo, error) {
	var todo models.Todo

	result := t.Db.First(&todo, todoId)

	if result.Error != nil {
		return todo, result.Error
	}

	todo.Item = todoDetails.Item
	todo.Completed = todoDetails.Completed

	t.Db.Save(&todo)

	return todo, nil
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &TodoRepositoryImpl{Db: db}
}
