#pragma once
#include "includes/json/json.hpp"
#include "config.hpp"
#include <sstream>
#include <ctime>
#include <chrono>
#include <time.h>
#include <iomanip>
using json = nlohmann::json;
using namespace std::chrono;

char *GetCurrentTime()
{
    time_t current_time;
    current_time = time(NULL);
    tm *curr_tm;
    char *date_string = new char[100];
    curr_tm = localtime(&current_time);

    strftime(date_string, 50, "%Y-%m-%d %T", curr_tm);
    return date_string;
}

int isGreater(time_t time1, time_t time2)
{
    return difftime(time1, time2) > 0.0 ? 1 : 0;
}
std::string GetOldestKey(json cnfg)
{
    auto APIKeys = cnfg["apikeys"];
    int counter = 0;
    int oldestOne = 0;
    std::string currentTime = GetCurrentTime();
    const char *time_details1 = currentTime.c_str();
    struct tm tm;
    strptime(time_details1, "%Y-%m-%d %T", &tm);
    time_t minTime = mktime(&tm);

    for (auto it = APIKeys.begin(); it != APIKeys.end(); ++it)
    {
        std::string date1 = cnfg["apikeys"][counter]["LastUsed"];

        const char *time_details = date1.c_str();
        struct tm tm1;
        strptime(time_details, "%Y-%m-%d %T", &tm1);
        time_t time = mktime(&tm1);

        if (!isGreater(time, minTime))
        {
            oldestOne = counter;
            minTime = time;
        }

        counter++;
    }
    std::cout << "Oldest one is " << oldestOne << std::endl;
    // Updating acess time
    std::string s = GetCurrentTime();
    cnfg["apikeys"][oldestOne]["LastUsed"] = s;

    // Updating file
    WriteConfig(cnfg);
    // Returning key
    return cnfg["apikeys"][oldestOne]["Key"];
}