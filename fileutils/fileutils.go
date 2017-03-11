package fileutils

import (
    "bufio"
    "os"
    "log"
    "fmt"
    "io"
    "path/filepath"
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


// GetLineending returns control codes for lineshifts in a given file.
// The first encountered line of at least 2 characters
// in length is used to safely determine whether it uses LF or CRLF.
func GetLineending(filename string) string {
    var lineending string
    reader := newBufferedReader(filename)
    for {
        line, err := reader.ReadString(0x0A)
        check(err)
        if len(line) > 1 {
            // Get the first line that is at least 2 characters long...
            lineending = line[len(line)-2:] // ... because the lineending is always at most 2 chars long
            break
        }
    }
    if lineending[0] > 0x1F {
        // If the first character is above 0x1F, it is not a control character,
        // and the last character is always bound to be a \n as that is what the line split by
        return lineending[1:]
    } else {
        return lineending
    }
}

func PrintFileInfo(filename string) {
    path, err := filepath.Abs(filename)
    check(err)
    file, err := os.Stat(filename)
    check(err)
    b := float64(file.Size())
    kib := b / 1024
    mib := kib / 1024
    gib := mib / 1024
    fmt.Printf("Information about '%s':\n", path)
    fmt.Printf("Size: %.0fB, %fKiB, %fMiB, %.9fGiB\n", b, kib, mib, gib)
    mode := file.Mode()
    if mode.IsDir() {
        fmt.Println("\tIs a directory")
    } else {
        fmt.Println("\tIs not a directory")
    }
    if mode.IsRegular() {
        fmt.Println("\tIs a regular file")
    } else {
        fmt.Println("\tIs not a regular file")
    }
    fmt.Printf("Unix permission bits: %s\n", file.Mode())
    if mode & os.ModeAppend != 0 {
        fmt.Println("\tIs append only")
    } else {
        fmt.Println("\tIs not append only")
    }
    if mode & os.ModeDevice != 0 {
        fmt.Println("\tIs a device file")
        if mode & os.ModeCharDevice != 0 {
            // If a file is a char device, it cannot be a block device
            // and vice versa.
            fmt.Println("\tIs a Unix character device")
            fmt.Println("\tIs not a Unix block device")
        } else {
            fmt.Println("\tIs not a Unix character device")
            fmt.Println("\tIs a Unix block device")
        }
    } else {
        fmt.Println("\tIs not a device file")
        fmt.Println("\tIs not a Unix character device")
        fmt.Println("\tIs not a Unix block device")
    }
    if mode & os.ModeSymlink != 0 {
        fmt.Println("\tIs a symbolic link")
    } else {
        fmt.Println("\tIs not a symbolic link")
    }
}
