package models

import "kala-v2/dao"

type Todo struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func CreateATodo(todo *Todo) (err error) {
	_, err = dao.DB.Exec("INSERT INTO kala (title,status) VALUES (?,?)", &todo.Title, &todo.Status)
	return err
}
func GetTodoList() (todo []*Todo,err error) {
	err = dao.DB.Select(&todo, "SELECT id,title,status FROM kala")
	if err != nil {
		return nil, err
	}
	return todo,nil
}
func UpdateATodo(id string,todo *Todo) (err error) {
	_,err=dao.DB.Exec("UPDATE kala SET status=? WHERE id=?", &todo.Status, id)
	return err
}
func DeleteATodo(id string) (err error) {
	_,err=dao.DB.Exec("DELETE FROM kala WHERE id=?", id)
	return err
}
