package fileutils

import (
    "io"
    "log"
    "sort"
)

func runeFrequencies(filename string) map[rune]int {
    reader := newBufferedReader(filename)
    runes := map[rune]int{}
    for {
        r, _, err := reader.ReadRune()
        if err == io.EOF {
            break
        } else if err != nil {
            log.Fatal(err)
        }
        if r > 0x19 { // Exclude common control codes
            runes[r] += 1
        }
    }
    return runes
}

func NumberOfLines(filename string) int64 {
    reader := newBufferedReader(filename)
    var lines int64
    for {
        _, err := reader.ReadString(0x0A)
        if err == io.EOF {
            break
        } else if err != nil {
            log.Fatal(err)
        }
        lines++
    }
    return lines
}

func TopRunes(filename string) PairList {
    runefreq := runeFrequencies(filename)
    pl := make(PairList, len(runefreq))
    i := 0
    for k, v := range runefreq {
        pl[i] = Pair{k, v}
        i++
    }
    sort.Sort(sort.Reverse(pl))
    return pl
}
