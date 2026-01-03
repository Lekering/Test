package http

import (
	"ReleaseHttpServer/todo"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPHandlers struct {
	todoList *todo.List
}

func NewHTTPHandlers(todo *todo.List) *HTTPHandlers {
	return &HTTPHandlers{
		todoList: todo,
	}
}

func jsonResponse(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func (h *HTTPHandlers) HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	var taskDto taskDto

	// Декодируем тело запроса
	if err := json.NewDecoder(r.Body).Decode(&taskDto); err != nil {
		errdto := NewErorDto(err.Error())
		jsonResponse(w, http.StatusBadRequest, errdto)
		return
	}
	if err := taskDto.Validate(); err != nil {
		jsonResponse(w, http.StatusBadRequest, NewErorDto(err.Error()))
		return
	}

	task := todo.NewTask(taskDto.Title, taskDto.Description)

	if err := h.todoList.AddTask(task); err != nil {
		jsonResponse(w, http.StatusInternalServerError, NewErorDto(err.Error()))
		return
	}
	jsonResponse(w, http.StatusCreated, todo.NewTask(taskDto.Title, taskDto.Description))
}

func (h *HTTPHandlers) HandleGetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.todoList.ListTasks()
	jsonResponse(w, http.StatusOK, tasks)
}

func (h *HTTPHandlers) HandleGetOneTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	task, err := h.todoList.GetTasks(title)
	if err != nil {
		jsonResponse(w, http.StatusNotFound, NewErorDto(err.Error()))
		return
	}

	jsonResponse(w, http.StatusOK, task)
}

func (h *HTTPHandlers) HandleUnDoneTask(w http.ResponseWriter, r *http.Request) {
	notDoneTasks := h.todoList.NotDoneTasks()

	jsonResponse(w, http.StatusOK, notDoneTasks)
}

func (h *HTTPHandlers) HandleDoneTask(w http.ResponseWriter, r *http.Request) {
	var doneTask doneDto

	title := mux.Vars(r)["title"]

	if err := json.NewDecoder(r.Body).Decode(&doneTask); err != nil {
		jsonResponse(w, http.StatusBadRequest, NewErorDto(err.Error()))
		return
	}

	var task todo.Task
	var err error

	if doneTask.Done {
		task, err = h.todoList.DoneTasks(title)
	} else {
		task, err = h.todoList.UnDoneTasks(title)
	}

	if err != nil {
		jsonResponse(w, http.StatusNotFound, NewErorDto(err.Error()))
		return
	}

	jsonResponse(w, http.StatusOK, task)
}

func (h *HTTPHandlers) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	if err := h.todoList.DeleteTask(title); err != nil {
		jsonResponse(w, http.StatusNotFound, NewErorDto(err.Error()))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
