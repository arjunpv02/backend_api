package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
	"handlers"
    "models"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/students", handlers.GetStudents).Methods("GET")
    r.HandleFunc("/students", handlers.CreateStudent).Methods("POST")
    r.HandleFunc("/students/{id}", handlers.GetStudentByID).Methods("GET")
    r.HandleFunc("/students/{id}", handlers.UpdateStudent).Methods("PUT")
    r.HandleFunc("/students/{id}", handlers.DeleteStudent).Methods("DELETE")

    fmt.Println("Server is running on port 8000...")
    log.Fatal(http.ListenAndServe(":8000", r))
}
