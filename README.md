# Base64 VLQ Encode/Decode in Golang

![Test](https://github.com/sampsonbryce/go-vlq/actions/workflows/ci.yml/badge.svg)

References: 
- https://github.com/mozilla/source-map
- https://github.com/Rich-Harris/vl

## Why?

I wanted to parse JS sourcemaps in golang and understand how they really worked. I could find a good vlq encoder/decoder in golang that I could understand so I wrote my own.

## Usage

```go
import github.com/sampsonbryce/go-vlq

toEncode := []int{0, 0, 0, 0}
encoded := vlq.Encode(toEncode) // AAAA

decoded := vlq.Decode(encoded) // [0, 0, 0, 0]
```