package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var baseURL = "http://localhost:8080"

type student struct {
	ID    string
	Name  string
	Grade float32
}

var data = []student{
	student{"A1", "ethan", 21},
	student{"A2", "hawke", 22},
	student{"A3", "gibran", 2.3},
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var result, err = json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func fetchUsers() ([]student, error) {
	var (
		err    error
		client = &http.Client{}
		data   []student
	)

	request, err := http.NewRequest("GET", baseURL+"/users", nil)

	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func main() {

	http.HandleFunc("/users", users)
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)

	var users, err = fetchUsers()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	for _, each := range users {
		fmt.Printf("ID: %s\t Name: %s\t Grade %f\n", each.ID, each.Name, each.Grade)
	}
}
