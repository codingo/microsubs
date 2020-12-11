module dnsdumpster

require dnsdumpster/config v0.0.0

replace dnsdumpster/config v0.0.0 => ./config/

require dnsdumpster/helper v0.0.0

replace dnsdumpster/helper v0.0.0 => ./helper/

require (
	dnsdumpster/query v0.0.0
	github.com/subfinder/subfinder v0.1.1 // indirect
)

replace dnsdumpster/query v0.0.0 => ./query/
