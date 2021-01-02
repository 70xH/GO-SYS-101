package main

import (
  "os"
  "fmt"
)

func main() {

  if len(os.Args) >= 2 {                                                // return if any argument is passed
    fmt.Println("pwd: Expected 0 args")
    return
  }

  pwd, err := os.Getwd()                                                // get the current/working directory

  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Println(pwd)                                                      // print the directory to console
}
