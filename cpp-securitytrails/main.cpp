#include <iostream>
#include <vector>
#include "includes/CLI11/CLI11.hpp"
#include "config.hpp"
#include "helper.hpp"
#include "query.hpp"
#include "includes/json/json.hpp"
using json = nlohmann::json;

using namespace std;
std::string sourceDomain;
std::string outputFile;
bool silent;
std::string configFile;
int main(int argc, char **argv)
{
    CLI::App app{"This queries the securitytrails.com and prints discovered subdomains"};
    configFile = "config.json";
    silent = false;
    app.add_option("-d,--domain", sourceDomain, "Target domain e.g apple.com (required)");
    app.add_option("-o,--output", outputFile, "Output file (optional)");
    app.add_option("-s,--silent", silent, "To display only the subdomains");
    app.add_option("-c,--config", configFile, "Configuration file. (optional)");

    CLI11_PARSE(app, argc, argv);
    if (sourceDomain == "")
    {
        cout << "Domain is required." << endl;
        exit(0);
    }
    cout << "Domain is : " << sourceDomain << endl;
    json configuration = ReadConfig(configFile);
    string service = configuration["service"];

    cout << "Service is : " << service << endl;
    std::string apiKey = GetOldestKey(configuration);

    cout << "API key: " << apiKey << endl;
    vector<std::string> subDomains = Query(sourceDomain, silent, service, apiKey);

    if (outputFile != "")
    {

        FILE *fp;

        fp = fopen(outputFile.c_str(), "w");
        if (fp == NULL)
        {
            cout << "Error while writing to file\n";
            exit(-1);
        }
        for (int i = 0; i < subDomains.size(); i++)
        {
            const char *line = subDomains[i].c_str();
            fprintf(fp, "%s\n", line);
        }
        fclose(fp);
    }
    return 0;
}