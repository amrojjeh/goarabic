
# GoArabic
[![GoDoc](https://godoc.org/github.com/amrojjeh/goarabic?status.svg)](https://godoc.org/github.com/amrojjeh/goarabic)

A Go Lang package for dealing with Arabic text. Forked from [01walid's package](github.com/01walid/goarabic).

## Current functionalities
- Glyph representation of the given Arabic text.
- Strip Tashkeel (Arabic Vowels).
- SmartLengh: return the length of the given Arabic String without considering Tashkeel (Arabic Vowels).
- Strip Tatweel
- rune-wise (UTF-8) reverse of the Arabic text, leaving out the Latin one.

## TODO
- Refine existing functions
- Add more utilities for dealing with Arabic text

## Usage

### Importing
```go
go get github.com/amrojjeh/goarabic@latest
```
##### Example Usage
```go
package main

import (
    "fmt"
    "github.com/amrojjeh/goarabic"
)

func main() {
    fmt.Println(goarabic.RemoveTashkeel("نًصٌ عَربيُّ"))
    fmt.Println(goarabic.ToGlyph("تجربة النص العربي"))
}

```

## Contributing
Contributions are greatly appreciated. Please fork this repository, make your changes, and open a pull request. More test cases and considerations might be needed, you can run tests using `go test` for the existing functionalities.

This a [SemVer](http://semver.org/)sioned package.
