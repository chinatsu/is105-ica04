package fileutils

import (
    "os"
    "log"
)

func OsReadFile(filename string) {
    file, err := os.Open(filename)
    check(err)
    defer file.Close()
    byteSlice := make([]byte, 64)
    bytesRead, err := file.Read(byteSlice)
    check(err)
    log.Printf("Number of bytes read: %d\n", bytesRead)
    log.Printf("Data read: %s\n", byteSlice)
}
