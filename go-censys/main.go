package main

import (
	"censys/config"
	"censys/helper"
	"censys/query"
	"flag"
	"log"
)

// Command line flags
var (
	sourceDomain = flag.String("d", "", "Target domain e.g apple.com (required)")
	outputFile   = flag.String("o", "", "Output file (optional)")
	silent       = flag.Bool("silent", false, "To display only the subdomains")
	configFile   = flag.String("c", "config.json", "Configuration file. (optional)")
)

func main() {
	flag.Parse()
	if *sourceDomain == "" {
		log.Fatalln("Domain is required.")
	}
	configuration := config.ReadConfig(*configFile)

	username, key := helper.GetOldestKey(&configuration)

	domains := query.Query(*sourceDomain, username, key, configuration.Censyspages, *silent, configuration.Service)
	if *outputFile != "" {
		helper.WriteDomainsToFile(*outputFile, domains)
	}
}
