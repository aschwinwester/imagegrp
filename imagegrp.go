package main

import (
	"github.com/aschwinwester/imagegrp/grep"
	"fmt"
)

func main() {


	options := grep.GetOptions()

	if options.URL == "" {
		fmt.Println("provide URL as argument and prefix with optional flags")
		return
	}


	grep.Fetch(options)

}

