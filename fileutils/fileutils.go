package fileutils

import (
    "bufio"
    "os"
    "log"
    "fmt"
    "io"
)

// check is a generic check function to use instead of
// the usual if sentences.
func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

// newBufferedReader reads a file by filename and returns
// a buffered reader for functions to later use.
func newBufferedReader(filename string) (*bufio.Reader) {
    // Open file by filename...
    f, err := os.Open(filename)
    check(err)
    // ... and return a bufio.Reader
    return bufio.NewReader(f)
}

// PrintLinesFromFile prints every line in a file including safely escaped
// sequences for control codes.
func PrintLinesFromFile(filename string) {
    reader := newBufferedReader(filename)
    for {
        line, err := reader.ReadString(0x0A)
        fmt.Printf("%+q\n", line) // Display newlines with %+q
        if err == io.EOF {
            // Read until EOF
            break
        } else if err != nil {
            log.Fatal(err)
        }
    }
}
