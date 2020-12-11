// Package dnsdumpster is a Parser for subdomains from DNSDumpster
package query

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"

	"dnsdumpster/helper"
)

// all subdomains found
var subdomains []string
var gCookies []*http.Cookie

// Query function returns all subdomains found using the service.
func Query(domain string, silent bool, serviceName string) []string {

	// CookieJar to hold csrf cookie
	var curCookieJar *cookiejar.Jar
	curCookieJar, _ = cookiejar.New(nil)

	// Make a http request to DNSDumpster
	resp, gCookies, err := helper.GetHTTPCookieResponse("https://dnsdumpster.com", gCookies, 180)
	if err != nil {
		if !silent {
			fmt.Printf("\ndnsdumpster: %v\n", err)
		}
		return subdomains
	}

	// Get the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		if !silent {
			fmt.Printf("\ndnsdumpster: %v\n", err)
		}
		return subdomains
	}

	src := string(body)

	re := regexp.MustCompile("<input type='hidden' name='csrfmiddlewaretoken' value='(.*)' />")
	match := re.FindAllStringSubmatch(src, -1)

	// CSRF Middleware token for POST Request
	csrfmiddlewaretoken := match[0]

	// Set cookiejar values
	u, _ := url.Parse("https://dnsdumpster.com")
	curCookieJar.SetCookies(u, gCookies)

	hc := http.Client{Jar: curCookieJar}
	form := url.Values{}

	form.Add("csrfmiddlewaretoken", csrfmiddlewaretoken[1])
	form.Add("targetip", domain)

	// Create a post request to get subdomain data
	req, err := http.NewRequest("POST", "https://dnsdumpster.com", strings.NewReader(form.Encode()))
	if err != nil {
		if !silent {
			fmt.Printf("\ndnsdumpster: %v\n", err)
		}
		return subdomains
	}

	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Referer", "https://dnsdumpster.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.0.1) Gecko/2008071615 Fedora/3.0.1-1.fc9 Firefox/3.0.1")

	resp, err = hc.Do(req)
	if err != nil {
		if !silent {
			fmt.Printf("\ndnsdumpster: %v\n", err)
		}
		return subdomains
	}

	// Get the response body
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		if !silent {
			fmt.Printf("\ndnsdumpster: %v\n", err)
		}
		return subdomains
	}

	src = string(body)

	// Find the table holding host records
	Regex, _ := regexp.Compile("<td class=\"col-md-4\">(.*\\..*\\..*)<br>")
	match = Regex.FindAllStringSubmatch(src, -1)

	// String to hold initial subdomains
	var initialSubs []string

	for _, data := range match {
		initialSubs = append(initialSubs, data[1])
	}

	validSubdomains := helper.Validate(domain, initialSubs)

	for _, subdomain := range validSubdomains {
		if silent {
			fmt.Printf("\n%s", subdomain)
		} else {
			fmt.Printf("\n[%s] %s", serviceName, subdomain)
		}

		subdomains = append(subdomains, subdomain)
	}

	return subdomains
}
