package main

import (
	"github.com/shamaton/custom-linter-sample/mylinter"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(mylinter.Analyzer) }
