module dnsgrep

require dnsgrep/config v0.0.0

replace dnsgrep/config v0.0.0 => ./config/

require dnsgrep/helper v0.0.0

replace dnsgrep/helper v0.0.0 => ./helper/

require (
	dnsgrep/query v0.0.0
	github.com/parnurzeal/gorequest v0.2.16 // indirect
	moul.io/http2curl v1.0.0 // indirect
)

replace dnsgrep/query v0.0.0 => ./query/

require dnsgrep/stringset v0.0.0

replace dnsgrep/stringset v0.0.0 => ./stringset/

require dnsgrep/domainutil v0.0.0

replace dnsgrep/domainutil v0.0.0 => ./domainutil/
