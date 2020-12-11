
// Written By : @ice3man (Nizamul Rana)
//
// Distributed Under MIT License
// Copyrights (C) 2018 Ice3man
//

// Package securitytrails is a golang SecurityTrails API client for subdomain discovery.
package query

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
	"securitytrails/config"
	"securitytrails/helper"
)

type securitytrailsObject struct {
	Subdomains []string `json:"subdomains"`
}

var securitytrailsData securitytrailsObject

// all subdomains found
var subdomains []string

// Query function returns all subdomains found using the service.
func Query(domain string, silent bool,serviceName string,cnfg *config.Configuration,paid bool) []string {
	flag:=true
	tries:=0
	for(flag){

		key,totalKeys:= helper.GetOldestKey(cnfg, paid)
		tries=tries+1
		if(tries==totalKeys){
			panic("All keys are exhausted.")
		}
		log.Println("Api key using : ", key)

		// Get credentials for performing HTTP Basic Auth
		securitytrailsKey := key

		if securitytrailsKey == "" {
			return subdomains
		}

		client := &http.Client{}
		req, err := http.NewRequest("GET", "https://api.securitytrails.com/v1/domain/"+domain+"/subdomains", nil)
		if err != nil {
			fmt.Printf("\npassivetotal: %v\n", err)
			return subdomains
		}

		req.Header.Add("APIKEY", securitytrailsKey)

		resp, err := client.Do(req)
		if(resp.StatusCode==429){
			fmt.Printf("%v limit reached.\n",securitytrailsKey)
			continue
		}
		if err != nil {
			fmt.Printf("\nsecuritytrails: %v\n", err)
			return subdomains
		}

		// Get the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("\nsecuritytrails: %v\n", err)
			return subdomains
		}

		// Decode the json format
		err = json.Unmarshal([]byte(body), &securitytrailsData)
		if err != nil {
			fmt.Printf("\nsecuritytrails: %v\n", err)
			return subdomains
		}

		// Append each subdomain found to subdomains array
		for _, subdomain := range securitytrailsData.Subdomains {
			finalSubdomain := subdomain + "." + domain
			if(silent){
				fmt.Printf("\n%s", finalSubdomain)
			}else{
				fmt.Printf("\n[%s] %s",serviceName, finalSubdomain)
			}
			subdomains = append(subdomains, finalSubdomain)
		}
		break;
	}

	return subdomains
}
