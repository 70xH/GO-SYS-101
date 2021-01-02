package main

import (
  "os"
  "fmt"
)

// os.Chdir only changes the application working directory (not the shell working directory)

func main() {

  var err error

  if len(os.Args) != 2 {

    fmt.Println("cd: Expected 0-1 args")
    return

  }

  err = os.Chdir(os.Args[1])                                  // return a *PathError if any error

  if err != nil {
    fmt.Println(err)
    return
  }

  dir, err := os.Getwd()                                      // to see the current working directory

  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Println(dir)
}
