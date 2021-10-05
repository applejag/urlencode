# urlencode

Super basic URL encoding utility. I needed one, so I decided to make one.

## Installation

Requires Go v1.16 (or higher)

```console
$ go install github.com/jilleJr/urlencode
```

## Features

- Encodes

- Decodes

- Colored output to highlight what's encoded/decoded

## Usage

```console
$ urlencode --help
Usage of urlencode: [-qd] [values...]

Encodes the input value for HTTP URL by default and prints
the encoded value to STDOUT.

Input is taken from the given arguments and prints the results
one per line, or uses each line from STDIN if no args are supplied.

Flags:
  -d, --decode   decodes, instead of encodes
  -h, --help     show this help text
  -q, --query    encode/decode value as query parameter value
```

