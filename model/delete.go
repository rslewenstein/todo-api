package model

import "fmt"

func DeleteTodo(name string) error {
	insertQ, err := con.Query("delete from Todo where name=?", name)
	if err != nil {
		fmt.Println(err)
		defer insertQ.Close()
		return err
	}
	defer insertQ.Close()
	return nil
}
