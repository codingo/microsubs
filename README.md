# Microsubs

A collection of files that can be used for individual source subdomain enumeration, or pulled together in a bash script, or interlace file (https://github.com/codingo/Interlace) to fully map a target.

Each file can be called on the command line using the following:

| Argument  | Description                                                 | Implemented |
| --------- | ----------------------------------------------------------- | ----------- |
| -c string | Configuration file to use (optional) (default: config.json) | ✔           |
| -k string | Key value to search for in associated configuration file    | ✔           |
| -premium  | When set, only use keys marked as premium keys              | ✔           |
| -d string | Target domain (e.g. apple.com) (required)                   | ✔           |
| -nC       | don't use colours in output                                 |             |
| -timout   | number of seconds to wait before timing out (default 90)    | ✔           |
| -o string | Output file (optional)                                      | ✔           |
| -silent   | Display only results without source next to them            | ✔           |
| -h        | Display this menu                                           | ✔           |
| -v        | Show Verbose output                                         | ✔           |
| -version  | Display microsubs version                                   | ✔           |

Config.json can be shared over all modules. Each service has set of API keys, and can use multiple keys with the last used date stored against them. Keys are cycled based on when they were last used. Some services won't have API keys, and will instead have other properties stored against them.

For example, here's a config.json that could be used with either the DNSGrep module, the Shodan module, or the URLScan module:

```{
    "normalkeys": {
        "passivetotal": {
            "apikeys": [
                {
                    "lastused": "2020-07-19 05:15:50.744273",
                    "key": "APIKEY"
                },
                {
                    "lastused": "2020-07-19 05:18:10.788507",
                    "key": "APIKEY"
                },
                {
                    "lastused": "2020-07-17 15:39:41.494927",
                    "key": "APIKEY"
                },
                {
                    "lastused": "2020-07-17 15:39:41.494927",
                    "key": "APIKEY"
                },
                {
                    "lastused": "2020-07-17 15:39:41.494927",
                    "key": "APIKEY"
                }
            ],
            "service": "passivetotal.org",
            "username": "username"
        }
        "dnsdumpster": {
            "service": "dnsdumpster.com"
        },
        "alienvault": {
            "apikeys": [
                {
                    "lastused": "2020-07-20 19:48:11.858701",
                    "key": "APIKEY"
                },
                {
                    "lastused": "2020-07-17 15:40:46.046895",
                    "key": "APIKEY"
                },
                {
                    "lastused": "2020-07-17 15:42:19.718550",
                    "key": "APIKEY"
                },
                {
                    "lastused": "2020-07-17 15:40:46.837664",
                    "key": "APIKEY"
                }
            ],
            "service": "otx.alienvault.com"
        }
    },
    "premiumkeys": {
        "passivetotal": {
            "apikeys": [
                {
                    "lastused": "2020-07-17 15:39:41.494927",
                    "key": "APIKEY"
                },
                {
                    "lastused": "2020-07-17 15:39:41.494927",
                    "key": "APIKEY"
                },
                {
                    "lastused": "2020-07-17 15:39:41.494927",
                    "key": "APIKEY"
                },
                {
                    "lastused": "2020-07-17 15:39:41.494927",
                    "key": "APIKEY"
                },
                {
                    "lastused": "2020-07-17 15:39:41.494927",
                    "key": "APIKEY"
                }
            ],
            "service": "passivetotal.org",
            "username": "username"
        },
        "dnsdumpster": {
            "service": "dnsdumpster.com"
        },
        "alienvault": {
            "apikeys": [
                {
                    "lastused": "2020-07-17 15:40:44.915433",
                    "key": "APIKEY"
                },
                {
                    "lastused": "2020-07-17 15:40:46.046895",
                    "key": "APIKEY"
                },
                {
                    "lastused": "2020-07-17 15:42:19.718550",
                    "key": "APIKEY"
                },
                {
                    "lastused": "2020-07-17 15:40:46.837664",
                    "key": "APIKEY"
                }
            ],
            "service": "otx.alienvault.com"
        }
    }
}
```

Modules should each live in their own folder, but should all share the same code base for arguments, or managing configurations to allow them to be built as quickly as possible. Ideally, a new service should be able to be provisioned inside of an hour, and a new command line argument that would impact on all modules changed in a single location.

Services to Build

| Service                  | API Link                                                                  | Description | Python |
| ------------------------ | ------------------------------------------------------------------------- | ----------- | ------ |
| SecurityTrails           | https://api.securitytrails.com/v1/                                        |             | ✔      |
| Censys                   | https://www.censys.io/api/v1                                              |             | ❌     |
| DNSDumpster              | https://dnsdumpster.com                                                   |             | ✔      |
| DNSGrep                  | http://dns.bufferover.run/dns?q=                                          |             | ✔      |
| PassiveTotal             | https://api.passivetotal.org                                              |             | ✔      |
| Shodan                   | https://api.shodan.io/shodan/host/search?query=hostname:                  |             | ❌     |
| URLScan                  | https://urlscan.io/api/v1/result/                                         |             | ❌     |
| Virus Total              | https://www.virustotal.com/vtapi/v2/domain/report                         |             | ✔      |
| Wayback Archive          |                                                                           |             |
| AbuseIPDB                |                                                                           |             |
| AlientVault OTX          | https://otx.alienvault.com:443/api/v1/indicators/domain/                  |             | ✔      |
| Apility                  |                                                                           |             |
| Bad Packets              |                                                                           |             |
| BinaryEdge.io            | https://api.binaryedge.io/v2/query/domains/subdomain/                     |             | ✔      |
| Bing API for Bing Search |                                                                           |             |
| Botscout.com             |                                                                           |             |
| Builtwith.com            | https://api.builtwith.com/domain-api                                      |             |
| CIRCL.lu                 |                                                                           |             |
| Leak-Lookup              |                                                                           |             |
| Clearbit.com             | https://clearbit.com/docs?python#enrichment-api-company-api-domain-lookup |             |
| Fraudguard.io            |                                                                           |             |
| Fullcontact.com          |                                                                           |             |

### Usage

````
usage: mainentry.py [-h] [--version] [-c CONFIG] [-k KEYVALUE] [-p] [-v] [-vv]
                    [-nC NOCOLOR] [-t TIMEOUT] [-o OUTPUT] [-s SILENT]
                    domain

A collection of files that can be used for individual source subdomain
enumeration, or pulled together in a bash script.

positional arguments:
  domain                Target domain (e.g. apple.com) (required)

optional arguments:
  -h, --help            show this help message and exit
  --version             show program's version number and exit
  -c CONFIG, --config CONFIG
                        Configuration file to use (optional) (default:
                        config.json)
  -k KEYVALUE, --keyvalue KEYVALUE
                        Key value to search for in associated configuration
                        file supported services ('shodan', 'securitytrails',
                        'censys', 'dnsdumpster', 'dnsgrep', 'passivetotal',
                        'urlscan', 'virustotal', 'waybackarchive',
                        'alienvault')
  -p, --premium         When set, only use keys marked as premium keys
  -v, --verbose         set loglevel to INFO
  -vv, --very-verbose   set loglevel to DEBUG
  -nC NOCOLOR, --nocolor NOCOLOR
                        don't use colours in output
  -t TIMEOUT, --timeout TIMEOUT
                        number of seconds to wait before timing out (default
                        90)
  -o OUTPUT, --output OUTPUT
                        Output file (optional)
  -s SILENT, --silent SILENT
                        Display only results without source next to them```
````
