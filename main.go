package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	port := ":3333"
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/courses", getCourses)

	err := http.ListenAndServe(port, nil)

	if err != nil {
		fmt.Println(err)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {

	fmt.Println("get /Users")

	io.WriteString(w, "This is my user endpint!\n")

}

func getCourses(w http.ResponseWriter, r *http.Request) {

	fmt.Println("get /courses")

	io.WriteString(w, "This is my course endpint!\n")

}
