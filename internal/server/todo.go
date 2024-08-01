package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	db "github.com/phatnguyen138/go_api/internal/db/sqlc"
	"github.com/phatnguyen138/go_api/model"
	"github.com/phatnguyen138/go_api/utils"
)

func (s *Server) ListTodoHanlder(w http.ResponseWriter, r *http.Request) {
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

	var resp []model.TodoResponse
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

func (s *Server) CreateTodoHandler(w http.ResponseWriter, r *http.Request) {

}
