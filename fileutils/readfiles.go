package fileutils

import (
    "os"
    "log"
    "io"
    "io/ioutil"
    "bufio"
)

// BufioReadLines reads a file line by line and appends it to byteSlice
// which is then logged to stdout
func BufioReadLines(filename string) []byte {
  f, err := os.Open(filename)
  check(err)
  defer f.Close()
  reader := bufio.NewReader(f)
  var byteSlice []byte
  for {
    line, err := reader.ReadBytes('\n')
    if err == io.EOF {
      break
    } else if err != nil {
      log.Fatal(err)
    }
  byteSlice = append(byteSlice, line...)
  }
  return byteSlice
}

// BufioReadBytes reads a file byte by byte and appends it to byteSlice
// which is then logged to stdout
func BufioReadBytes(filename string) []byte {
    var byteSlice []byte
    f, err := os.Open(filename)
    check(err)
    defer f.Close()
    reader := bufio.NewReader(f)
    for {
        b, err := reader.ReadByte()
        if err == io.EOF {
            break
        } else if err != nil {
            log.Fatal(err)
        }
        byteSlice = append(byteSlice, b)
    }
    return byteSlice
}

// IoutilReadFile reads the entire file into the byteSlice at once
func IoutilReadFile(filename string) []byte {
    byteSlice, err := ioutil.ReadFile(filename)
    check(err)
    return byteSlice
}

func IoReadFile(filename string) []byte {
    file, err := os.Open(filename)
    check(err)
    defer file.Close()
    info, err := os.Stat(filename)
    check(err)
    byteSlice := make([]byte, info.Size())
    _, err = io.ReadFull(file, byteSlice)
    check(err)
    return byteSlice
}

func OsReadFile(filename string) []byte {
    file, err := os.Open(filename)
    check(err)
    defer file.Close()
    info, err := os.Stat(filename)
    check(err)
    byteSlice := make([]byte, info.Size())
    _, err = file.Read(byteSlice)
    check(err)
    return byteSlice
}
