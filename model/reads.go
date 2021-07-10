package model

import "todo-api/views"

func ReadAll() ([]views.PostRequest, error){
	rows, err := con.Query("SELECT * FROM Todo")
	if err != nil{
		return nil, err
	}
	todos := []views.PostRequest{}
	for rows.Next(){
		data := views.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}
	return todos, nil
}