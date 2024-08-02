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
		Id: strconv.Itoa(int(todo.ID)),
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

func ConvertUpdateTodoToParam(todo dto.UpdateTodoRequest) (db.UpdateTodoParams,error) {
	id, err := strconv.Atoi(todo.Id)

	if err != nil {
		println("Wrong id")
		return db.UpdateTodoParams{}, err
	}

	resp := db.UpdateTodoParams{
		ID: int32(id),
		Title: todo.Title,
		Description: sql.NullString{
			String: todo.Description,
			Valid: true,
		},
		Completed: sql.NullBool{
			Bool: *todo.Completed,
			Valid: true,
		},
	}
	return resp, nil
}

func ConvertTodoToUpdateTodoParam (todo db.Todo) (db.UpdateTodoParams, error) {
	return db.UpdateTodoParams{
		ID: todo.ID,
		Title: todo.Title,
		Description: todo.Description,
		Completed: todo.Completed,
	}, nil
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