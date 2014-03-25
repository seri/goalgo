package util

import (
    "fmt"
    "io"
    "os"
)

type Dumpable interface {
    WriteTo(w io.Writer)
}

func Fail(filename, message string, data Dumpable) {
    fmt.Printf("Failed: %s\n", message)
    fmt.Printf("Exit now. See %s for details\n", filename)

    f, err := os.Create(filename)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    fmt.Fprintln(f, message)
    data.WriteTo(f)
    os.Exit(1)
}