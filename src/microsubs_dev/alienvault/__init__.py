# from microsubs_dev import __logger__

import requests


def query(api_key, domain, timeout):
    try:
        subdomains = []
        header = {'X-OTX-API-KEY': api_key}
        passive_dns = requests.get(
            'https://otx.alienvault.com:443/api/v1/indicators/domain/' + domain + '/passive_dns', headers=header, timeout=timeout)
        info = passive_dns.json()
        for i in info['passive_dns']:
            subdomains.append(i['hostname'])
        subdomains = set(subdomains)
        return subdomains
    except Exception as e:
        raise e
