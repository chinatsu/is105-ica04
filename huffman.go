package main

import "./huffman"
import "fmt"
import "log"

func main() {
  d, err := huffman.DecodeHuffmanString("011011011100100000110")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(d)
}
