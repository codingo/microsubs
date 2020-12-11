package helper

import (
	"html"
	"net/url"
	"strings"
	"unicode"
)

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

// SubdomainExists checks if a key exists in an array
func SubdomainExists(key string, values []string) bool {
	for _, data := range values {
		if key == data {
			return true
		}
	}
	return false
}

// ExtractSubdomains extracts a subdomain from a big blob of text
func ExtractSubdomainsFromText(text, domain string) (urls []string) {
	allUrls := ExtractSubdomains(text, domain)

	return Validate(domain, allUrls)
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

// ExtractSubdomains finds all subdomains from a given text
func ExtractSubdomains(text, domain string) (urls []string) {
	allUrls := findAllUrls(text)
	var finalUrls []string

	for _, u := range allUrls {
		finalUrls = append(finalUrls, handleURI(u)...)
	}

	// Filter by domains and remove duplicates
	finalUrls = filterByDomain(finalUrls, domain)

	return finalUrls
}

func findAllUrls(text string) (urls []string) {
	for i, r := range text {
		if r == '.' {
			bck := string(r)
			//Go back till first valid ascii or number
			for backIndex := i - 1; backIndex >= 0; backIndex-- {
				rr := rune(text[backIndex])
				if isValidRuneBack(rr) {
					bck = string(rr) + bck
				} else {
					break
				}
			}

			//Go forth till the last valid ascii or number
			for forwardIndex := i + 1; forwardIndex < len(text); forwardIndex++ {
				rr := rune(text[forwardIndex])
				if isValidRuneForward(rr) {
					bck = bck + string(rr)
				} else {
					break
				}
			}
			urls = append(urls, bck)
		}
	}

	return urls
}

func isValidRuneBack(r rune) bool {
	return unicode.IsNumber(r) || unicode.IsLetter(r) || r == ':' || r == '/' || r == '_' || r == '-' || r == '%'
}

func isValidRuneForward(r rune) bool {
	return isValidRuneBack(r) || r == '.'
}

func handleURI(u string) []string {
	var urls []string
	// Try to parse as normal URI
	if u, err := url.ParseRequestURI(u); err == nil {
		urls = append(urls, u.Host)
		return urls

	}

	// Html Unescape
	u = html.UnescapeString(u)

	// Query Unescape
	u, _ = url.QueryUnescape(u)

	replacer := strings.NewReplacer(
		"u003d", " ",
		"/", " ",
		"\\", " ",
		"-site:", " ",
		"-www", "www",
	)

	// Suppress bad chars
	u = replacer.Replace(u)

	// Suppress bad starting characters
	u = suppressLeftChar(u)

	// Split on spaces
	return strings.Split(u, " ")
}

func suppressLeftChar(s string) string {
	if strings.HasPrefix(s, "-www") {
		return s[1:]
	}

	if strings.HasPrefix(s, "-site:") {
		return s[6:]
	}

	for i, r := range s {
		if r == '/' {
			return s[i:]
		}
	}

	return s
}

func filterByDomain(urls []string, domain string) []string {
	result := []string{}
	seen := map[string]string{}
	for _, u := range urls {
		if strings.HasSuffix(u, domain) {
			if _, ok := seen[u]; !ok {
				result = append(result, u)
				seen[u] = u
			}
		}
	}
	return result
}
