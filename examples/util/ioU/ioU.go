package ioU

import (
    "bufio"
    "os"
    "strings"
)

type LineHandler func(int, string)

func EachLine(filename string, handler LineHandler) {
    f, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    r := bufio.NewReader(f)
    var line string

    for lineNo := 0;; lineNo++ {
        if line, err = r.ReadString('\n'); err != nil {
            break
        }
        handler(lineNo, strings.Trim(line, "\n"))
    }
}
