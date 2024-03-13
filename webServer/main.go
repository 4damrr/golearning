package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
)

type Todo struct {
	ID     uuid.UUID
	Todo   string
	Detail string
}

var todos = []Todo{
	Todo{ID: uuid.New(), Todo: "Makan", Detail: "Ayam Bakar"},
	Todo{ID: uuid.New(), Todo: "Minum", Detail: "Jus Jeruk"},
	Todo{ID: uuid.New(), Todo: "IPK Gibran", Detail: "2,3"},
}

func todosHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		var result, err = json.Marshal(todos)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return

	case "POST":
		var newTodo Todo
		err := json.NewDecoder(r.Body).Decode(&newTodo)
		if err != nil {
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		newTodo.ID = uuid.New()

		todos = append(todos, newTodo)
		json.NewEncoder(w).Encode(newTodo)

	case "DELETE":
		var deleteID Todo
		bytee, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(bytee, &deleteID)
		if err != nil {
			http.Error(w, "", http.StatusNotFound)
		}
		for i, todo := range todos {
			if todo.ID == deleteID.ID {
				todos = append(todos[:i], todos[i+1:]...)
				break
			}
		}
		http.Error(w, "", http.StatusNotFound)

	default:
		http.Error(w, "", http.StatusBadRequest)
	}
}

func todosById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var id = r.URL.Query().Get("ID")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "Error parsing uuid", http.StatusInternalServerError)
		return
	}
	fmt.Print("test")

	switch r.Method {
	case "GET":
		for _, todo := range todos {
			if todo.ID == parsedID {
				var todo, _ = json.Marshal(todo)
				w.Write(todo)
				return
			}
		}
	case "PUT":
		var updatedTodo Todo
		updatedTodo.ID = parsedID
		fmt.Print(updatedTodo)
		err := json.NewDecoder(r.Body).Decode(&updatedTodo)
		if err != nil {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		for i, todo := range todos {
			if todo.ID == parsedID {
				todo = updatedTodo
				var todo, _ = json.Marshal(todo)
				todos = append(todos[:i], todos[i+1:]...)
				todos = append(todos, updatedTodo)
				w.Write(todo)
				return
			}
		}
	case "DELETE":
		for i, todo := range todos {
			if todo.ID == parsedID {
				todos = append(todos[:i], todos[i+1:]...)
				break
			}
		}
		http.Error(w, "", http.StatusNotFound)

	default:
		http.Error(w, "", http.StatusBadRequest)
	}
}

func main() {
	http.HandleFunc("/todos", todosHandler)
	http.HandleFunc("/todo", todosById)
	http.ListenAndServe(":8080", nil)
}
