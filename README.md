# strand

A lightweight, powerful Golang library for generating random strings with both cryptographically secure and seeded options.

[![PkgGoDev](https://pkg.go.dev/badge/everlastingbeta/strand)](https://pkg.go.dev/github.com/everlastingbeta/strand)
[![Go Report Card](https://goreportcard.com/badge/everlastingbeta/strand?style=flat-square)](https://goreportcard.com/report/everlastingbeta/strand)
![test](https://github.com/everlastingbeta/strand/workflows/test/badge.svg)
![golangci-lint](https://github.com/everlastingbeta/strand/workflows/golangci-lint/badge.svg)

## Features

- Generate cryptographically secure random strings using `crypto/rand`
- Create deterministic random strings with custom seeds using `math/rand/v2`
- Context-aware functions for cancellation support
- Predefined character sets for common use cases
- Simple, clean API with both error-returning and panic-on-error versions

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
    // Generate a secure random string with uppercase letters
    token, err := strand.String(16, strand.UppercaseAlphabet)
    if err != nil {
        panic(err)
    }
    fmt.Println("Secure token:", token)

    // Generate a predictable string with a custom charset and seed
    password := strand.SeededString(12, strand.ALL, 42)
    fmt.Println("Deterministic password:", password)
}
```

## Usage Examples

### Cryptographically Secure Random Generation

Use these functions for security-sensitive applications like tokens, passwords, and API keys.

```go
// Generate a secure random byte slice with alphanumeric characters
bytes, err := strand.Bytes(12, strand.AlphaNumeric)
if err != nil {
    // Handle error
}
fmt.Printf("Random bytes: %v\n", bytes)

// Generate a secure random string with custom character set
apiKey, err := strand.String(32, strand.AlphaNumeric + "-_")
if err != nil {
    // Handle error
}
fmt.Println("API key:", apiKey)

// Non-error-returning versions (will panic on error)
token := strand.MustString(16, strand.ALL)
fmt.Println("Secure token:", token)
```

### Deterministic Random Generation

Use these functions when you need reproducible results with a specific seed.

```go
// Generate a random byte slice with a timestamp-based seed
bytes := strand.SeededBytes(8, strand.Numbers)
fmt.Printf("Seeded bytes: %v\n", bytes)

// Generate a random string with a custom seed and charset
id := strand.SeededString(10, "ACDEFGHJKLMNPQRSTUVWXYZ23456789", 12345)
fmt.Println("Deterministic ID:", id)
```

### Context-Aware Functions

For operations that might need to be canceled or have timeouts.

```go
import (
    "context"
    "time"
)

// Create a context with timeout
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// Generate random strings with context support
result, err := strand.StringWithContext(ctx, 20, strand.ALL)
if err != nil {
    // Handle context cancellation or other errors
}

// Also available for seeded versions
seededResult, err := strand.SeededStringWithContext(ctx, 20, strand.AlphaNumeric, 42)
if err != nil {
    // Handle errors
}
```

## Available Character Sets

Strand provides several predefined character sets for convenience:

| Constant | Description |
|----------|-------------|
| `UppercaseAlphabet` | Uppercase letters (A-Z) |
| `LowercaseAlphabet` | Lowercase letters (a-z) |
| `Alphabet` | All letters (a-z, A-Z) |
| `Numbers` | Digits (0-9) |
| `AlphaNumeric` | All letters and digits |
| `Symbols` | Common special characters |
| `ALL` | All alphanumeric characters and symbols |

You can also define your own custom character sets as strings.

## Security Considerations

- The `Bytes()` and `String()` functions use `crypto/rand` and are suitable for security-sensitive applications.
- The `SeededBytes()` and `SeededString()` functions use `math/rand/v2` and are NOT cryptographically secure. Use them only when predictable output is required.

## License

[MIT](https://github.com/everlastingbeta/strand/blob/main/LICENSE)
