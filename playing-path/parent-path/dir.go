package main

import (
  "os"
  "fmt"
  "path/filepath"
)

func main() {

  // check for args

  if len(os.Args) > 2 {
    fmt.Println("dir: Expected 0-1 args, given", len(os.Args))
    return
  }

  // declare variables

  var parPath string

  if len(os.Args) == 2 {
    parPath = filepath.Dir(os.Args[1])                                              // gets the parent path of a file or directory
  } else {
    dir, err := os.Getwd()                                                          // get the pwd if no argument is passed

    if err != nil {
      fmt.Println(err)
      return
    }

    parPath = filepath.Dir(dir)
  }

  fmt.Println(parPath)
}
