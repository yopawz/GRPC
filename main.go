package main

import (
	"first-app/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/insert/employee", middleware.CreateEmployee).Methods("POST")
	router.HandleFunc("/employee", middleware.GetAllEmployee).Methods("GET")

	http.ListenAndServe(":8000", router)
}
