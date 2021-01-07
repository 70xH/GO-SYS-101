package main

import (
  "os"
  "io"
  "fmt"
)

func main() {

  // check for args

  if len(os.Args) != 2 {
    panic("readByteChunks: Expected 1 arg")
    return
  }

  file, err := os.Open(os.Args[1])                  // return a pointer to the File Descriptor (*File) or on failure an erro

  if err != nil {
    panic(err)
    return
  }

  defer file.Close()                                // close the file, avoiding any leaks

  bte := make([]byte, 32)

  for n := 0; err == nil; {
    n, err = file.Read(bte)                         // reads the buffer and prints the contents
    if err == nil {
      fmt.Println(string(bte[:n]))                  // print the stuff, limited to the size of the byte array
    } else if err != nil && err != io.EOF {         // exit on EOF or an error
      panic(err)
      return
    }
  }
}
