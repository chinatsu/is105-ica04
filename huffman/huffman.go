package huffman

import (
  "errors"
)


func DecodeHuffmanString(code string) (string, error) {
  var i int
  var decoded string
  var err error
  key := getKey()
  for {
    i = 2
    if len(code) < i {
      err = errors.New("Invalid length of code")
      break
    } else if key[code[:i]] == 0 {
      i = 3
    }
    if len(code) < i {
      err = errors.New("Invalid length of code")
      break
    }
    if key[code[:i]] == 0 {
      err = errors.New("Malformed huffman code, no matching key")
      break
    }
    decoded += string(key[code[:i]])
    code = code[i:]
    if len(code) == 0 {
      break
    }
  }
  if err != nil {
    return "", err
  } else {
    return decoded, nil
  }
}

func getKey() map[string]rune {
  return map[string]rune{
    "000": 'A',
    "001": 'B',
    "111": 'C',
    "10": 'D',
    "110": 'E',
    "01": 'F',
  }
}
