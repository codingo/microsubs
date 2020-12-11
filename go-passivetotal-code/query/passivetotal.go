package query

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type passivetotalObject struct {
	Subdomains []string `json:"subdomains"`
}

var passivetotalData passivetotalObject

// all subdomains found
var subdomains []string

// Query function returns all subdomains found using the service.
func Query(domain string, key string, silent bool, serviceName string, username string) []string {

	if username == "" || key == "" {
		return subdomains
	}

	// Create JSON Get body
	var request = []byte(`{"query":"` + domain + `"}`)

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.passivetotal.org/v2/enrichment/subdomains", bytes.NewBuffer(request))
	if err != nil {
		if !silent {
			fmt.Printf("\npassivetotal: %v\n", err)
		}
		return subdomains
	}

	req.SetBasicAuth(username, key)

	// Set content type as application/json
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		if !silent {
			fmt.Printf("\npassivetotal: %v\n", err)
		}
		return subdomains
	}

	// Get the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		if !silent {
			fmt.Printf("\npassivetotal: %v\n", err)
		}
		return subdomains
	}

	// Decode the json format
	err = json.Unmarshal([]byte(body), &passivetotalData)
	if err != nil {
		if !silent {
			fmt.Printf("\npassivetotal: %v\n", err)
		}
		return subdomains
	}

	// Append each subdomain found to subdomains array
	for _, subdomain := range passivetotalData.Subdomains {
		finalSubdomain := subdomain + "." + domain

		if silent {
			fmt.Printf("\n%s", finalSubdomain)
		} else {
			fmt.Printf("\n[%s] %s", serviceName, finalSubdomain)
		}

		subdomains = append(subdomains, finalSubdomain)
	}

	return subdomains
}
