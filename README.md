# GoJokes

A lightweight, zero-dependency Go library and CLI Tool for Random jokes

## Installation

### Library

```bash
go get github.com/yuvrajmodz/gojokes
```

### CLI

```bash
go install github.com/yuvrajmodz/gojokes/cmd/gojokes@latest
```

## Library Usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/yuvrajmodz/gojokes"
)

func main() {
	joke, err := gojokes.Random()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(joke)
}
```

## CLI Usage

```bash
# Print a random joke
gojokes

# Print a joke from a specific category
gojokes -c programming
gojokes --category golang

# List all available categories
gojokes --list
```

## Available Categories

| Category     | Description                        |
|--------------|------------------------------------|
| programming  | General software development humour |
| golang       | Go-specific jokes                  |
| linux        | Linux and Unix humour              |
| neutral      | Wordplay and general jokes         |

## API Reference

### `Random() (string, error)`

Returns a random joke selected uniformly from all available categories.

### `Category(name string) (string, error)`

Returns a random joke from the named category. Returns an error if the category does not exist.

### `Categories() []string`

Returns a sorted list of all available category names.

## Design

- **Zero dependencies** — uses only the Go standard library.
- **Embedded data** — all JSON joke files are compiled into the binary via `//go:embed`.
- **Single load** — categories are parsed once at startup using `sync.Once`.
- **Modern randomness** — uses `math/rand/v2` for unbiased selection.
- **Production Optimized** — Extremely Super Fast

## Requirements

- Go 1.21 or Later

## License

MIT
