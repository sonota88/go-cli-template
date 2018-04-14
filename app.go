package main

import (
  "fmt"
  "os"
)

func main(){
  var currentDir string
  var projectDir string
  var subcmd string

  fmt.Println(os.Args[0])
  currentDir = os.Args[1]
  projectDir = os.Args[2]
  subcmd = os.Args[3]
  
  fmt.Println(currentDir)
  fmt.Println(projectDir)
  fmt.Println(subcmd)

  mainArgs := os.Args[4:]
  fmt.Println(mainArgs)
}
