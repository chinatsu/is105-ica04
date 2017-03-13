package main

import (
    "./fileutils"
    "fmt"
    "os"
    "flag"
)

func main() {
    filePtr := flag.String("f", "", "input file (required)")
    runePtr := flag.Bool("runes", true, "count top runes in file")
    linePtr := flag.Bool("lines", false, "count lines in file")
    numPtr := flag.Int("num", 5, "amount of runes to display in runes mode")
    flag.Parse()
    if *filePtr == "" {
        fmt.Fprintf(os.Stderr, "Usage of filestats:\n")
        flag.PrintDefaults()
    } else {
        if *linePtr {
            fmt.Printf("Amount of lines in '%s': %d\n", *filePtr, fileutils.NumberOfLines(*filePtr))
        } else if *runePtr {
            runes(*filePtr, *numPtr)
        }
    }
}


func runes(filename string, num int) {
    toprunes := fileutils.TopRunes(filename)
    if toprunes.Len() < num {
        num = toprunes.Len()
    }
    for i := 0; i < num; i++ {
        fmt.Printf("%d: '%c' (%d occurrences)\n", i+1, toprunes[i].Key, toprunes[i].Value)
    }
}
