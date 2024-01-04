package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./user.env")
	if err != nil {
		panic(err)
	}
	fmt.Println("name:", os.Getenv("name"))
	fmt.Println("age:", os.Getenv("age"))
}
