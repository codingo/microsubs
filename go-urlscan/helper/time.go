package helper

import "time"
import "log"
import "urlscan/config"

// GetCurrentTime returns the current time
func GetCurrentTime() string {
	dt := time.Now().Format(time.RFC3339)
	return dt
}

//ParseTime parsed the time from string to time
func ParseTime(datetime string) time.Time {
	time, err := time.Parse(time.RFC3339, datetime)
	if err != nil {
		log.Println("Error while parsing date :", err)
	}
	return time
}

// GetOldestKey returns the oldest key used and update it usage
func GetOldestKey(cnfg *config.Configuration) string {
	oldestOne := 0
	minTime := time.Now()
	for i, key := range cnfg.APIKeys {
		time := ParseTime(key.LastUsed)
		if time.Before(minTime) {
			oldestOne = i
			minTime = time
		}
	}
	// updating the acess time
	cnfg.APIKeys[oldestOne].LastUsed = GetCurrentTime()
	// Updating configuration file
	config.WriteUpdatedConfig(*cnfg)

	return cnfg.APIKeys[oldestOne].Key
}
