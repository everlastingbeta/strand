# strand

Simple golang library that allows for quick generation of random strings.

[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/everlastingbeta/strand)
[![Go Report Card](https://goreportcard.com/badge/everlastingbeta/strand?style=flat-square)](https://goreportcard.com/report/everlastingbeta/strand)

## Example

[Library provided charsets](https://godoc.org/github.com/everlastingbeta/strand#pkg-constants)

```go
package main

import (
  "fmt"

  "github.com/everlastingbeta/strand"
)

func main() {
  // create a cryptographic pseudorandom byte slice
  cryptoBytes, err := strand.Bytes(12, strand.ALL)
  if err == nil {
    fmt.Printf("random crypto bytes: %v ==> %s\n", cryptoBytes, string(cryptoBytes))
  }

  // create a cryptographic pseudorandom string
  cryptoString, err := strand.String(12, strand.UppercaseAlphabet)
  if err == nil {
    fmt.Println("random crypto string: ", cryptoString)
  }

  // create a seeded pseudorandom byte slice
  seededBytes := strand.SeededBytes(12, strand.Symbols)
  fmt.Printf("random seeded bytes: %v ==> %s\n", seededBytes, string(seededBytes))

  // create a seeded pseudorandom string with a custom charset and seed
  seededString := strand.SeededString(12, "customCHARSET123!@#", 934006630)
  fmt.Println("random seeded string: ", seededString)

  /*  example output
  random crypto bytes: [108 105 57 119 36 126 121 84 89 71 94 64] ==> li9w$~yTYG^@
  random crypto string:  AIPJFZALAFOY
  random seeded bytes: [45 63 46 124 58 62 91 123 35 41 123 125] ==> -?.|:>[{#){}
  random seeded string:  tu2CcAc2C@HR
  */
}
```

## License

[MIT](https://github.com/everlastingbeta/strand/blob/master/LICENSE)
