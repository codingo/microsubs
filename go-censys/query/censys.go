package query

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"censys/helper"
)

// all subdomains found
var subdomains []string

type resultsq struct {
	Data  []string `json:"parsed.extensions.subject_alt_name.dns_names"`
	Data1 []string `json:"parsed.names"`
}

type response struct {
	Results  []resultsq `json:"results"`
	Metadata struct {
		Pages int `json:"pages"`
	} `json:"metadata"`
}

// Query function returns all subdomains found using the service.
func Query(domain string, username string, key string, censysPages string, silent bool, serviceName string) []string {

	var uniqueSubdomains []string
	var initialSubdomains []string
	var hostResponse response

	// Default Censys Pages to process. I think 10 is a good value
	//DefaultCensysPages := 10

	if username == "" || key == "" {
		return subdomains
	}

	if censysPages != "all" {

		CensysPages, _ := strconv.Atoi(censysPages)

		for i := 1; i <= CensysPages; i++ {
			// Create JSON Get body
			var request = []byte(`{"query":"` + domain + `", "page":` + strconv.Itoa(i) + `, "fields":["parsed.names","parsed.extensions.subject_alt_name.dns_names"], "flatten":true}`)

			client := &http.Client{}
			req, err := http.NewRequest("POST", "https://www.censys.io/api/v1/search/certificates", bytes.NewBuffer(request))
			if err != nil {
				if !silent {
					fmt.Printf("\n%v: %v\n", serviceName, err)
				}
				return subdomains
			}

			req.SetBasicAuth(username, key)

			// Set content type as application/json
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Accept", "application/json")

			resp, err := client.Do(req)
			if err != nil {
				if !silent {
					fmt.Printf("\n%v: %v\n", serviceName, err)
				}
				return subdomains
			}

			// Get the response body
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				if !silent {
					fmt.Printf("\n%v: %v\n", serviceName, err)
				}
				return subdomains
			}

			err = json.Unmarshal([]byte(body), &hostResponse)
			if err != nil {
				if !silent {
					fmt.Printf("\n%v: %v\n", serviceName, err)
				}
				return subdomains
			}

			// Add all items found
			for _, res := range hostResponse.Results {
				initialSubdomains = append(initialSubdomains, res.Data...)
				initialSubdomains = append(initialSubdomains, res.Data1...)
			}

			validSubdomains := helper.Validate(domain, initialSubdomains)
			uniqueSubdomains = helper.Unique(validSubdomains)
		}

		// Append each subdomain found to subdomains array
		for _, subdomain := range uniqueSubdomains {

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
	} else if censysPages == "all" {

		// Create JSON Get body
		var request = []byte(`{"query":"` + domain + `", "page":1, "fields":["parsed.names","parsed.extensions.subject_alt_name.dns_names"], "flatten":true}`)

		client := &http.Client{}
		req, err := http.NewRequest("POST", "https://www.censys.io/api/v1/search/certificates", bytes.NewBuffer(request))
		if err != nil {
			if !silent {
				fmt.Printf("\n%v: %v\n", serviceName, err)
			}
			return subdomains
		}

		req.SetBasicAuth(username, key)

		// Set content type as application/json
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			if !silent {
				fmt.Printf("\n%v: %v\n", serviceName, err)
			}
			return subdomains
		}

		// Get the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			if !silent {
				fmt.Printf("\n%v: %v\n", serviceName, err)
			}

			return subdomains
		}

		err = json.Unmarshal([]byte(body), &hostResponse)
		if err != nil {
			if !silent {
				fmt.Printf("\n%v: %v\n", serviceName, err)
			}
			return subdomains
		}

		// Add all items found
		for _, res := range hostResponse.Results {
			initialSubdomains = append(initialSubdomains, res.Data...)
			initialSubdomains = append(initialSubdomains, res.Data1...)
		}

		validSubdomains := helper.Validate(domain, initialSubdomains)
		uniqueSubdomains = helper.Unique(validSubdomains)

		// Append each subdomain found to subdomains array
		for _, subdomain := range uniqueSubdomains {

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

		for i := 2; i <= hostResponse.Metadata.Pages; i++ {
			// Create JSON Get body
			var request = []byte(`{"query":"` + domain + `", "page":` + strconv.Itoa(i) + `, "fields":["parsed.names","parsed.extensions.subject_alt_name.dns_names"], "flatten":true}`)

			client := &http.Client{}
			req, err := http.NewRequest("POST", "https://www.censys.io/api/v1/search/certificates", bytes.NewBuffer(request))
			if err != nil {
				if !silent {
					fmt.Printf("\n%v: %v\n", serviceName, err)
				}
				return subdomains
			}

			req.SetBasicAuth(username, key)

			// Set content type as application/json
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Accept", "application/json")

			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}

			// Get the response body
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				if !silent {
					fmt.Printf("\n%v: %v\n", serviceName, err)
				}
				return subdomains
			}

			err = json.Unmarshal([]byte(body), &hostResponse)
			if err != nil {
				if !silent {
					fmt.Printf("\n%v: %v\n", serviceName, err)
				}
				return subdomains
			}

			// Add all items found
			for _, res := range hostResponse.Results {
				initialSubdomains = append(initialSubdomains, res.Data...)
				initialSubdomains = append(initialSubdomains, res.Data1...)
			}

			validSubdomains := helper.Validate(domain, initialSubdomains)
			uniqueSubdomains = helper.Unique(validSubdomains)

			// Append each subdomain found to subdomains array
			for _, subdomain := range uniqueSubdomains {

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
		}
	}

	return subdomains
}
