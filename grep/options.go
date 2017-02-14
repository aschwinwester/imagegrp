package grep

import (
	"strings"
	"flag"
	"log"
)

type Options struct {
	Verbose bool
	URL string
	Folder string
	CSSClass string
}

func GetOptions() Options {


	var options Options
	var verbose= flag.Bool("v", false, "Verbose logging")
	var folderName = flag.String("f", "", "Folder name")
	var css = flag.String("cssClass", "", "css class value")
	flag.Parse()

	options.Folder = *folderName
	options.Verbose = *verbose
	options.CSSClass = *css

	if (options.Verbose) {
		log.Printf("Found folder %s", options.Folder)
		log.Printf("CSS class %s", options.CSSClass)
	}
	for _, arg := range flag.Args() {
		if (options.Verbose) {
			log.Printf("found arg %s", arg)
		}

		if strings.HasPrefix(arg, "http") {
			options.URL = arg
		}
	}




	return options
}

