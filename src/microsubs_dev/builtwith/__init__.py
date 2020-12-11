import requests
import json


def query(api_key, domain, timeout):
    subdomains = []
    response = requests.get(
        "https://api.builtwith.com/v14/api.json?KEY={0}&LOOKUP={1}".format(api_key, domain), headers={})
    data = response.json()
    for result in data['Results']:
        paths = result['Result']['Paths']
        for path in paths:
            subdomains.append(path["SubDomain"]+"."+domain)
    return set(subdomains)
