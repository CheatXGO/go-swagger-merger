package main

import (
	"flag"

	"github.com/ntauth/go-swagger-merger/helpers"
	"golang.org/x/exp/slices"
)

func main() {
	merger := helpers.NewMerger()

	var outputFileName string

	flag.StringVar(&outputFileName, "o", "apis.swagger.json", "")
	flag.Parse()

	// Sort the files lexicographically so that the swagger annotations
	// artifact always comes first.
	files := flag.Args()
	slices.Sort(files)

	for _, f := range files {
		err := merger.AddFile(f)
		if err != nil {
			panic(err)
		}
	}

	err := merger.Save(outputFileName)
	if err != nil {
		panic(err)
	}
}
