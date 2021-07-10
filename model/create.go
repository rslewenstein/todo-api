package model

import "fmt"

func CreateTodo() error {
	insertQ, err := con.Query("insert into Todo values(?, ?)", "Rafael", "Primeiro Teste")
	defer insertQ.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
