package model

import "fmt"

func Add(a int, b int) int {
	return a + b
}

// --------------------------------
// Commands

func Cmd1() {
	fmt.Println("cmd1")
}

func Cmd2() {
	fmt.Println("cmd2")
}
