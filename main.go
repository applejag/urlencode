package main

import (
	"fmt"
	"os"
	"bufio"

	"github.com/spf13/pflag"
	"github.com/fatih/color"
	"github.com/mattn/go-colorable"
)

var flags struct {
	Query bool
	Decode bool
	ShowHelp bool
}

var stdout = colorable.NewColorableStdout()
var stderr = colorable.NewColorableStderr()

var errColor = color.New(color.FgHiRed)

func main() {
	pflag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage of %s: [-qd] [values...]

Encodes the input value for HTTP URL by default and prints
the encoded value to STDOUT.

Input is taken from the given arguments and prints the results
one per line, or uses each line from STDIN if no args are supplied.

Flags:
`, os.Args[0])
		pflag.PrintDefaults()
	}

	pflag.BoolVarP(&flags.Query, "query", "q", false, "encode/decode value as query parameter value")
	pflag.BoolVarP(&flags.Decode, "decode", "d", false, "decodes, instead of encodes")
	pflag.BoolVarP(&flags.ShowHelp, "help", "h", false, "show this help text")

	pflag.Parse()

	if flags.ShowHelp {
		pflag.Usage()
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
	values []string
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
