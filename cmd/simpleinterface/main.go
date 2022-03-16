package main

import (
	"github.com/naoyafurudono/simpleinterface"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(simpleinterface.Analyzer) }
