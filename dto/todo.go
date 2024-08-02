package dto

type TodoResponse struct {
	Id 			string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   string `json:"completed"`
}

type CreateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date,omitempty"`
}

type UpdateTodoRequest struct {
	Id			string 	`json:"id"`
	Title       string 	`json:"title,omitempty"`
	Description string 	`json:"description,omitempty"`
	Completed   *bool   `json:"completed,omitempty"`
}
