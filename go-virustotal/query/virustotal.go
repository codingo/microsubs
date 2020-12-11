package query

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"virustotal/helper"
)

// all subdomains found
var subdomains []string

type virustotalapiObject struct {
	Subdomains []string `json:"subdomains"`
}

var virustotalapiData virustotalapiObject

// Local function to query virustotal API
// Requires an API key
func queryVirustotalAPI(domain string, key string, silent bool, serviceName string) (subdomains []string, err error) {

	// Make a search for a domain name and get HTTP Response
	resp, err := helper.GetHTTPResponse("https://www.virustotal.com/vtapi/v2/domain/report?apikey="+key+"&domain="+domain, 180)
	if err != nil {
		return subdomains, err
	}

	// Get the response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return subdomains, err
	}

	// Decode the json format
	err = json.Unmarshal([]byte(respBody), &virustotalapiData)
	if err != nil {
		return subdomains, err
	}

	// Append each subdomain found to subdomains array
	for _, subdomain := range virustotalapiData.Subdomains {

		// Fix Wildcard subdomains containing asterisk before them
		if strings.Contains(subdomain, "*.") {
			subdomain = strings.Split(subdomain, "*.")[1]
		}

		if silent {
			fmt.Printf("\n%s", subdomain)
		} else {
			fmt.Printf("\n[%s] %s", serviceName, subdomain)
		}

		subdomains = append(subdomains, subdomain)
	}

	return subdomains, nil
}

// Query function returns all subdomains found using the service.
func Query(domain string, key string, silent bool, serviceName string) []string {

	var subdomains []string

	if key == "" {
		return subdomains
	}
	// Get subdomains via API
	subdomains, err := queryVirustotalAPI(domain, key, silent, serviceName)
	if err != nil {
		if !silent {
			fmt.Printf("\nvirustotal: %v\n", err)
		}
		return subdomains
	}
	return subdomains
}
