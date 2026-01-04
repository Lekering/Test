package http

import (
	"errors"
	"net/http"
	"path/filepath"

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

	// API маршруты
	api := r.PathPrefix("/api").Subrouter()
	api.Path("/tasks").Methods("GET").Queries("done", "false").HandlerFunc(h.HTTPHandlers.HandleUnDoneTask)
	api.Path("/tasks").Methods("GET").HandlerFunc(h.HTTPHandlers.HandleGetTasks)
	api.Path("/tasks").Methods("POST").HandlerFunc(h.HTTPHandlers.HandleCreateTask)
	api.Path("/tasks/{title}").Methods("GET").HandlerFunc(h.HTTPHandlers.HandleGetOneTask)
	api.Path("/tasks/{title}").Methods("PATCH").HandlerFunc(h.HTTPHandlers.HandleDoneTask)
	api.Path("/tasks/{title}").Methods("DELETE").HandlerFunc(h.HTTPHandlers.HandleDeleteTask)

	// Статические файлы
	webDir, _ := filepath.Abs("./web")
	staticDir := http.Dir(webDir)
	fileServer := http.FileServer(staticDir)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	// Главная страница
	r.Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(webDir, "index.html"))
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil
}
