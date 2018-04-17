package main

import (
	"./src/main"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var currentDir string
	var projectDir string

	fmt.Println(os.Args[0])
	currentDir = os.Args[1]
	projectDir = os.Args[2]

	fmt.Println(currentDir)
	fmt.Println(projectDir)

	mainArgs := os.Args[3:]
	fmt.Println(mainArgs)

	a, _ := strconv.Atoi(mainArgs[0])
	b, _ := strconv.Atoi(mainArgs[1])
	model.Add(a, b)
}
