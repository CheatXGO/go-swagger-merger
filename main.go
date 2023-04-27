package main

import (
	"flag"

	"github.com/ntauth/go-swagger-merger/helpers"
)

func main() {
	merger := helpers.NewMerger()

	var outputFileName string

	flag.StringVar(&outputFileName, "o", "apis.swagger.json", "")
	flag.Parse()

	for _, f := range flag.Args() {
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
