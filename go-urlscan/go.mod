module urlscan

require urlscan/config v0.0.0

replace urlscan/config v0.0.0 => ./config/

require urlscan/helper v0.0.0

replace urlscan/helper v0.0.0 => ./helper/

require urlscan/query v0.0.0

replace urlscan/query v0.0.0 => ./query/

require (
	github.com/caffix/cloudflare-roundtripper v0.0.0-20181218223503-4c29d231c9cb // indirect
	github.com/keegancsmith/rpc v1.1.0 // indirect
	github.com/robertkrimen/otto v0.0.0-20180617131154-15f95af6e78d // indirect
	github.com/stamblerre/gocode v0.0.0-20190327203809-810592086997 // indirect
	gopkg.in/sourcemap.v1 v1.0.5 // indirect
	urlscan/stringset v0.0.0
)

replace urlscan/stringset v0.0.0 => ./stringset/
