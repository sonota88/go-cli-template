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

func Cat(opts map[string]string) {
	path := opts["in"]
	_, isInspect := opts["-A"]
	if isInspect {
		cat_a(path)
	} else {
		cat(path)
	}
}

func cat(path string) {
	var fp *os.File
	var err error

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

func readLine(file *os.File, buf *[]byte, bufSize int) ([]byte, bool) {
	var line []byte
	lfpos := findLfPos(*buf)

	if lfpos >= 0 {
		splitPos := lfpos + 1
		line = append(line, (*buf)[:splitPos]...)
		*buf = (*buf)[splitPos:]
		return line, false
	} else {
		line = append(line, (*buf)...)
	}

	for {
		readbuf := make([]byte, bufSize)
		n, err := file.Read(readbuf)
		*buf = readbuf[:n]
		if n == 0 {
			if len(line) > 0 {
				return line, false
			} else {
				return line, true
			}
		}
		if err != nil {
			panic(err)
		}

		lfpos := findLfPos(*buf)
		if lfpos >= 0 {
			splitPos := lfpos + 1
			line = append(line, (*buf)[:splitPos]...)
			*buf = (*buf)[splitPos:]
			return line, false
		} else {
			line = append(line, (*buf)[:n]...)
		}
	}
}

func cat_a(path string) {
	var r *bufio.Reader

	if path == "" {
		r = bufio.NewReader(os.Stdin)
	} else {
		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		r = bufio.NewReader(file)
	}

	for {
		line, err := r.ReadString('\n')
		printLine(line)
		if err == io.EOF {
			break
		}
	}

	fmt.Println("[EOF]")
}
