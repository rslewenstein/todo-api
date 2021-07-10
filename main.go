package main

import (
	"fmt"
	"log"
	"net/http"

	"todo-api/controller"
	"todo-api/model"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	fmt.Println("Serving...")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
