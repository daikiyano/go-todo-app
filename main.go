package main

import (
	"fmt"
	"log"
	"todo_app/app/models"
)

func main() {

	log.Println("test")

	fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test2"
	// u.Email = "test2@gmail.com"
	// u.PassWord = "testtest"
	// fmt.Println(u)
	// u.CreateUser()

	// u, _ := models.GetUser(1)
	// fmt.Println(u)

	// u.Name = "Test2"
	// u.Email = "test2@gmail.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u.DeleteUser()
	// user, _ := models.GetUser(3)
	// user.CreateTodo("third Todo")
	// fmt.Println(user)
	// todos, _ := models.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }
	user2, _ := models.GetUser(3)
	todos, _ := user2.GetTodosByUser()
	for _, v := range todos {
		fmt.Println(v)
	}

}
