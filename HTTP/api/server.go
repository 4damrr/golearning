package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var BaseURL = "http://localhost:8080"

type Data struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Detail string    `json:"detail"`
}

type Server struct {
	*mux.Router
	todos []Data
}

func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
		todos:  []Data{},
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/todos", s.listTodo()).Methods("GET")
	s.HandleFunc("/todos", s.addTodo()).Methods("POST")
}

func (s *Server) listTodo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(s.todos); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) addTodo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i Data
		if err := json.NewDecoder(r.Body).Decode(&s.todos); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		i.ID = uuid.New()
		s.todos = append(s.todos, i)

	}

}

func (s *Server) removeTodo(w http.ResponseWriter, r *http.Request) {

}
