package main

import (
	"github.com/aschwinwester/imagegrp/grep"
	"fmt"
)

func main() {


	options := grep.GetOptions()

	if options.URL == "" {
		fmt.Println("provide URL as argument")
		return
	}


	grep.Fetch(options)

}

