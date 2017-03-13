package main

import "./fileutils"
import "flag"

func main() {
    filePtr := flag.String("f", "", "filename")
    flag.Parse()
    if *filePtr != "" {
        fileutils.OsReadFile(*filePtr)
    }
}
