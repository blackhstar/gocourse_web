package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	//http
	//port := ":3333"
	router := mux.NewRouter()

	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/courses", getCourses).Methods("GET")

	srv := &http.Server{
		Handler:      http.TimeoutHandler(router, time.Second*5, "Timeout!"),
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 3 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	//http
	//err := http.ListenAndServe(port, nil)
	//if err != nil {
	//	fmt.Println(err)
	//}
}

func getUsers(w http.ResponseWriter, r *http.Request) {

	//http
	//fmt.Println("get /Users")
	//io.WriteString(w, "This is my user endpint!\n")

	json.NewEncoder(w).Encode(map[string]bool{"ok": true})

}

func getCourses(w http.ResponseWriter, r *http.Request) {

	//http
	//fmt.Println("get /courses")
	//io.WriteString(w, "This is my course endpint!\n")

	json.NewEncoder(w).Encode(map[string]bool{"ok": true})

}
