# strand

A lightweight Go library for generating random strings — both cryptographically secure (`crypto/rand`) and deterministic seeded (`math/rand/v2`) — with context cancellation support and predefined character sets.

[![PkgGoDev](https://pkg.go.dev/badge/everlastingbeta/strand)](https://pkg.go.dev/github.com/everlastingbeta/strand)
[![Go Report Card](https://goreportcard.com/badge/everlastingbeta/strand?style=flat-square)](https://goreportcard.com/report/everlastingbeta/strand)
![test](https://github.com/everlastingbeta/strand/workflows/test/badge.svg)
![golangci-lint](https://github.com/everlastingbeta/strand/workflows/golangci-lint/badge.svg)

## Features

- Cryptographically secure random output via `crypto/rand`
- Deterministic seeded output via `math/rand/v2` for reproducible fixtures
- Context-aware variants for cancellation and timeouts
- Predefined character sets, plus support for any custom string
- Zero runtime dependencies

## Requirements

Go 1.26 or newer.

## Installation

```sh
go get -u github.com/everlastingbeta/strand
```

## Quick Start

```go
package main

import (
    "fmt"

    "github.com/everlastingbeta/strand"
)

func main() {
    // Cryptographically secure random string.
    token, err := strand.String(16, strand.UppercaseAlphabet)
    if err != nil {
        panic(err)
    }
    fmt.Println("Secure token:", token)

    // Deterministic string from a fixed seed.
    fixture := strand.SeededString(12, strand.ALL, 42)
    fmt.Println("Seeded fixture:", fixture)
}
```

## Usage Examples

### Cryptographically Secure Random Generation

Use these functions for security-sensitive output like tokens, passwords, and API keys.

```go
// Random byte slice drawn from an alphanumeric charset.
b, err := strand.Bytes(12, strand.AlphaNumeric)
if err != nil {
    // Handle error.
}
fmt.Println("Random bytes:", string(b))

// Random string from a custom charset.
apiKey, err := strand.String(32, strand.AlphaNumeric+"-_")
if err != nil {
    // Handle error.
}
fmt.Println("API key:", apiKey)

// Panic-on-error variants for callers that statically know inputs are valid.
token := strand.MustString(16, strand.ALL)
fmt.Println("Secure token:", token)
```

### Deterministic Random Generation

Use these functions when you need reproducible results.

```go
// Time-seeded (non-reproducible across runs).
b := strand.SeededBytes(8, strand.Numbers)
fmt.Println("Seeded bytes:", string(b))

// Fixed-seed (identical output every run).
id := strand.SeededString(10, "ACDEFGHJKLMNPQRSTUVWXYZ23456789", 12345)
fmt.Println("Deterministic ID:", id)
```

### Context-Aware Functions

For operations that might need to be canceled or time out.

```go
import (
    "context"
    "time"
)

ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

result, err := strand.StringWithContext(ctx, 20, strand.ALL)
if err != nil {
    // Handle context cancellation or other errors.
}

seededResult, err := strand.SeededStringWithContext(ctx, 20, strand.AlphaNumeric, 42)
if err != nil {
    // Handle errors.
}
```

## Available Character Sets

| Constant            | Contents                                  |
|---------------------|-------------------------------------------|
| `UppercaseAlphabet` | `A-Z`                                     |
| `LowercaseAlphabet` | `a-z`                                     |
| `Alphabet`          | `a-z` and `A-Z`                           |
| `Numbers`           | `0-9`                                     |
| `AlphaNumeric`      | `Alphabet` + `Numbers`                    |
| `Symbols`           | ``<>,\./|?;:[]{}+=_-()*&^%$#@!~``         |
| `ALL`               | `AlphaNumeric` + `Symbols`                |

Any string may be passed as a custom charset.

## Errors

The secure generators return one of the following sentinel errors, suitable for use with `errors.Is`:

| Error              | Condition                                 |
|--------------------|-------------------------------------------|
| `ErrInvalidSize`   | `size <= 0`                               |
| `ErrEmptyCharset`  | `charset == ""`                           |
| `ErrRandomFailure` | the underlying entropy source failed (wrapped) |

Context-aware variants additionally return an error wrapping `ctx.Err()` if the context is already done. `MustBytes` and `MustString` panic on the same conditions instead of returning the error.

```go
_, err := strand.String(0, strand.Alphabet)
if errors.Is(err, strand.ErrInvalidSize) {
    // ...
}
```

## Security Considerations

- `Bytes`, `String`, `MustBytes`, `MustString`, and their `*WithContext` variants use `crypto/rand` and are suitable for security-sensitive output.
- `SeededBytes`, `SeededString`, and their `*WithContext` variants use `math/rand/v2`. They are **not** cryptographically secure — use them only when reproducible output is required.

## License

[MIT](https://github.com/everlastingbeta/strand/blob/main/LICENSE)
