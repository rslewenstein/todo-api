package controller

import (
	"net/http"
	"todo-api/model"
)

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			//take some data
			// save it!
			if err := model.CreateTodo(); err != nil {
				w.Write([]byte("Some error"))
				return
			}
		}
	}
}
