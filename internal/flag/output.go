// Copyright 2026 The Flux Authors
// SPDX-License-Identifier: Apache-2.0

package flag

import (
	"fmt"
	"slices"
	"strings"

	"github.com/spf13/pflag"
)

type Output string

func (o *Output) String() string {
	return string(*o)
}

func ValidateOutput(output string, allowed []string) error {
	if strings.TrimSpace(output) == "" {
		return fmt.Errorf("no output format given, must be one of: %s",
			strings.Join(allowed, ", "))
	}
	if !slices.Contains(allowed, output) {
		return fmt.Errorf("unsupported output format '%s', must be one of: %s",
			output, strings.Join(allowed, ", "))
	}
	return nil
}

type OutputValue interface {
	pflag.Value
	Description() string
}

func NewOutputValue(output *Output, allowed ...string) OutputValue {
	return &outputValue{
		output:  output,
		allowed: allowed,
	}
}

type outputValue struct {
	output  *Output
	allowed []string
}

func (o *outputValue) String() string {
	return o.output.String()
}

func (o *outputValue) Set(str string) error {
	if err := ValidateOutput(str, o.allowed); err != nil {
		return err
	}
	*o.output = Output(str)
	return nil
}

func (o *outputValue) Type() string {
	return strings.Join(o.allowed, "|")
}

func (o *outputValue) Description() string {
	return fmt.Sprintf("output format, can be one of: %s", strings.Join(o.allowed, ", "))
}
