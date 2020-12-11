package main

import (
	"flag"
	"log"
	"shodan/config"
	"shodan/helper"
	"shodan/query"
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

	key := helper.GetOldestKey(&configuration)

	log.Println("Api key using : ", key)

	domains := query.Query(*sourceDomain, key, *silent, configuration.Service, configuration.MaxPages)
	if *outputFile != "" {
		helper.WriteDomainsToFile(*outputFile, domains)
	}
}
