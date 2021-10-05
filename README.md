# urlencode

Super basic URL encoding utility. I needed one, so I decided to make one.

![Screenshot from 2021-10-05 17-19-29](https://user-images.githubusercontent.com/2477952/136052966-e00dcf9e-eece-4478-a7e5-e6703d95508d.png)

## Installation

Requires Go v1.16 (or higher)

```console
$ go install github.com/jilleJr/urlencode
```

## Features

- Encodes

- Decodes

- Colored output to highlight what's encoded/decoded

- Read from arguments, or STDIN if none supplied

## Usage

```console
$ urlencode --help
urlencode v1.0.0  Copyright (C) 2021  Kalle Jillheden

  License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>.
  This program comes with ABSOLUTELY NO WARRANTY; for details type '--license-w'
  This is free software, and you are welcome to redistribute it
  under certain conditions; type '--license-c' for details.

Encodes the input value for HTTP URL by default and prints
the encoded value to STDOUT.

Input is taken from the given arguments and prints the results
one per line, or uses each line from STDIN if no args are supplied.

Flags:
  -d, --decode    decodes, instead of encodes
  -h, --help      show this help text and exit
  -q, --query     encode/decode value as query parameter value
      --version   show version and exit
```

## License

Copyright &copy; 2021 Kalle Jillheden

License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>.
See full license text in the [LICENSE](./LICENSE) file.
