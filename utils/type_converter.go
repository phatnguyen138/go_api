package utils

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/phatnguyen138/go_api/dto"
	db "github.com/phatnguyen138/go_api/internal/db/sqlc"
)

func ConvertTodoToResponse(todo db.Todo) dto.TodoResponse {
	return dto.TodoResponse{
		Title:       todo.Title,
		Description: todo.Description.String,
		Completed:   strconv.FormatBool(todo.Completed.Bool),
	}
}

func ConvertCreateTodoToParam(todo dto.CreateTodoRequest) (db.CreateTodoParams,error) {
	dueDate , err := StringToNullTime(todo.DueDate)
	if err != nil {
		println("Convert Create DTO into param failed with" + err.Error())
		return db.CreateTodoParams{}, err
	}
	
	resp := db.CreateTodoParams{
		Title: todo.Title,
		Description: sql.NullString{
			String: todo.Description,
			Valid:  true,
		},
		DueDate: dueDate,
	}

	return resp, nil
}

func StringToNullTime (s string) (sql.NullTime,error) {
	var nullTime sql.NullTime
    if s != "" {
        parsedTime, err := time.Parse("2006-01-02", s)
        if err != nil {
            return nullTime, err
        }
        nullTime = sql.NullTime{
            Time:  parsedTime,
            Valid: true,
        }
    }

	return nullTime,nil
}