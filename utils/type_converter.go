package utils

import (
	"strconv"

	db "github.com/phatnguyen138/go_api/internal/db/sqlc"
	"github.com/phatnguyen138/go_api/model"
)

func ConvertTodoToResponse(todo db.Todo) model.TodoResponse {
	return model.TodoResponse{
		Title:       todo.Title,
		Description: todo.Description.String,
		Completed:   strconv.FormatBool(todo.Completed.Bool),
	}
}
