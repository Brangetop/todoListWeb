package http

import (
	"net/http"

	"brange.net/todoListWeb/todo"
)

type HTTPHandlers struct {
	todoList *todo.List
}

func newHTTPHandlers(todoList *todo.List) *HTTPHandlers {
	return &HTTPHandlers{
		todoList: todoList,
	}
}

// /tasks, POST, JSON
func (h *HTTPHandlers) HandleCreateTask(w http.ResponseWriter, r *http.Request) {

}

// /tasks/{title}, GET, pattern
func (h *HTTPHandlers) HandleGetTasks(w http.ResponseWriter, r *http.Request) {

}

// /tasks, GET, -
func (h *HTTPHandlers) HandleGetAllTasks(w http.ResponseWriter, r *http.Request) {

}

// /tasks?completed=false, GET, query
func (h *HTTPHandlers) HandleAllUncompletedTasks(w http.ResponseWriter, r *http.Request) {

}

// /tasks/{title}, PATCH, pattern +JSON
func (h *HTTPHandlers) HandleCompleteTask(w http.ResponseWriter, r *http.Request) {

}

// /tasks/{title}, DELETE, pattern + JSON
func (h *HTTPHandlers) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {

}
