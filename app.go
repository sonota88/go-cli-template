package main

import (
	"./src/main"
	// "./src/main/maparse"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var subcmdFlag string

	fmt.Println(os.Args[0])
	subcmdFlag = os.Args[1]
	if subcmdFlag == "subcmd" {
		mainSubcmd(os.Args[2:])
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
	result := model.Add(a, b)
	fmt.Println(result)
}

func mainSubcmd(args []string) {
	var currentDir string
	var projectDir string
	var subcmd string

	currentDir = args[0]
	projectDir = args[1]
	subcmd = args[2]

	fmt.Println(currentDir)
	fmt.Println(projectDir)
	fmt.Println(subcmd)

	mainArgs := args[3:]
	fmt.Println(mainArgs)

	if subcmd == "cmd1" {
		model.Cmd1()
	} else if subcmd == "cmd2" {
		model.Cmd2()
	} else if subcmd == "cat" {
		model.Cat(mainArgs)
	} else {
		fmt.Println("Command not supported")
		os.Exit(1)
	}
}
