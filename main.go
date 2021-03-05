package main

import (
	"fmt"
	"log"
	"todo_app/app/models"
	"todo_app/config"
)

func main() {
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.LogFile)
	fmt.Println(config.Config.LogFile)

	log.Println("test")

	fmt.Println(models.Db)

	u := &models.User{}
	u.Name = "test"
	u.Email = "test@gmail.com"
	u.PassWord = "testtest"
	fmt.Println(u)
	u.CreateUser()

	// u, _ := models.GetUser(1)
	// fmt.Println(u)

	// u.Name = "Test2"
	// u.Email = "test2@gmail.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u.DeleteUser()
	user, _ := models.GetUser(2)
	user.CreateTodo("Second Todo")
	// fmt.Println(user)
	todos, _ := models.GetTodos()
	for _, v := range todos {
		fmt.Println(v)
	}

}
