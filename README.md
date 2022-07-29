<!--
SPDX-FileCopyrightText: 2021 Kalle Fagerberg

SPDX-License-Identifier: CC0-1.0
-->

# urlencode

[![REUSE status](https://api.reuse.software/badge/github.com/jilleJr/urlencode)](https://api.reuse.software/info/github.com/jilleJr/urlencode)

Super basic URL encoding utility. I needed one, so I decided to make one.

![urlencode-usage-screenshot](https://user-images.githubusercontent.com/2477952/136087896-7bc8c5ca-cbef-414f-bee1-50dbd5440259.png)

## Installation

Requires Go v1.18 (or higher)

```console
$ go install github.com/jilleJr/urlencode@latest
```

## Features

- Encodes

- Decodes

- Colored output to highlight what's encoded/decoded

- Read from STDIN or from a file

## Usage

```console
$ urlencode --help
urlencode v1.1.0  Copyright (C) 2021  Kalle Fagerberg

  License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>.
  This program comes with ABSOLUTELY NO WARRANTY; for details type '--license-w'
  This is free software, and you are welcome to redistribute it
  under certain conditions; type '--license-c' for details.

Encodes/decodes the input value for HTTP URL and prints
the encoded/decoded value to STDOUT.
  urlencode              // read from STDIN
  urlencode myfile.txt   // read from myfile.txt

Flags:
  -a, --all                use all input at once, instead of line-by-line
      --completion shell   generate shell completions (for "bash", "zsh", "fish", or "powershell")
  -d, --decode             decodes, instead of encodes
  -e, --encoding encoding  encode/decode format (default: "path-segment")
  -h, --help               help for urlencode
      --help-completion    help for adding shell completions
  -v, --version            version for urlencode

Valid encodings, and their intended usages:
                         http://user:pass@site.com/index.html?q=value#Frag
  -e s, -e path-segment  --------------------------index.html-------------
  -e p, -e path          -------------------------/index.html-------------
  -e q, -e query         ------------------------------------?q=value-----
  -e h, -e host          -----------------site.com------------------------
  -e c, -e cred          -------user:pass@--------------------------------
  -e f, -e frag          --------------------------------------------#Frag

                         http://[::1%25eth0]/home/index.html
  -e z, -e zone          --------------eth0-----------------
```

## License

Written and maintained by [@jilleJr](https://github.com/jilleJr).
Licensed under the GNU GPL 3.0 or later, or the CC0 1.0, depending on the file.

This repository is [REUSE](https://reuse.software/) compliant.
