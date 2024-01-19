package main

import (
	"cmp"
	"flag"

	"github.com/CheatXGO/go-swagger-merger/helpers"
	"golang.org/x/exp/slices"
)

func main() {
	var (
		outputFileName string
		outputTitle    string
		outputVersion  string
	)

	flag.StringVar(&outputFileName, "o", "apis.swagger.json", "")
	flag.StringVar(&outputTitle, "t", "title", "")
	flag.StringVar(&outputVersion, "v", "version", "")
	flag.Parse()

	// Sort the files lexicographically in reverse so that the swagger annotations
	// artifact always comes last. This is required so that the merged file contains
	// the annotations info.
	files := flag.Args()
	slices.SortFunc(files, func(f1, f2 string) int {
		return cmp.Compare(f1, f2)
	})

	merger := helpers.NewMerger(outputTitle, outputVersion)

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
