package main

import (
	"simpleinterface"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(simpleinterface.Analyzer) }
