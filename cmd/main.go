package main

import (
	"fmt"
	"os"
)

// Go 程序的默认入口函数
func main() {
	fmt.Println(os.UserHomeDir())
}
