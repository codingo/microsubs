package main

import (
	"flag"
	"log"
	"securitytrails/config"
	"securitytrails/helper"
	"securitytrails/query"
)

// Command line flags
var (
	sourceDomain = flag.String("d", "", "Target domain e.g apple.com (required)")
	outputFile   = flag.String("o", "", "Output file (optional)")
	silent       = flag.Bool("silent", false, "To display only the subdomains")
	paid         = flag.Bool("paid", false, "To use the paid keys")
	configFile   = flag.String("c", "config.json", "Configuration file. (optional)")
)

func main() {
	flag.Parse()
	if *sourceDomain == "" {
		log.Fatalln("Domain is required.")
	}
	configuration := config.ReadConfig(*configFile)

	domains := query.Query(*sourceDomain, *silent, configuration.Service, &configuration, *paid)
	if *outputFile != "" {
		helper.WriteDomainsToFile(*outputFile, domains)
	}
}
