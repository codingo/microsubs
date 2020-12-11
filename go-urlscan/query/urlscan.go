package query

import (
	"encoding/json"
	"fmt"
	"strings"
	"urlscan/helper"
	"urlscan/stringset"
)

const (
	// UserAgent is the default user agent used by Amass during HTTP requests.
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.119 Safari/537.36"
)

// all subdomains found
var subdomains []string
var serviceName string
var silent bool

// Query function returns all subdomains found using the service.
func Query(domain string, key string, slnt bool, service string) []string {
	if key == "" {
		return subdomains
	}
	serviceName = service
	silent = slnt
	url := searchURL(domain)
	page, err := helper.RequestWebPage(url, nil, nil, "", "")
	if err != nil {
		if !silent {
			fmt.Printf("\n%v %v\n", serviceName, err)
		}
		return subdomains
	}
	// Extract the subdomain names from the REST API results
	var results struct {
		Results []struct {
			ID string `json:"_id"`
		} `json:"results"`
		Total int `json:"total"`
	}
	if err := json.Unmarshal([]byte(page), &results); err != nil {
		if !silent {
			fmt.Printf("\n%v %v\n", serviceName, err)
		}
		return subdomains
	}

	var ids []string
	if results.Total > 0 {
		for _, result := range results.Results {
			ids = append(ids, result.ID)
		}
	} else {
		if id := attemptSubmission(domain, key); id != "" {
			ids = []string{id}
		}
	}

	subs := stringset.New()
	for _, id := range ids {
		subs.Union(getSubsFromResult(id))
	}
	return subs.Slice()
}

func getSubsFromResult(id string) stringset.Set {
	subs := stringset.New()

	url := resultURL(id)
	page, err := helper.RequestWebPage(url, nil, nil, "", "")
	if err != nil {
		if !silent {
			fmt.Printf("\n%v %v\n", serviceName, err)
		}
		return subs
	}
	// Extract the subdomain names from the REST API results
	var data struct {
		Lists struct {
			IPs        []string `json:"ips"`
			Subdomains []string `json:"linkDomains"`
		} `json:"lists"`
	}
	if err := json.Unmarshal([]byte(page), &data); err == nil {
		subs.InsertMany(data.Lists.Subdomains...)
	}
	domains := subs.Slice()
	for _, d := range domains {
		if silent {
			fmt.Printf("\n%s", d)
		} else {
			fmt.Printf("\n[%s] %s", serviceName, d)
		}
	}
	return subs
}

func attemptSubmission(domain string, key string) string {

	headers := map[string]string{
		"API-Key":      key,
		"Content-Type": "application/json",
	}
	url := "https://urlscan.io/api/v1/scan/"
	body := strings.NewReader(submitBody(domain))
	page, err := helper.RequestWebPage(url, body, headers, "", "")
	if err != nil {
		if !silent {
			fmt.Printf("\n4%v %v\n", serviceName, err)
		}
		return ""
	}
	// Extract the subdomain names from the REST API results
	var result struct {
		Message string `json:"message"`
		ID      string `json:"uuid"`
		API     string `json:"api"`
	}
	if err := json.Unmarshal([]byte(page), &result); err != nil {
		return ""
	}
	if result.Message != "Submission successful" {
		return ""
	}
	// Keep this data source active while waiting for the scan to complete
	for {
		_, err = helper.RequestWebPage(result.API, nil, nil, "", "")
		if err == nil || err.Error() != "404 Not Found" {
			break
		}
	}
	return result.ID
}

func searchURL(domain string) string {
	return fmt.Sprintf("https://urlscan.io/api/v1/search/?q=domain:%s", domain)
}

func resultURL(id string) string {
	return fmt.Sprintf("https://urlscan.io/api/v1/result/%s/", id)
}

func submitBody(domain string) string {
	return fmt.Sprintf("{\"url\": \"%s\", \"public\": \"on\", \"customagent\": \"%s\"}", domain, UserAgent)
}
