package main

import (
	"flag"
	"fmt"
	"github.com/eiri/mchain/mchain"
	"os"
)

func main() {
	var fileName string
	var count int
	flag.StringVar(&fileName, "input", "", "path to source text")
	flag.IntVar(&count, "count", 25, "number of words to generate")
	flag.Parse()

	seen := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
	if !seen["input"] {
		fmt.Fprintln(os.Stderr, "Error: flag -input required")
		flag.PrintDefaults()
		os.Exit(2)
	}

	m := mchain.NewMarkovChain(fileName)
	text := m.Generate(count)
	fmt.Println(text)
}
