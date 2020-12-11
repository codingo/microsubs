#pragma once
#include <vector>
#include "includes/request/HTTPRequest.hpp"
#include "includes/json/json.hpp"

using json = nlohmann::json;

std::vector<std::string> Query(std::string domain, bool silent, std::string serviceName, std::string apiKey)
{
    std::vector<std::string> sDomains;
    if (apiKey == "")
    {
        return sDomains;
    }
    try
    {
        // you can pass http::InternetProtocol::V6 to Request to make an IPv6 request
        http::Request request("http://api.securitytrails.com/v1/domain/" + domain + "/subdomains");

        // pass parameters as a map
        std::map<std::string, std::string> parameters = {};
        http::Response response = request.send("GET", parameters, {"APIKEY:" + apiKey});
        std::string body = std::string(response.body.begin(), response.body.end());

        json j = json::parse(body);

        auto subdomains = j["subdomains"];

        for (int i = 0; i < subdomains.size(); i++)
        {
            std::string s = subdomains[i];
            std::string finalSubomain = s + "." + domain;
            sDomains.push_back(finalSubomain);
            if (silent)
            {
                std::cout << finalSubomain << std::endl;
            }
            else
            {
                std::cout << "[" << serviceName << "] " << finalSubomain << std::endl;
            }
        }
    }
    catch (const std::exception &e)
    {
        std::cerr << "Request failed, error: " << e.what() << std::endl;
    }
    return sDomains;
}