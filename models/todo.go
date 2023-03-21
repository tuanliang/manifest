package models

import "manifest/dao"

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func CreateTodo(todo *Todo) error {
	if err := dao.DB.Create(&todo).Error; err != nil {
		return err
	}
	return nil
}
func GetAllTodo() (todoList []*Todo, err error) {
	if err := dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return todoList, nil
}
func GetTodoById(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err := dao.DB.Where("id=?", id).First(&todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}
func UpdateTodo(todo *Todo) error {
	if err := dao.DB.Save(todo).Error; err != nil {
		return err
	}
	return nil
}
func DeleteTodo(id string) error {
	err := dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	if err != nil {
		return err
	}
	return nil
}
