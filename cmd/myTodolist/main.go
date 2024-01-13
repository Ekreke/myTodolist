package main

import (
	"os"

	"github.com/ekreke/myTodolist/internal/mytodolist"
)

func main() {
	command := mytodolist.NewMyTodolistCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
