package main

import (
	"os"

	mytodolist "github.com/ekreke/myTodolist/internal"
)

func main() {
	command := mytodolist.NewMyTodolistCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
