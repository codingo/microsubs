module passivetotal

require passivetotal/config v0.0.0

replace passivetotal/config v0.0.0 => ./config/

require passivetotal/helper v0.0.0

replace passivetotal/helper v0.0.0 => ./helper/

require (
	github.com/subfinder/subfinder v0.1.1 // indirect
	passivetotal/query v0.0.0
)

replace passivetotal/query v0.0.0 => ./query/
