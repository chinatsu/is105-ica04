package fileutils

import (
  "bufio"
  "os"
  "log"
  "fmt"
  "io"
)

func PrintLinesFromFile(filename string) {
  reader := NewBufferedReader(filename)
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

func GetLineending(filename string) string {
  var lineending string
  reader := NewBufferedReader(filename)
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

func NewBufferedReader(filename string) (*bufio.Reader) {
  // Open file by filename...
  f, err := os.Open(filename)
  if err != nil {
    log.Fatal(err)
  }
  // ... and return a bufio.Reader
  return bufio.NewReader(f)
}
