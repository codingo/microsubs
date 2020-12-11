package query

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"shodan/helper"
)

type ShodanResult struct {
	Matches []shodanObject `json:"matches"`
	Result  int            `json:"result"`
	Error   string         `json:"error"`
}

// Structure of a single dictionary of output by crt.sh
type shodanObject struct {
	Hostnames []string `json:"hostnames"`
}

var shodanResult ShodanResult

// all subdomains found
var subdomains []string

// Query function returns all subdomains found using the service.
func Query(domain string, key string, silent bool, serviceName string, maxPages int) []string {
	shodanAPIKey := key

	if shodanAPIKey == "" {
		return subdomains
	}

	for currentPage := 0; currentPage <= maxPages; currentPage++ {
		resp, err := helper.GetHTTPResponse("https://api.shodan.io/shodan/host/search?query=hostname:"+domain+"&page="+strconv.Itoa(currentPage)+"&key="+shodanAPIKey, 180)
		if err != nil {
			fmt.Printf("\nshodan: %v\n", err)
			return subdomains
		}

		// Get the response body
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("\nshodan: %v\n", err)
			return subdomains
		}

		// Decode the json format
		err = json.Unmarshal([]byte(respBody), &shodanResult)
		if err != nil {
			fmt.Printf("\nshodan: %v\n", err)
			return subdomains
		}

		if shodanResult.Error != "" {
			return subdomains
		}

		// Append each subdomain found to subdomains array
		for _, block := range shodanResult.Matches {
			for _, hostname := range block.Hostnames {

				// Fix Wildcard subdomains containg asterisk before them
				if strings.Contains(hostname, "*.") {
					hostname = strings.Split(hostname, "*.")[1]
				}

				if silent {
					fmt.Printf("\n%s", hostname)
				} else {
					fmt.Printf("\n[%s] %s", serviceName, hostname)
				}
				subdomains = append(subdomains, hostname)

				subdomains = append(subdomains, hostname)
			}
		}
	}

	return subdomains
}
