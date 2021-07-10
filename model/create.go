package model

import "fmt"

func CreateTodo(name, todo string) error {
	insertQ, err := con.Query("insert into Todo values(?, ?)", name, todo)
	if err != nil {
		fmt.Println(err)
		defer insertQ.Close()
		return err
	}
	defer insertQ.Close()
	return nil
}
