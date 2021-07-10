package main

import (
	"net/http"

	"todo-api/controller"
)

func main() {
	mux := controller.Register()

	http.ListenAndServe(":3000", mux)
}
