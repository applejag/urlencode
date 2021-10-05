# urlencode

Super basic URL encoding utility. I needed one, so I decided to make one.

![Screenshot from 2021-10-05 19-29-55](https://user-images.githubusercontent.com/2477952/136073443-840ee017-27b1-45c0-be23-e33bb3a43127.png)

## Installation

Requires Go v1.16 (or higher)

```console
$ go install github.com/jilleJr/urlencode
```

## Features

- Encodes

- Decodes

- Colored output to highlight what's encoded/decoded

- Read from STDIN or from a file

## Usage

```console
$ urlencode --help
urlencode v1.0.0  Copyright (C) 2021  Kalle Jillheden

  License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>.
  This program comes with ABSOLUTELY NO WARRANTY; for details type '--license-w'
  This is free software, and you are welcome to redistribute it
  under certain conditions; type '--license-c' for details.

Encodes/decodes the input value for HTTP URL by default and prints
the encoded/decoded value to STDOUT.
  urlencode             // read from STDIN
  urlencode myfile.txt  // read from myfile.txt

Flags:
  -a, --all                      use all input at once, instead of line-by-line
  -d, --decode                   decodes, instead of encodes
  -e, --encoding "path-segment"  encode/decode format
  -h, --help                     show this help text and exit
      --version                  show version and exit

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

Copyright &copy; 2021 Kalle Jillheden

License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>.
See full license text in the [LICENSE](./LICENSE) file.
