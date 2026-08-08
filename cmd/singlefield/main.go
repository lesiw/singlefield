package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"lesiw.io/singlefield"
)

func main() { singlechecker.Main(singlefield.Analyzer) }
