package http

import (
	"ReleaseHttpServer/todo"
	"encoding/json"
	"net/http"
	"time"
)

type HTTPHandlers struct {
	todoList *todo.List
}

func NewHTTPHandlers(todo *todo.List) *HTTPHandlers {
	return &HTTPHandlers{
		todoList: todo,
	}
}

func (h *HTTPHandlers) HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	var todo = todo.NewTask("fgshsh", "sggs")

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		errdto := NewErorDto(err.Error())

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errdto)
		return
	}
}

type ErorDto struct {
	Massege string
	Time    time.Time
}

func NewErorDto(err string) *ErorDto {
	return &ErorDto{
		Massege: err,
		Time:    time.Now(),
	}
}
