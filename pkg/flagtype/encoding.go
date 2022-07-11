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

	"github.com/spf13/cobra"
)

type Encoding string

const (
	EncodePathSegment    Encoding = "path-segment"
	EncodePath           Encoding = "path"
	EncodeQueryComponent Encoding = "query"
	EncodeHost           Encoding = "host"
	EncodeZone           Encoding = "zone"
	EncodeUserPassword   Encoding = "cred"
	EncodeFragment       Encoding = "frag"
)

// String is used both by fmt.Print and by Cobra in help text
func (e *Encoding) String() string {
	return string(*e)
}

// Set must have pointer receiver so it doesn't change the value of a copy
func (e *Encoding) Set(v string) error {
	switch v {
	case "s", "path-segment":
		*e = EncodePathSegment
	case "p", "path":
		*e = EncodePath
	case "q", "query":
		*e = EncodeQueryComponent
	case "h", "host":
		*e = EncodeHost
	case "z", "zone":
		*e = EncodeZone
	case "c", "cred":
		*e = EncodeUserPassword
	case "f", "frag":
		*e = EncodeFragment
	default:
		return fmt.Errorf("invalid encoding: %q", v)
	}
	return nil
}

// Type is only used in help text
func (e *Encoding) Type() string {
	return "encoding"
}

func CompleteEncoding(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return []string{
		"path-segment\tSegment between two / slashes /",
		"s\tSegment between two / slashes /",
		"path\tAll path segments, including the slashes /",
		"p\tAll path segments, including the slashes /",
		"query\tQuery parameter (key or value), e.g ?key=value",
		"q\tQuery parameter (key or value), e.g ?key=value",
		"host\tHostname (FQDN)",
		"h\tHostname (FQDN)",
		"cred\tCredentials (username:password@)",
		"c\tCredentials (username:password@)",
		"frag\tFragment parameter, everything past the hash #",
		"f\tFragment parameter, everything past the hash #",
		"zone\tIPv6 zone parameter",
		"z\tIPv6 zone parameter",
	}, cobra.ShellCompDirectiveNoFileComp
}
