package server

import (
	"encoding/json"
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