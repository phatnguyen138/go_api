package dto

type TodoResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   string `json:"completed"`
}

type CreateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date,omitempty"`
}
