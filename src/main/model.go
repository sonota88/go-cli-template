package model

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

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

func Cat(args []string) {
	var fp *os.File
	var err error

	path := args[0]
	fmt.Println(path)

	fp, err = os.Open(path)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReaderSize(fp, 4096)

	for {
		line, _, err := reader.ReadLine()
		fmt.Println(string(line))
		if err == io.EOF {
			break
		}
	}
}
