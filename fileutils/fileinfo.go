package fileutils

import (
    "path/filepath"
    "os"
    "fmt"
)

// PrintFileInfo prints information about a file, in the format
// specified in the assignment.
func PrintFileInfo(filename string) {
    path, err := filepath.Abs(filename)
    check(err)
    file, err := os.Lstat(filename)
    check(err)
    b := float64(file.Size())
    kib := b / 1024
    mib := kib / 1024
    gib := mib / 1024
    fmt.Printf("Information about '%s':\n", path)
    fmt.Printf("Size: %.0fB, %fKiB, %fMiB, %.9fGiB\n", b, kib, mib, gib)
    mode := file.Mode()
    // Sorry about the golf
    m := map[bool]string{true: "Is", false: "Is not"}
    fmt.Printf("\t%s a directory\n", m[mode.IsDir()])
    fmt.Printf("\t%s a regular file\n", m[mode.IsRegular()])
    fmt.Printf("Unix permission bits: %s\n", mode)
    fmt.Printf("\t%s append only\n", m[mode&os.ModeAppend != 0])
    fmt.Printf("\t%s a device file\n", m[mode&os.ModeDevice != 0])
    fmt.Printf("\t%s a Unix character device\n", m[(mode&os.ModeDevice != 0)&&(mode&os.ModeCharDevice != 0)])
    fmt.Printf("\t%s a Unix block device\n", m[(mode&os.ModeDevice != 0)&&(mode&os.ModeCharDevice == 0)])
    fmt.Printf("\t%s a symbolic link\n", m[mode&os.ModeSymlink != 0])
}
