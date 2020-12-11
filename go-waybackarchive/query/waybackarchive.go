package query

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"


	"waybackarchive/helper"
)

// all subdomains found
var subdomains []string

// Query function returns all subdomains found using the service.
func Query(domain string, silent bool, serviceName string) []string {


	pagesResp, err := helper.GetHTTPResponse("http://web.archive.org/cdx/search/cdx?url=*."+domain+"&showNumPages=true", 180)
	if err != nil {
		if !silent {
			fmt.Printf("\n%v: %v\n",serviceName, err)
		}
		return subdomains
	}

	b, err := ioutil.ReadAll(pagesResp.Body)
	if err != nil {
		if !silent {
			fmt.Printf("\n%v: %v\n",serviceName, err)
		}
		return subdomains
	}

	numPages, err := strconv.Atoi(strings.Split(string(b), "\n")[0])
	if err != nil {
		if !silent {
			fmt.Printf("\n%v: %v\n",serviceName, err)
		}
		return subdomains
	}

	for i := 0; i <= numPages; i++ {
		resp, err := helper.GetHTTPResponse("http://web.archive.org/cdx/search/cdx?url=*."+domain+"/*&output=json&fl=original&collapse=urlkey&page="+string(i), 180)
		if err != nil {
			if !silent {
				fmt.Printf("\n%v: %v\n",serviceName, err)
			}
			return subdomains
		}

		// Get the response body
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			if !silent {
				fmt.Printf("\n%v: %v\n",serviceName, err)
			}
			return subdomains
		}

		initialSubs := helper.ExtractSubdomainsFromText(string(respBody), domain)
		validSubdomains := helper.Unique(initialSubs)

		for _, subdomain := range validSubdomains {
			if helper.SubdomainExists(subdomain, subdomains) == false {
				if silent {
					fmt.Printf("\n%s", subdomain)
				} else {
					fmt.Printf("\n[%s] %s", serviceName, subdomain)
				}
				subdomains = append(subdomains, subdomain)
			}
		}
	}

	return subdomains
}
