// Copyright 2026 The Flux Authors
// SPDX-License-Identifier: Apache-2.0

package junitxml

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	. "github.com/onsi/gomega"

	apiv1 "github.com/fluxcd/flux-schema/api/v1beta1"
)

func TestFromReport(t *testing.T) {
	g := NewWithT(t)

	var report apiv1.Report
	wantBytes, err := os.ReadFile("testdata/report.json")
	g.Expect(err).To(Succeed())
	g.Expect(json.Unmarshal(wantBytes, &report)).To(Succeed())

	gotTS := FromReport(report)

	var wantTS TestSuites
	wantBytes, err = os.ReadFile("testdata/report.junit.xml")
	g.Expect(err).To(Succeed())
	g.Expect(xml.Unmarshal(wantBytes, &wantTS)).To(Succeed())

	g.Expect(cmp.Diff(
		wantTS,
		gotTS,
		cmpopts.IgnoreFields(TestSuites{}, "XMLName"),
	)).To(BeEmpty())

	var buf bytes.Buffer
	g.Expect(gotTS.Write(&buf)).To(Succeed())
	g.Expect(cmp.Diff(wantBytes, buf.Bytes())).To(BeEmpty())
}
