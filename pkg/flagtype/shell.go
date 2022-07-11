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

package flagtype

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

type Shell string

const (
	ShellBash       Shell = "bash"
	ShellZsh        Shell = "zsh"
	ShellFish       Shell = "fish"
	ShellPowerShell Shell = "powershell"
)

// String is used both by fmt.Print and by Cobra in help text
func (s *Shell) String() string {
	return string(*s)
}

// Set must have pointer receiver so it doesn't change the value of a copy
func (s *Shell) Set(v string) error {
	switch strings.ToLower(v) {
	case "bash":
		*s = ShellBash
	case "zsh":
		*s = ShellZsh
	case "fish":
		*s = ShellFish
	case "powershell", "pwsh":
		*s = ShellPowerShell
	default:
		return fmt.Errorf(`invalid shell: %q, must be one of "bash", "zsh", "fish", or "powershell"`, v)
	}
	return nil
}

// Type is only used in help text
func (s *Shell) Type() string {
	return "shell"
}

func CompleteShell(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return []string{
		"bash\tCompletions for Bourne-again shell",
		"zsh\tCompletions for Z-shell",
		"fish\tCompletions for Fish shell",
		"powershell\tCompletions for Microsoft PowerShell",
		"pwsh\tCompletions for Microsoft PowerShell",
	}, cobra.ShellCompDirectiveNoFileComp
}
