// Copyright (c) 2020-2021 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package pi

import "testing"

func TestProposalNameIsValid(t *testing.T) {
	// Setup pi plugin
	p, cleanup := newTestPiPlugin(t)
	defer cleanup()

	tests := []struct {
		name string
		want bool
	}{
		// empty test
		{
			"",
			false,
		},
		// 7 characters
		{
			"abcdefg",
			false,
		},

		// 81 characters
		{
			"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			false,
		},
		// 8 characters
		{
			"12345678",
			true,
		},
		{
			"valid title",
			true,
		},
		{
			" - title: is valid; title. !.,  ",
			true,
		},
		{
			" - title: is valid; title.   ",
			true,
		},
		{
			"\n\n#This-is MY tittle###",
			false,
		},
		{
			"{this-is-the-title}",
			false,
		},
		{
			"\t<this- is-the title>",
			false,
		},
		{
			"{this   -is-the-title}   ",
			false,
		},
		{
			"###this is the title***",
			false,
		},
		{
			"###this is the title@+",
			true,
		},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			isValid := p.proposalNameIsValid(test.name)
			if isValid != test.want {
				t.Errorf("got %v, want %v", isValid, test.want)
			}
		})
	}
}
