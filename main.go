package main

import (
	"fmt"
	"udemy_todo/app/controllers"
	"udemy_todo/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	// u := &models.User{}

	// u.Name = "test"
	// u.Email = "test@exmaple.com"
	// u.Password = "testtest"
	// fmt.Println(u)

	// u.CreateUser()

	// u, _ := models.GetUser(1)
	// fmt.Println(u)

	// u.Name = "test2"
	// u.Email = "test2@emam.com"
	// u.UpdateUser()

	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// user, _ := models.GetUser(1)
	// user.CreateTodo("Second Todo")

	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	// todos, _ := models.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// todosByUser, _ := user.GetTodosByUser()
	// for _, v := range todosByUser {
	// 	fmt.Println(v)
	// }

	// t, _ := models.GetTodo(1)
	// t.Content = "Update Todo"
	// t.UpdateTodo()

	// t, _ := models.GetTodo(1)
	// t.DeleteTodo()
}
