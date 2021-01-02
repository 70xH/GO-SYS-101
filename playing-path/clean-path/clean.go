package main

import (
  "os"
  "fmt"
  "path/filepath"
)

func main() {

  // check for args

  if len(os.Args) > 2 {
    fmt.Println("clean: Expected 0-1 args, given", len(os.Args))
    return
  }

  // declare variables

  var cleanPath string

  if len(os.Args) == 2 {
    cleanPath = filepath.Clean(os.Args[1])                                          // returns no error
  } else {
    dir, err := os.Getwd()                                                          // when no argument is passed, get the current directory

    if err != nil {
      fmt.Println(err)
      return
    }

    cleanPath = filepath.Clean(dir)
  }

  fmt.Println(cleanPath)
}
