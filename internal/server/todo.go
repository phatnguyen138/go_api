package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/phatnguyen138/go_api/dto"
	db "github.com/phatnguyen138/go_api/internal/db/sqlc"
	"github.com/phatnguyen138/go_api/utils"
)

func (s *Server) ListTodoHandlder(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	limitStr := queryParams.Get("limit")
	offsetStr := queryParams.Get("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		http.Error(w, "Invalid offset parameter", http.StatusBadRequest)
		return
	}

	var resp []dto.TodoResponse
	listParam := db.ListTodosParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	todos, err := s.query.ListTodos(r.Context(), listParam)

	if err != nil {
		http.Error(w, "Invalid list query", http.StatusBadRequest)
		return
	}

	for _, eachtodo := range todos {
		resp = append(resp, utils.ConvertTodoToResponse(eachtodo))
	}

	jsonResp, err := json.Marshal(resp)

	if err != nil {
		http.Error(w, "Invalid response", http.StatusBadRequest)
		return
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) GetTodoHandler(w http.ResponseWriter,r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	if len(parts) != 3 {
		http.Error(w,"Invlid URL",http.StatusBadRequest)
	}

	id, err := strconv.Atoi(parts[2])

	if err != nil {
		http.Error(w,err.Error(), http.StatusBadRequest)
		return
	}

	todo, err := s.query.GetTodo(r.Context(),int32(id))
	println(int32(id))

	if err != nil {
		http.Error(w,err.Error(), http.StatusBadRequest)
		return
	}

	todoResp := utils.ConvertTodoToResponse(todo)

	resp, err := json.Marshal(todoResp)

	if err != nil {
		http.Error(w,"Invavlid todo convert", http.StatusBadRequest)
		return
	}

	_,_ = w.Write(resp)
}

func (s *Server) CreateTodoHandler(w http.ResponseWriter, r *http.Request){
	var todoCreate dto.CreateTodoRequest
	err := json.NewDecoder(r.Body).Decode(&todoCreate)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	todoParam, err := utils.ConvertCreateTodoToParam(todoCreate)
	
	if err != nil {
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}
	newTodo, err := s.query.CreateTodo(r.Context(),todoParam)

	if err != nil {
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}

	todoResp := utils.ConvertTodoToResponse(newTodo)

	resp, err := json.Marshal(todoResp)

	if err != nil {
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}

	_,_ = w.Write(resp)
}

func (s *Server) DeleteTodoHanlder (w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	if len(parts) != 3 {
		http.Error(w,"Invlid URL",http.StatusBadRequest)
	}

	id, err := strconv.Atoi(parts[2])

	if err != nil {
		http.Error(w,err.Error(), http.StatusBadRequest)
		return
	}

	deletedTodo, err := s.query.DeleteTodo(r.Context(),int32(id))

	if err != nil {
		http.Error(w,err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(utils.ConvertTodoToResponse(deletedTodo))

	if err != nil {
		http.Error(w,err.Error(), http.StatusBadRequest)
		return
	}

	_,_ = w.Write(resp)
}

func (s *Server) UpdateTodoHandler (w http.ResponseWriter, r *http.Request){
	var updateTodo dto.UpdateTodoRequest
	err := json.NewDecoder(r.Body).Decode(&updateTodo)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todoId, err := strconv.Atoi(updateTodo.Id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	oldTodo, err := s.query.GetTodo(r.Context(),int32(todoId))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newTodo, err := utils.ConvertUpdateTodoToParam(updateTodo)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

    // Update only the fields that are present in the request
    if updateTodo.Title != "" {
        oldTodo.Title = newTodo.Title
	}

    if updateTodo.Description != "" {
        oldTodo.Description = newTodo.Description
    }

    if updateTodo.Completed != nil {
        oldTodo.Completed = newTodo.Completed
		fmt.Printf("Updated completed")
    }
	
	newTodoParam, err := utils.ConvertTodoToUpdateTodoParam(oldTodo)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedTodo, err := s.query.UpdateTodo(r.Context(), newTodoParam)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todoResp := utils.ConvertTodoToResponse(updatedTodo)

	resp, err := json.Marshal(todoResp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_,_ = w.Write(resp)
}