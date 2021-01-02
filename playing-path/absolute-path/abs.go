package main

import (
  "os"
  "fmt"
  "path/filepath"
)

func main() {

  // check for args

  if len(os.Args) > 2 {
    fmt.Println("abs: Expected 0-1 args, given", os.Args)
    return
  }

  // declare the variables

  var path string
  var err error

  if len(os.Args) == 2 {
    path, err = filepath.Abs(os.Args[1])
  } else {
    dir, err := os.Getwd()

    if err != nil {
      fmt.Println(err)
      return
    }

    path, err = filepath.Abs(dir)
  }

  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Println(path)
}
