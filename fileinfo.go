package main

import (
  "./fileutils"
  "flag"
  "fmt"
  "os"
)

func main() {
    filePtr := flag.String("f", "", "input file")
    flag.Parse()
    if *filePtr == "" {
        fmt.Fprintf(os.Stderr, "Usage of fileinfo:\n")
        flag.PrintDefaults()
    } else {
        fileutils.PrintFileInfo(*filePtr)
    }
}
