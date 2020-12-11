module query

require dnsgrep/helper v0.0.0

replace dnsgrep/helper v0.0.0 => ../helper/

require dnsgrep/stringset v0.0.0

replace dnsgrep/stringset v0.0.0 => ../stringset/


require dnsgrep/domainutil v0.0.0

replace dnsgrep/domainutil v0.0.0 => ../domainutil/

require (
	github.com/bobesa/go-domain-util v0.0.0-20190911083921-4033b5f7dd89 // indirect
)