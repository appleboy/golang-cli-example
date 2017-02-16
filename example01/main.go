package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var showVersion bool

	flag.BoolVar(&showVersion, "version", false, "Print version information.")
	flag.BoolVar(&showVersion, "v", false, "Print version information.")
	flag.Parse()

	// Show version and exit
	if showVersion {
		fmt.Println("Version 1.0.0")
		os.Exit(0)
	}
}
