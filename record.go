package AsciiArt

import (
    "fmt"
    "strings"
    "io"
)

type Record struct {
    Number string
}

type flag string

// Write the record's number and related flags to wr
func (i *Record) Write(wr io.Writer, flags ...flag) {
    if strings.Contains(i.Number, "?") {
        flags = append(flags, "ILLEGAL")
    }

    fmt.Fprintf(wr, i.Number)
    for _, flag := range flags {
        fmt.Fprintf(wr, " %v", flag)
    }
    fmt.Fprintln(wr)
}
