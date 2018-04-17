package main

import (
	"./src/main"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var subcmdFlag string

	fmt.Println(os.Args[0])
	subcmdFlag = os.Args[1]
	if subcmdFlag == "subcmd" {
		// TODO
	} else {
		mainSimple(os.Args[2:])
	}
}

func mainSimple(args []string) {
	var currentDir string
	var projectDir string

	currentDir = args[0]
	projectDir = args[1]

	fmt.Println(currentDir)
	fmt.Println(projectDir)

	mainArgs := args[2:]
	fmt.Println(mainArgs)

	a, _ := strconv.Atoi(mainArgs[0])
	b, _ := strconv.Atoi(mainArgs[1])
	model.Add(a, b)
}
