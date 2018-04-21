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

func findLfPos(bs []byte) int {
	pos := -1
	for i := 0; i < len(bs); i++ {
		if bs[i] == /*LF*/ 0x0a {
			pos = i
			break
		}
	}
	return pos
}

func printLine(line string) {
	rline := []rune(line)
	for i := 0; i < len(rline); i++ {
		c := string(rline[i])
		if c == "\t" {
			fmt.Print("^I")
		} else if c == "\r" {
			fmt.Print("^M")
		} else if c == "\n" {
			fmt.Print("$\n")
		} else {
			fmt.Print(c)
		}
	}
}

func Cata(args []string) {
	file, err := os.Open(args[0])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var linebuf []byte
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if n == 0 {
			break
		} else {
			if err != nil {
				panic(err)
			}
			for {
				lfpos := findLfPos(buf)
				if lfpos == -1 {
					for i := 0; i < n; i++ {
						linebuf = append(linebuf, buf[i])
					}
					break
				}
				for i := 0; i <= lfpos; i++ {
					linebuf = append(linebuf, buf[i])
				}
				printLine(string(linebuf))
				linebuf = []byte{}
				buf = buf[(lfpos + 1):]
			}
		}
	}
	if len(linebuf) > 0 {
		printLine(string(linebuf))
	}
	fmt.Println("[EOF]")
}
