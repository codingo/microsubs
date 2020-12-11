# from microsubs_dev import __logger__

import urllib.request
import requests
import json
import sys


def query(api_key, domain, timeout):
    subdomains = []
    try:
        header = {'X-Key': api_key}
        url = 'https://api.binaryedge.io/v2/query/domains/subdomain/' + domain
        response = requests.get(url, headers=header, timeout=timeout)
        info = response.json()
        for i in info['events']:
            subdomains.append(i)
        subdomains = set(subdomains)
        return subdomains
    except Exception as e:
        raise e
