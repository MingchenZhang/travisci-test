package main

import (
	"fmt"
	go_tty "github.com/mattn/go-tty"
)

func main() {
	fmt.Println("This is an automated build and release test of travis ci")
	tty, _ := go_tty.Open()
	defer tty.Close()
	pwd, _ := tty.ReadPassword()
	fmt.Printf("password: %s\n", pwd)
}
