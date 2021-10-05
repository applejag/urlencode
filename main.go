package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/mattn/go-colorable"
	"github.com/spf13/pflag"
)

const version = "v1.0.0"

var flags struct {
	Query                 bool
	Decode                bool
	ShowHelp              bool
	ShowVersion           bool
	ShowLicenseWarranty   bool
	ShowLicenseConditions bool
}

var stdout = colorable.NewColorableStdout()
var stderr = colorable.NewColorableStderr()

var errColor = color.New(color.FgHiRed)

func main() {
	versionText := fmt.Sprintf(`urlencode %s  Copyright (C) 2021  Kalle Jillheden

  License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>.
  This program comes with ABSOLUTELY NO WARRANTY; for details type '--license-w'
  This is free software, and you are welcome to redistribute it
  under certain conditions; type '--license-c' for details.`, version)

	pflag.Usage = func() {
		fmt.Fprintf(os.Stderr, `%s

Encodes the input value for HTTP URL by default and prints
the encoded value to STDOUT.

Input is taken from the given arguments and prints the results
one per line, or uses each line from STDIN if no args are supplied.

Flags:
`, versionText)
		pflag.PrintDefaults()
	}

	pflag.BoolVarP(&flags.Query, "query", "q", false, "encode/decode value as query parameter value")
	pflag.BoolVarP(&flags.Decode, "decode", "d", false, "decodes, instead of encodes")
	pflag.BoolVarP(&flags.ShowHelp, "help", "h", false, "show this help text and exit")
	pflag.BoolVar(&flags.ShowVersion, "version", false, "show version and exit")

	pflag.BoolVarP(&flags.ShowLicenseConditions, "license-c", "", false, "show license conditions")
	pflag.BoolVarP(&flags.ShowLicenseWarranty, "license-w", "", false, "show license warranty")
	pflag.CommandLine.MarkHidden("license-c")
	pflag.CommandLine.MarkHidden("license-w")

	pflag.Parse()

	if flags.ShowHelp {
		pflag.Usage()
		os.Exit(0)
	}

	if flags.ShowVersion {
		fmt.Println(versionText)
		os.Exit(0)
	}

	if flags.ShowLicenseConditions {
		fmt.Println(licenseConditions)
		os.Exit(0)
	}

	if flags.ShowLicenseWarranty {
		fmt.Println(licenseWarranty)
		os.Exit(0)
	}

	var scanner Scanner
	var enc = encodePath

	if flags.Query {
		enc = encodeQueryComponent
	}

	if pflag.NArg() == 0 {
		scanner = bufio.NewScanner(os.Stdin)
		defer os.Stdin.Close()
	} else {
		scanner = &StringScanner{
			values: pflag.Args(),
		}
	}

	for scanner.Scan() {
		value := scanner.Text()

		if flags.Decode {
			escaped, err := unescape(value, enc)
			if err != nil {
				fmt.Fprintln(stderr, errColor.Sprint("err:"), err)
				os.Exit(2)
			}
			fmt.Fprint(stdout, escaped)
		} else {
			fmt.Fprint(stdout, escape(value, enc))
		}

		fmt.Println()
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(stderr, errColor.Sprint("err:"), err)
		os.Exit(2)
	}
}

type Scanner interface {
	Scan() bool
	Text() string
	Err() error
}

type StringScanner struct {
	values    []string
	nextIndex int
}

func (ss *StringScanner) Scan() bool {
	if ss.nextIndex >= len(ss.values) {
		return false
	}
	ss.nextIndex++
	return true
}

func (ss *StringScanner) Text() string {
	return ss.values[ss.nextIndex-1]
}

func (ss *StringScanner) Err() error {
	return nil
}
