package AsciiArt_test
import (
    "github.com/roymiloh/AsciiArt"
    "os"
    "bufio"
    "strings"
    "io"
)

func ExampleWriter_Write() {
    input :=
        "    _  _     _  _  _  _  _ \n" +
        "  | _| _||_||_ |_   ||_||_|\n" +
        "  ||_  _|  | _||_|  ||_| _|\n\n" +
        "    _  _     _  _  _  _  _ \n" +
        "  | _| _||_||_ |_   ||_||_|\n" +
        "  ||_  _|  | _||_|   |_| _|\n\n" +
        " _  _  _     _  _     _  _ \n" +
        "  | _| _||_||_ |_   ||_||_|\n" +
        "  ||_  _|  | _||_|  ||_| _|\n\n"

    r := AsciiArt.NewReader(strings.NewReader(input))
    wr := bufio.NewWriter(os.Stdout)

    // a channel indicates on finish of writing process
    done := make(chan bool, 1)
    flushRate := 10

    // buffered channel for throttling purposes
    chRec := make(chan *AsciiArt.Record, flushRate)

    // writing process in his own goroutine
    go func() {
        recCount := 0

        for record := range chRec {
            recCount += 1
            record.Write(wr)

            if recCount % flushRate == 0 {
                wr.Flush()
            }
        }

        wr.Flush()
        done <- true
    }()

    // reading and sending values to chRec
    for {
        record, err := r.Read()
        if err != nil {
            if err == io.EOF {
                close(chRec)
                break;
            } else {
                // terminate the process if anything unexpected happens
                // ..or do something else, it actually depends on the error.
                panic(err)
            }
        }
        chRec <- record
    }

    <- done

    // Output:
    // 123456789
    // 123456?89 ILLEGAL
    // 723456189
}
