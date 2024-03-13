package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type user struct {
	Username string
	Password string
}

type JsonResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func setJsonResp(Status int, Message string, w http.ResponseWriter) {
	var resp = JsonResponse{
		Status:  Status,
		Message: Message,
		Data:    nil,
	}
	response, _ := json.Marshal(resp)
	w.WriteHeader(resp.Status)
	w.Write(response)
}

var users = []user{
	user{Username: "bigbo55", Password: "smallbo55"},
	user{Username: "m4m4n", Password: "m1m1n"},
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user user
	err := json.NewDecoder(r.Body).Decode(&user)
	uname := user.Username
	pwd := user.Password
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case "GET":
		for _, user := range users {
			if user.Username == uname {
				if user.Password == pwd {
					fmt.Println("Login Success!")
					setJsonResp(http.StatusOK, "Login Success!", w)
					return
				}
				setJsonResp(http.StatusInternalServerError, "Passwordnya salah!", w)
				return
			}
		}
		setJsonResp(http.StatusBadRequest, "User not found!", w)
		return
	}
}

func main() {
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8080", nil)
}
