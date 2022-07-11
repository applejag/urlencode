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
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var progNameColor = color.New(color.FgGreen)
var progArgColor = color.New(color.FgMagenta)

var flagNameColor = color.New(color.FgCyan)
var flagValueColor = color.New(color.FgYellow)

var commentColor = color.New(color.FgHiBlack)

type encodingFieldHelp struct {
	short  string
	long   string
	substr string
}

func writeRow(sb *strings.Builder, example string, fields []encodingFieldHelp) {
	const dashes = "--------------------------------------------------------------------------"
	exampleDashes := dashes[:len(example)]
	for _, f := range fields {
		sb.WriteString("  ")
		flagNameColor.Fprint(sb, "-e")
		sb.WriteByte(' ')
		flagValueColor.Fprint(sb, f.short)
		sb.WriteString(", ")
		flagNameColor.Fprint(sb, "-e")
		sb.WriteByte(' ')
		flagValueColor.Fprint(sb, f.long)
		sb.WriteString("              "[len(f.long):])
		i := strings.Index(example, f.substr)
		commentColor.Fprint(sb, exampleDashes[:i])
		sb.WriteString(f.substr)
		commentColor.Fprint(sb, exampleDashes[i+len(f.substr):])
		sb.WriteByte('\n')
	}
}

func encodingsMessage() string {
	var sb strings.Builder
	sb.WriteString("Valid encodings, and their intended usages:\n")
	sb.WriteString("                         ")
	const example1 = "http://user:pass@site.com/index.html?q=value#Frag"
	commentColor.Fprint(&sb, example1)
	sb.WriteByte('\n')

	writeRow(&sb, example1, []encodingFieldHelp{
		{short: "s", long: "path-segment", substr: "index.html"},
		{short: "p", long: "path", substr: "/index.html"},
		{short: "q", long: "query", substr: "?q=value"},
		{short: "h", long: "host", substr: "site.com"},
		{short: "c", long: "cred", substr: "user:pass@"},
		{short: "f", long: "frag", substr: "#Frag"},
	})

	sb.WriteString("\n                         ")
	const example2 = "http://[::1%25eth0]/home/index.html"
	commentColor.Fprint(&sb, any(example2))
	sb.WriteByte('\n')

	writeRow(&sb, example2, []encodingFieldHelp{
		{short: "z", long: "zone", substr: "eth0"},
	})

	return sb.String()
}

func flagsMessage(c *cobra.Command) string {
	var sb strings.Builder
	sb.WriteString("Flags:\n")

	c.Flags().VisitAll(func(flag *pflag.Flag) {
		if flag.Hidden {
			return
		}
		sb.WriteString("  ")
		var width int
		if flag.Shorthand != "" {
			flagNameColor.Fprintf(&sb, "-%s", flag.Shorthand)
			sb.WriteString(", ")
			width += 3 + len(flag.Shorthand)
		} else {
			sb.WriteString("    ")
			width += 4
		}
		flagNameColor.Fprintf(&sb, "--%s", flag.Name)
		width += 2 + len(flag.Name)
		t := flag.Value.Type()
		if t == "string" {
			sb.WriteByte(' ')
			flagValueColor.Fprintf(&sb, "string")
			width += 7
		}
		const spaces = "                         "
		sb.WriteString(spaces[width:])
		sb.WriteString(flag.Usage)
		if flag.DefValue != "" && t == "string" {
			sb.WriteByte(' ')
			commentColor.Fprint(&sb, "(default: ")
			flagValueColor.Fprintf(&sb, `"%s"`, flag.DefValue)
			commentColor.Fprint(&sb, ")")
		}
		sb.WriteByte('\n')
	})

	return sb.String()
}

func sampleUsageMessage() string {
	var sb strings.Builder
	sb.WriteString(`
Encodes/decodes the input value for HTTP URL and prints
the encoded/decoded value to STDOUT.
`)

	sb.WriteString("  ")
	progNameColor.Fprint(&sb, os.Args[0])
	sb.WriteString("              ")
	commentColor.Fprint(&sb, "// read from STDIN")
	sb.WriteString("\n  ")
	progNameColor.Fprint(&sb, os.Args[0])
	sb.WriteByte(' ')
	progArgColor.Fprint(&sb, "myfile.txt")
	sb.WriteString("   ")
	commentColor.Fprint(&sb, "// read from myfile.txt")
	sb.WriteRune('\n')
	return sb.String()
}

func completionHelp() string {
	return fmt.Sprintf(`Bash:

  $ source <(%[1]s --completion=bash)

  # To load completions for each session, execute once:
  # Linux:
  $ %[1]s --completion=bash > /etc/bash_completion.d/%[2]s
  # macOS:
  $ %[1]s --completion=bash > $(brew --prefix)/etc/bash_completion.d/%[2]s

Zsh:

  # If shell completion is not already enabled in your environment,
  # you will need to enable it.  You can execute the following once:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  $ %[1]s --completion=zsh > "${fpath[1]}/_%[2]s"

  # You will need to start a new shell for this setup to take effect.

fish:

  $ %[1]s --completion=fish | source

  # To load completions for each session, execute once:
  $ %[1]s --completion=fish > ~/.config/fish/completions/%[2]s.fish

PowerShell:

  PS> %[1]s --completion=powershell | Out-String | Invoke-Expression

  # To load completions for every new session, run:
  PS> %[1]s --completion=powershell > %[2]s.ps1
  # and source this file from your PowerShell profile.`, os.Args[0], filepath.Base(os.Args[0]))
}
