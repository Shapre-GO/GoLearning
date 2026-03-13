package loglinter_test

import (
	"testing"

	"loglinter/analyzer"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestLogLinter(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, analyzer.Analyzer, "sample")
}
