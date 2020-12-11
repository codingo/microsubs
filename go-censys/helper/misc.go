package helper

import "strings"

// Unique Returns unique items in a slice
// Adapted from http://www.golangprograms.com/remove-duplicate-values-from-slice.html
func Unique(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

//Validate returns valid subdomains found ending with target domain
func Validate(domain string, strslice []string) (subdomains []string) {
	for _, entry := range strslice {
		if strings.HasSuffix(entry, "."+domain) {
			subdomains = append(subdomains, entry)
		}
	}

	return subdomains
}
