#pragma once

#include "includes/json/json.hpp"
using json = nlohmann::json;
std::string CONFIGFILE;
int print()
{
    return 22;
}

json ReadConfig(std::string filename)
{
    try
    {
        CONFIGFILE = filename;
        std::ifstream i(filename);
        json j;
        i >> j;
        return j;
    }
    catch (const std::exception &e)
    {
        std::cerr << e.what() << '\n';
        std::cerr << "Exiting.\n Error while reading " << filename << ", please validate it\n";
        exit(1);
    }
}

json WriteConfig(json j)
{
    std::ofstream o(CONFIGFILE);
    o << std::setw(4) << j << std::endl;
}