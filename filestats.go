package main

import (
    "./fileutils"
    "fmt"
    "os"
    "flag"
)

func main() {
    filePtr := flag.String("f", "", "input file")
    runePtr := flag.Bool("runes", true, "count top 5 runes")
    linePtr := flag.Bool("lines", false, "count lines")
    flag.Parse()
    if *filePtr == "" {
        fmt.Fprintf(os.Stderr, "Usage of filestats:\n")
        flag.PrintDefaults()
    } else {
        if *linePtr {
            fmt.Printf("Amount of lines in '%s': %d\n", *filePtr, fileutils.NumberOfLines(*filePtr))
        } else if *runePtr {
            runes(*filePtr)
        }
    }
}


func runes(filename string) {
    var r int
    toprunes := fileutils.TopRunes(filename)
    if toprunes.Len() > 5 {
        r = 5
    } else {
        r = toprunes.Len()
    }
    for i := 0; i < r; i++ {
        fmt.Printf("%d: '%c' (%d occurrences)\n", i+1, toprunes[i].Key, toprunes[i].Value)
    }
}
