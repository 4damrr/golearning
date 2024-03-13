package main

import (
	"encoding/json"
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

func main() {
	http.HandleFunc("/todos", todosHandler)
	http.ListenAndServe(":8080", nil)
}

// func todosGET(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	if r.Method == "GET" {
// 		var result, err = json.Marshal(todos)

// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		w.Write(result)
// 		return
// 	}

// 	http.Error(w, "", http.StatusBadRequest)
// }

// func todosPOST(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	if r.Method == "POST" {
// 		var newTodo Todo
// 		err := json.NewDecoder(r.Body).Decode(&newTodo)
// 		if err != nil {
// 			http.Error(w, "", http.StatusBadRequest)
// 			return
// 		}

// 		newTodo.ID = uuid.New()

// 		todos = append(todos, newTodo)
// 		json.NewEncoder(w).Encode(newTodo)

// 	}

// 	http.Error(w, "", http.StatusBadRequest)
// }

// func todosDELETE(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	if r.Method == "DELETE" {
// 		var deleteID Todo
// 		bytee, _ := ioutil.ReadAll(r.Body)
// 		err := json.Unmarshal(bytee, &deleteID)
// 		if err != nil {
// 			http.Error(w, "", http.StatusNotFound)
// 		}
// 		for i, todo := range todos {
// 			if todo.ID == deleteID.ID {
// 				todos = append(todos[:i], todos[i+1:]...)
// 				break
// 			}
// 		}
// 		http.Error(w, "", http.StatusNotFound)
// 	}

// 	http.Error(w, "", http.StatusBadRequest)
// }
