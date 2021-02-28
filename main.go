package main

import (
	"fmt"
	"todo_app/config"
)

func main() {
	fmt.Println(config.Config.Port)
}
