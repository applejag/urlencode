package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
	"github.com/mattn/go-colorable"
	"github.com/spf13/pflag"
)

const version = "v1.0.0"

var flags struct {
	Encode                string
	Decode                bool
	SeparateLines         bool
	ShowHelp              bool
	ShowVersion           bool
	ShowLicenseWarranty   bool
	ShowLicenseConditions bool
}

var stdout = colorable.NewColorableStdout()
var stderr = colorable.NewColorableStderr()

var errProgramNameColor = color.New(color.FgRed, color.Italic)
var errColor = color.New(color.FgHiRed, color.Bold)

func main() {
	versionText := fmt.Sprintf(`urlencode %s  Copyright (C) 2021  Kalle Jillheden

  License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>.
  This program comes with ABSOLUTELY NO WARRANTY; for details type '--license-w'
  This is free software, and you are welcome to redistribute it
  under certain conditions; type '--license-c' for details.`, version)

	pflag.Usage = func() {
		fmt.Fprintln(stderr, versionText)
		fmt.Fprintln(stderr, sampleUsageMessage())
		fmt.Fprintln(stderr, flagsMessage())
		fmt.Fprint(stderr, encodingsMessage())
	}

	pflag.StringVarP(&flags.Encode, "encoding", "e", "path-segment", "encode/decode format")
	pflag.BoolVarP(&flags.Decode, "decode", "d", false, "decodes, instead of encodes")
	pflag.BoolVarP(&flags.SeparateLines, "lines", "l", false, "encode/decode each line by themselves")
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

	var enc encoding
	switch flags.Encode {
	case "s", "path-segment":
		enc = encodePathSegment
	case "p", "path":
		enc = encodePath
	case "q", "query":
		enc = encodeQueryComponent
	case "h", "host":
		enc = encodeHost
	case "z", "zone":
		enc = encodeZone
	case "c", "cred":
		enc = encodeUserPassword
	case "f", "frag":
		enc = encodeFragment
	default:
		printErr(fmt.Errorf("invalid encoding: %q", flags.Encode))
		fmt.Fprint(stdout, encodingsMessage())
		os.Exit(1)
	}

	if pflag.NArg() > 1 {
		printErr(errors.New("must only supply up to one file name argument"))
		os.Exit(1)
	}

	var reader io.Reader
	if pflag.NArg() == 0 {
		reader = os.Stdin
		defer os.Stdin.Close()
	} else {
		filename := pflag.Arg(0)
		file, err := os.Open(filename)
		if err != nil {
			printErr(err)
			os.Exit(3)
		}
		reader = file
		defer file.Close()
	}

	var scanner Scanner
	if flags.SeparateLines {
		scanner = bufio.NewScanner(reader)
	} else {
		scanner = NewReadAllScanner(reader)
	}

	for scanner.Scan() {
		value := scanner.Text()

		if flags.Decode {
			escaped, err := unescape(value, enc)
			if err != nil {
				printErr(err)
				os.Exit(2)
			}
			fmt.Fprint(stdout, escaped)
		} else {
			fmt.Fprint(stdout, escape(value, enc))
		}

		fmt.Println()
	}

	if err := scanner.Err(); err != nil {
		printErr(err)
		os.Exit(2)
	}
}

func printErr(err error) {
	fmt.Fprintln(stderr, errProgramNameColor.Sprintf("%s:", os.Args[0]), errColor.Sprint("err:"), err)
}

type Scanner interface {
	Scan() bool
	Text() string
	Err() error
}

type readAllScanner struct {
	reader io.Reader
	bytes  []byte
	err    error
}

func NewReadAllScanner(reader io.Reader) Scanner {
	return &readAllScanner{
		reader: reader,
	}
}

func (s *readAllScanner) Scan() bool {
	if s.bytes != nil {
		return false
	}

	s.bytes, s.err = io.ReadAll(s.reader)
	return s.err == nil
}

func (s *readAllScanner) Text() string {
	return string(s.bytes)
}

func (s *readAllScanner) Err() error {
	return s.err
}
