package AsciiArt
import (
    "testing"
    "strings"
    "io"
)

var readTests = []struct {
    Name    string
    Input   string
    Output  []string
}{
    {
        "Simple",
        "    _  _     _  _  _  _  _ \n" +
        "  | _| _||_||_ |_   ||_||_|\n" +
        "  ||_  _|  | _||_|  ||_| _|\n\n",
        []string{"123456789"},
    },
    {
        "Include illegal",
        "    _  _     _  _  _  _  _ \n" +
        "  | _| _||_||  |_   ||_||_|\n" +
        "  ||_  _|  | _||_|  ||_| _|\n\n",
        []string{"1234?6789"},
    },
    {
        "All illegal",
        "    _  _     _  _  _  _  _ \n" +
        "     |  || ||  |    ||_||_|\n" +
        "  ||_  _|  | _||_|   | |  |\n\n",
        []string{"?????????"},
    },
    {
        "Multiple valid records",
        "    _  _     _  _  _  _  _ \n" +
        "  | _| _||_||_ |_   ||_||_|\n" +
        "  ||_  _|  | _||_|  ||_| _|\n\n" +
        "    _  _     _  _  _  _  _ \n" +
        "  | _| _||_||_ |_   ||_||_|\n" +
        "  ||_  _|  | _||_|  ||_| _|\n\n",
        []string{"123456789", "123456789"},
    },
    {
        "Multiple records with illegal",
        "    _  _     _  _  _  _  _ \n" +
        "  | _| _||_||_ |_   ||_||_|\n" +
        "  ||_  _|  | _||_|  ||_| _|\n\n" +
        "    _  _     _  _  _  _  _ \n" +
        "  | _| _||_||_ |_   ||_||_|\n" +
        "  ||_  _|  | _||_|   |_| _|\n\n" +
        " _  _  _     _  _     _  _ \n" +
        "  | _| _||_||_ |_   ||_||_|\n" +
        "  ||_  _|  | _||_|  ||_| _|\n\n",
        []string{"123456789", "123456?89", "723456189"},
    },
}

var readTestsError = []struct {
    Name    string
    Input   string
    Error   error
}{
    {
        "Error not enough digits",
        "    _  _     _  _  _  _ \n" +
        "  | _| _||_||_ |_   ||_|\n" +
        "  ||_  _|  | _||_|  ||_|\n\n",
        ErrNotEnoughDigits,
    },
    {
        "Error no break line at the end",
        "    _  _     _  _  _  _  _ \n" +
        "  | _| _||_||_ |_   ||_||_|\n" +
        "  ||_  _|  | _||_|  ||_| _|\n",
        ErrNoBreakLine,
    },
}

func TestRead(t *testing.T) {
    for _, tt := range readTests {
        r := NewReader(strings.NewReader(tt.Input))
        recordIndex := -1

        for {
            out, err := r.Read()
            recordIndex += 1
            if err == io.EOF {
                break;
            }
            if err != nil {
                t.Errorf("%s: %v", tt.Name, err)
            } else if out.Number != tt.Output[recordIndex] {
                t.Errorf("%s: out=%v want=%v", tt.Name, out.Number, tt.Output[recordIndex])
            }
        }
    }
}

func TestRead_Error(t *testing.T) {
    for _, tt := range readTestsError {
        r := NewReader(strings.NewReader(tt.Input))

        _, err := r.Read()
        if err == nil {
            t.Errorf("%s: should returns an error")
        } else if err.Error() != tt.Error.Error() {
            t.Errorf("%s: got another error: %v", tt.Name, err)
        }
    }
}
