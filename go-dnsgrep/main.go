package main

import (
	"dnsgrep/config"
	"dnsgrep/helper"
	"dnsgrep/query"
	"flag"
	"fmt"
	"log"
)

// Command line flags
var (
	sourceDomain = flag.String("d", "", "Target domain e.g apple.com (required)")
	outputFile   = flag.String("o", "", "Output file (optional)")
	silent       = flag.Bool("silent", false, "To display only discovered data")
	subdomains   = flag.Bool("subdomains", false, "Limit results to subdomains only.")
	configFile   = flag.String("c", "config.json", "Configuration file. (optional)")
)

func main() {
	flag.Parse()
	if *sourceDomain == "" {
		log.Fatalln("Domain is required.")
	}
	configuration := config.ReadConfig(*configFile)

	domains := query.Query(*sourceDomain, *silent, *subdomains, configuration.Services)
	if *outputFile != "" {
		go helper.WriteDomainsToFile(*outputFile, domains)
	}
	for _, domain := range domains {
		fmt.Printf("%v\n", domain)
	}
}
