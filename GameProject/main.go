package main

import (
	"fmt"
	"gameproject/entity"
	"gameproject/repository/mysql"
	"net/http"
)

func main() {
	http.HandleFunc("/users/register", userRegisterHandler)
	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}

func userRegisterHandler(writer http.ResponseWriter, request *http.Request) {

	if request.Method != http.MethodPost {

		fmt.Fprintf(writer, "invaild method")

	}
	fmt.Fprintf(writer, "Handler Working!")

}

func testMethods() {
	mysqlRepo := mysql.New()

	createdUser, error := mysqlRepo.Register(entity.User{
		Name:        "ali",
		ID:          1,
		PhoneNumber: "091455",
	})

	if error != nil {
		fmt.Errorf("error happend %w", error)
	} else {
		fmt.Println("created user", createdUser)
	}

}
