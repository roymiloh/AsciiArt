package AsciiArt

import (
    "fmt"
    "strings"
    "bytes"
)

type Record struct {
    Number string
}

type flag string

// Write the record's number and related flags to wr
func (i *Record) String(flags ...flag) string {
    if strings.Contains(i.Number, "?") {
        flags = append(flags, "ILLEGAL")
    }

    if len(flags) == 0 {
        return i.Number
    }

    wr := &bytes.Buffer{}
    fmt.Fprintf(wr, i.Number)
    for _, flag := range flags {
        fmt.Fprintf(wr, " %v", flag)
    }

    return wr.String()
}
