package simpleinterface_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/naoyafurudono/simpleinterface"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, simpleinterface.Analyzer, "a")
}
