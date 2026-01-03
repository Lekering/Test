package http

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	HTTPHandlers *HTTPHandlers
}

func NewHTTPServer(hand *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		HTTPHandlers: hand,
	}
}

func (h *HTTPServer) StartServer() error {
	r := mux.NewRouter()

	r.Path("/tasks").Methods("GET").Queries("done", "false").HandlerFunc(h.HTTPHandlers.HandleUnDoneTask)
	r.Path("/tasks").Methods("GET").HandlerFunc(h.HTTPHandlers.HandleGetTasks)
	r.Path("/tasks").Methods("POST").HandlerFunc(h.HTTPHandlers.HandleCreateTask)
	r.Path("/tasks/{title}").Methods("GET").HandlerFunc(h.HTTPHandlers.HandleGetOneTask)
	r.Path("/tasks/{title}").Methods("PATCH").HandlerFunc(h.HTTPHandlers.HandleDoneTask)
	r.Path("/tasks/{title}").Methods("DELETE").HandlerFunc(h.HTTPHandlers.HandleDeleteTask)

	if err := http.ListenAndServe(":8080", r); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil
}
