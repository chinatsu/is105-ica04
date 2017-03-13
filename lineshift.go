package main

import (
    "./fileutils"
    "fmt"
    "os"
)

func main() {
    args := os.Args[1:]
    if len(args) == 1 {
        fmt.Printf("%+q\n", fileutils.GetLineending(args[0])) // Safely print newline characters
    } else {
        fmt.Println("Usage: lineshift [filename]")
    }
}
