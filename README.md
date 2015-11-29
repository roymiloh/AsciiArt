# AsciiArt

A reader for ascii art numbers:

     _  _  _        _  _  _  _
      ||_||_|  ||_| _|| | _||_
      ||_||_|  |  ||_ |_| _| _|
      
      output: 788142035
      
If a digit does not fit, it returns `?` for that digit.
      
## Usage

    reader := AsciiArt.NewReader(myInputReader)
    // read the next value
    record, err := reader.Read()
    record.Write(os.Stdout)

## Running tests
    go test 'github.com/roymiloh/AsciiArt' -v --run=Test

There's also an example test (`--run=Example`)<br />
[Code Coverage](http://gocover.io/github.com/roymiloh/AsciiArt)

## Configurations
For another mapping of numbers, you should fork it and change it internally, at least until next versions.
You should also override the following configurations (of `Reader`):
- DigitsPerRecord
- LinesPerRecord
- ColumnsPerDigit

## TODO
- Writer (from raw numbers to ascii art numbers)
- Better customization
