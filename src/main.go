package main

import (
	"fmt"
	"net/http"
	"todolist/router"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting up http server...")

	r := mux.NewRouter()

	router.SetupRoutes(r)

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		panic(err)
	}
}
