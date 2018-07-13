////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a BSD-style  license that can be found in the LICENSE file.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Package strings implements helpful utility functions for strings.
package strings_test

import (
	. "gitlab.com/abitofhelpinc/producer-consumer/common"
	"testing"
)

func TestCleanStringForPlatform(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CleanStringForPlatform(tt.args.str); got != tt.want {
				t.Errorf("CleanStringForPlatform() = %v, want %v", got, tt.want)
			}
		})
	}
}
