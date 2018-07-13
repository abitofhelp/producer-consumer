////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a BSD-style  license that can be found in the LICENSE file.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Package strings implements helpful utility functions for strings.
package strings

import (
	"runtime"
	"strings"
)

// Adjust platform specific encodings for a string.
// When using STDIN on various platforms the end of line can be different.
// An example is that Windows lines end with "\r\n", which is known as CRLF.
// On UNIX, the lines end with "\n", which is known as LF.
func CleanStringForPlatform(str string) string {
	if str == "" {
		return str
	}

	if runtime.GOOS == "windows" {
		str = strings.TrimRight(str, "\r\n")
	} else {
		str = strings.TrimRight(str, "\n")
	}
	return str
}
