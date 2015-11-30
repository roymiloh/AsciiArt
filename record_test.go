package AsciiArt
import (
    "testing"
)

var recordWriteTests = []struct{
    Name    string
    Input   *Record
    Output  string
    Flags   []flag
}{
    {
        Name: "Simple",
        Input: &Record{ Number: "123456789" },
        Output: "123456789",
    },
    {
        Name: "Illegal flag",
        Input: &Record{ Number: "1?3456789" },
        Output: "1?3456789 ILLEGAL",
    },
    {
        Name: "Custom flag",
        Input: &Record{ Number: "123456789" },
        Output: "123456789 ERROR",
        Flags: []flag{"ERROR"},
    },
    {
        Name: "Multiple flags",
        Input: &Record{ Number: "123456789" },
        Output: "123456789 ERROR PERFECT",
        Flags: []flag{"ERROR", "PERFECT"},
    },
    {
        Name: "Illegal and custom flags",
        Input: &Record{ Number: "1234567?9" },
        Output: "1234567?9 ERROR ILLEGAL",
        Flags: []flag{"ERROR"},
    },
}

func TestRecord_Write(t *testing.T) {
    for _, tt := range recordWriteTests {
        out := tt.Input.String(tt.Flags...)
        if out != tt.Output {
            t.Errorf("%s: out=%v want=%v", tt.Name, out, tt.Output)
        }
    }
}
