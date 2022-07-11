// SPDX-FileCopyrightText: 2021 Kalle Fagerberg
//
// SPDX-License-Identifier: GPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
	"github.com/jilleJr/urlencode/pkg/flagtype"
	"github.com/jilleJr/urlencode/pkg/license"
	"github.com/mattn/go-colorable"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const version = "v1.1.0"

var versionText = fmt.Sprintf(`urlencode %s  Copyright (C) 2021  Kalle Jillheden

  License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>.
  This program comes with ABSOLUTELY NO WARRANTY; for details type '--license-w'
  This is free software, and you are welcome to redistribute it
  under certain conditions; type '--license-c' for details.`, version)

var flags = struct {
	Encode                flagtype.Encoding
	Decode                bool
	AllLines              bool
	ShowLicenseWarranty   bool
	ShowLicenseConditions bool
	Completions           flagtype.Shell
	ShowCompletionsHelp   bool
}{
	Encode: flagtype.EncodePathSegment,
}

var (
	stdout = colorable.NewColorableStdout()
	stderr = colorable.NewColorableStderr()

	errProgramNameColor    = color.New(color.FgRed, color.Italic)
	errColor               = color.New(color.FgHiRed, color.Bold)
	errUseHelpFlagTipColor = color.New(color.FgHiBlack, color.Italic)
)

var rootCmd = &cobra.Command{
	Use:   "urlencode",
	Short: "Encodes/decodes the input value for HTTP URLs",
	Long: `Encodes/decodes the input value for HTTP URLs
and prints the encoded/decoded value to STDOUT.`,
	Args:    cobra.MaximumNArgs(1),
	Version: versionText,
	Run: func(cmd *cobra.Command, args []string) {
		if flags.ShowLicenseConditions {
			fmt.Println(license.Conditions)
			return
		}

		if flags.ShowLicenseWarranty {
			fmt.Println(license.Warranty)
			return
		}

		if flags.ShowCompletionsHelp {
			fmt.Println(completionHelp())
			return
		}

		if flags.Completions != "" {
			switch flags.Completions {
			case flagtype.ShellBash:
				cmd.GenBashCompletionV2(os.Stdout, true)
			case flagtype.ShellZsh:
				cmd.GenZshCompletion(os.Stdout)
			case flagtype.ShellFish:
				cmd.GenFishCompletion(os.Stdout, true)
			case flagtype.ShellPowerShell:
				cmd.GenPowerShellCompletion(os.Stdout)
			default:
				printErr(fmt.Errorf("unsupported shell: %q", flags.Completions))
				os.Exit(1)
			}
			return
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
		if flags.AllLines {
			scanner = NewReadAllScanner(reader)
		} else {
			scanner = bufio.NewScanner(reader)
		}

		for scanner.Scan() {
			value := scanner.Text()

			if flags.Decode {
				escaped, err := unescape(value, flags.Encode)
				if err != nil {
					printErr(err)
					os.Exit(2)
				}
				fmt.Fprint(stdout, escaped)
			} else {
				fmt.Fprint(stdout, escape(value, flags.Encode))
			}

			fmt.Println()
		}

		if err := scanner.Err(); err != nil {
			printErr(err)
			os.Exit(2)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		printErr(err)
		os.Exit(1)
	}
}

func init() {
	versionText := fmt.Sprintf(`urlencode %s  Copyright (C) 2021  Kalle Jillheden

  License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>.
  This program comes with ABSOLUTELY NO WARRANTY; for details type '--license-w'
  This is free software, and you are welcome to redistribute it
  under certain conditions; type '--license-c' for details.`, version)

	rootCmd.SetHelpFunc(func(c *cobra.Command, args []string) {
		fmt.Fprintln(stderr, versionText)
		fmt.Fprintln(stderr, sampleUsageMessage())
		fmt.Fprintln(stderr, flagsMessage(c))
		fmt.Fprint(stderr, encodingsMessage())
	})
	// We have our own error handling in Execute()
	rootCmd.SilenceErrors = true
	// Only print help if calling with --help
	rootCmd.SilenceUsage = true

	rootCmd.Flags().VarP(&flags.Encode, "encoding", "e", "encode/decode format")
	rootCmd.RegisterFlagCompletionFunc("encoding", flagtype.CompleteEncoding)
	rootCmd.Flags().BoolVarP(&flags.Decode, "decode", "d", false, "decodes, instead of encodes")
	rootCmd.Flags().BoolVarP(&flags.AllLines, "all", "a", false, "use all input at once, instead of line-by-line")
	rootCmd.Flags().Var(&flags.Completions, "completion", `generate shell completions (for "bash", "zsh", "fish", or "powershell")`)
	rootCmd.RegisterFlagCompletionFunc("completion", flagtype.CompleteShell)
	rootCmd.Flags().BoolVar(&flags.ShowCompletionsHelp, "help-completion", false, "help for adding shell completions")

	rootCmd.Flags().BoolVarP(&flags.ShowLicenseConditions, "license-c", "", false, "show license conditions")
	rootCmd.Flags().BoolVarP(&flags.ShowLicenseWarranty, "license-w", "", false, "show license warranty")
	rootCmd.Flags().MarkHidden("license-c")
	rootCmd.Flags().MarkHidden("license-w")
}

func printErr(err error) {
	fmt.Fprintln(stderr, errProgramNameColor.Sprint("urlencode:"), errColor.Sprint("err:"), err)
	fmt.Fprintln(stderr, errUseHelpFlagTipColor.Sprintf(`tip: Call "%s --help" to see usage`, os.Args[0]))
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
