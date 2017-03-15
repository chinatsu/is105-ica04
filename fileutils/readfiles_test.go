package fileutils

import (
    "testing"
)

func BenchmarkOsReadText1(b *testing.B) {
    benchmarkOsReadFile("../files/text1.txt", b)
}

func BenchmarkOsReadPg100(b *testing.B) {
    benchmarkOsReadFile("../files/pg100.txt", b)
}

func benchmarkOsReadFile(filename string, b *testing.B) {
    for j := 0; j < b.N; j++ {
        b.StopTimer()
        b.StartTimer()
        _ = OsReadFile(filename)
    }
}

func BenchmarkIoReadText1(b *testing.B) {
    benchmarkIoReadFile("../files/text1.txt", b)
}

func BenchmarkIoReadPg100(b *testing.B) {
    benchmarkIoReadFile("../files/pg100.txt", b)
}

func benchmarkIoReadFile(filename string, b *testing.B) {
    for j := 0; j < b.N; j++ {
        b.StopTimer()
        b.StartTimer()
        _ = IoReadFile(filename)
    }
}

func BenchmarkIoutilReadText1(b *testing.B) {
    benchmarkIoutilReadFile("../files/text1.txt", b)
}

func BenchmarkIoutilReadPg100(b *testing.B) {
    benchmarkIoutilReadFile("../files/pg100.txt", b)
}

func benchmarkIoutilReadFile(filename string, b *testing.B) {
    for j := 0; j < b.N; j++ {
        b.StopTimer()
        b.StartTimer()
        _ = IoutilReadFile(filename)
    }
}

func BenchmarkBufioLineReadText1(b *testing.B) {
    benchmarkBufioReadLines("../files/text1.txt", b)
}

func BenchmarkBufioLineReadPg100(b *testing.B) {
    benchmarkBufioReadLines("../files/pg100.txt", b)
}

func benchmarkBufioReadLines(filename string, b *testing.B) {
    for j := 0; j < b.N; j++ {
        b.StopTimer()
        b.StartTimer()
        _ = BufioReadLines(filename)
    }
}

func BenchmarkBufioByteReadText1(b *testing.B) {
    benchmarkBufioReadBytes("../files/text1.txt", b)
}

func BenchmarkBufioByteReadPg100(b *testing.B) {
    benchmarkBufioReadBytes("../files/pg100.txt", b)
}

func benchmarkBufioReadBytes(filename string, b *testing.B) {
    for j := 0; j < b.N; j++ {
        b.StopTimer()
        b.StartTimer()
        _ = BufioReadBytes(filename)
    }
}
