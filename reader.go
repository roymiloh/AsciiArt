package AsciiArt
import (
    "bufio"
    "io"
    "bytes"
    "errors"
)

// mapping
var numbers = map[string]rune{
    " _ " +
    "| |" +
    "|_|": '0',

    "   " +
    "  |" +
    "  |": '1',

    " _ " +
    " _|" +
    "|_ ": '2',

    " _ " +
    " _|" +
    " _|": '3',

    "   " +
    "|_|" +
    "  |": '4',

    " _ " +
    "|_ " +
    " _|": '5',

    " _ " +
    "|_ " +
    "|_|": '6',

    " _ " +
    "  |" +
    "  |": '7',

    " _ " +
    "|_|" +
    "|_|": '8',

    " _ " +
    "|_|" +
    " _|": '9',
}

type Reader struct {
    DigitsPerRecord  int
    LinesPerRecord   int
    ColumnsPerDigit int
    r               *bufio.Reader
}

var (
    ErrNotEnoughDigits = errors.New("wrong number of digits in value")
    ErrNoBreakLine = errors.New("no break line detected at all")
)

func NewReader(r io.Reader) *Reader {
    return &Reader{
        r: bufio.NewReader(r),
        DigitsPerRecord: 9,
        LinesPerRecord: 4,
        ColumnsPerDigit: 3,
    }
}

// Read reads the next value and wrap it in a Record
func (r *Reader) Read() (record *Record, err error) {
    rawValue, err := r.readRawRecord()
    if err != nil {
        return nil, err
    }

    value := r.parseRecord(rawValue)
    return &Record{ Number: value }, nil
}

// readRawRecord reads entire record as is and returns it
func (r *Reader) readRawRecord() (rawValue [][]byte, err error) {
    rawLineLength := r.DigitsPerRecord * r.ColumnsPerDigit

    for i := 0; i < r.LinesPerRecord - 1; i++ {
        rawLine, _, err := r.r.ReadLine()

        if err != nil {
            // possible io.EOF
            return nil, err
        }
        if (len(rawLine) != rawLineLength) {
            return nil, ErrNotEnoughDigits
        }

        rawValue = append(rawValue, rawLine)
    }

    // ignore line break below each value, but validates
    // it actually exists.
    br, _, _ := r.r.ReadRune()
    if br != '\n' {
        // fault tolerance, at least the next values could be read
        // in case there's no break line
        r.r.UnreadRune()
        return nil, ErrNoBreakLine
    }

    return rawValue, nil
}

// parseRecord gets a raw record and returns his normalized string representation.
func (r *Reader) parseRecord(rawValue [][]byte) string {
    // using bytes.Buffer instead of concatenation for better performance.
    var value bytes.Buffer

    for pos := 0; pos < r.DigitsPerRecord * r.ColumnsPerDigit; pos += r.ColumnsPerDigit {
        digit := r.parseDigit(rawValue, pos)
        value.WriteRune(digit)
    }

    return value.String()
}

// parseDigit gets a raw record and position and returns his normalized
// digit representation.
func (r *Reader) parseDigit(rawValue [][]byte, pos int) rune {
    var digit bytes.Buffer

    for line := 0; line < r.LinesPerRecord - 1; line++ {
        digit.Write(rawValue[line][pos : pos + r.ColumnsPerDigit])
    }

    return getNumber(digit.String())
}

// getNumber gets a raw representation of digit and returns
// the rune corresponding one or '?' if it couldn't detect it.
func getNumber(rawDigit string) rune {
    if val, ok := numbers[rawDigit]; ok {
        return val
    }

    return '?'
}
