// Copyright 2026 The Flux Authors
// SPDX-License-Identifier: Apache-2.0

package flag

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestOutput_Set(t *testing.T) {
	tests := []struct {
		name      string
		str       string
		expect    string
		expectErr bool
	}{
		{"text", "text", "text", false},
		{"yaml", "yaml", "yaml", false},
		{"json", "json", "json", false},
		{"junit", "junit", "junit", false},
		{"unsupported", "xml", "", true},
		{"empty", "", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)

			o := new(Output)
			ov := NewOutputValue(o, "text", "yaml", "json", "junit")
			err := ov.Set(tt.str)
			if tt.expectErr {
				g.Expect(err).To(HaveOccurred())
			} else {
				g.Expect(err).ToNot(HaveOccurred())
			}
			g.Expect(o.String()).To(Equal(tt.expect))
		})
	}
}
