package main

import (
  "os"
  "fmt"
  "io/ioutil"
)

func main () {

  // check for args

  if len(os.Args) != 2 {
    panic("rf: expected 1 args")
    return
  }

  file, err := ioutil.ReadFile(os.Args[1])                          // returns a byte array

  if err != nil {
    panic(err)
    return
  }

  fmt.Println(string(file))                                         // prints the byte array as string
}
