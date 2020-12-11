package query

import (
	"fmt"
	"os"

	"dnsgrep/stringset"
	"strings"

	"dnsgrep/domainutil"

	"github.com/parnurzeal/gorequest"
)

type DNSGrep struct {
	FdnsA []string `json:"FDNS_A"`
}

// var response chan

// Query function returns all subdomains found using the service.
func Query(domain string, silent bool, subdomain bool, services []string) []string {
	subs := stringset.New()
	noResposne := 0
	if len(services) == 0 {
		return subs.Slice()
	}

	for _, service := range services {
		query := fmt.Sprintf("%s%s", service, domain)
		var grepResult DNSGrep
		_, _, err := gorequest.New().Get(query).EndStruct(&grepResult)
		if err != nil {
			if !silent {
				fmt.Printf("%v\n", err)
			}
		}
		if len(grepResult.FdnsA) == 0 {
			noResposne++
		} else {
			subs.Union(getDomains(grepResult.FdnsA, subdomain))
		}
	}
	if noResposne == len(services) {
		if !silent {
			fmt.Println("No results found!")
		}
		os.Exit(1)
	}
	return subs.Slice()
}

func getDomains(records []string, subdomain bool) stringset.Set {
	var subdomains []string
	subs := stringset.New()
	for _, record := range records {
		split := strings.Split(record, ",")
		if len(split) == 2 {
			if subdomain {
				if domainutil.HasSubdomain(split[1]) {
					subdomains = append(subdomains, split[1])
				}
			} else {
				subdomains = append(subdomains, split[1])
			}
		}
	}
	subs.InsertMany(subdomains...)
	return subs
}
