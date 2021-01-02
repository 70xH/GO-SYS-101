package main

import (
  "os"
  "fmt"
  "path/filepath"
)

func main() {

  // check for args

  if len(os.Args) > 2 {
    fmt.Println("base: Expected 0-1 args, given", len(os.Args))
    return
  }

  // declare variables

  var basePath string

  if len(os.Args) == 2 {
    basePath = filepath.Base(os.Args[1])                      // does not return any error
  } else {
    dir, err := os.Getwd()                                    // if no argument is passed use the current path

    if err != nil {
      fmt.Println(err)
      return
    }

    basePath = filepath.Base(dir)
  }

  fmt.Println(basePath)
}
