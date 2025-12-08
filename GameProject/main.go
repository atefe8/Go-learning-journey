package main

import (
	"encoding/json"
	"fmt"
	"gameproject/repository/mysql"
	"gameproject/service/userservice"
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/user/register", userRegisterHandler)
	http.HandleFunc("/health-check", healthCheckHandler)
	http.HandleFunc("/user/login", userLoginHandler)
	fmt.Println("server start working...")
	err := http.ListenAndServe(":2021", nil)

	if err != nil {
		panic(err)
	}

}

func userLoginHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "method is not valid", http.StatusBadRequest)
	}

	data, error := io.ReadAll(req.Body)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}

	var userLoginRequest userservice.UserLoginRequest
	unError := json.Unmarshal(data, &userLoginRequest)
	if unError != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}

	repo := userservice.New(mysql.New())
	user, logErr := repo.Login(userLoginRequest)
	if logErr != nil {
		http.Error(w, logErr.Error(), http.StatusInternalServerError)
		return
	}

	type Response struct {
		Message string
		Data    interface{}
	}

	json.NewEncoder(w).Encode(Response{
		"user find!",
		user,
	})

}

func userRegisterHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "method is not valid", http.StatusBadRequest)
	}

	data, error := io.ReadAll(req.Body)

	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
	}

	var userRegisterRequest userservice.UserRegisterRequest

	marshalError := json.Unmarshal(data, &userRegisterRequest)
	if marshalError != nil {
		http.Error(w, marshalError.Error(), http.StatusInternalServerError)
	}

	repo := userservice.New(mysql.New())
	res, registerError := repo.Register(userRegisterRequest)
	if registerError != nil {
		http.Error(w, registerError.Error(), http.StatusInternalServerError)
	}

	type Response struct {
		Message string
		Data    interface{}
	}

	json.NewEncoder(w).Encode(Response{
		"user successfully created",
		res,
	})

}

func healthCheckHandler(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodGet {
		fmt.Fprintf(w, "invaild method")
		return

	}
	fmt.Fprintf(w, "server works")
}
