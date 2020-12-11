package helper

import "strings"

//Validate returns valid subdomains found ending with target domain
func Validate(domain string, strslice []string) (subdomains []string) {
	for _, entry := range strslice {
		if strings.HasSuffix(entry, "."+domain) {
			subdomains = append(subdomains, entry)
		}
	}

	return subdomains
}
