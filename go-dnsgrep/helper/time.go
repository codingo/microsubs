package helper

import (
	"time"
    "log"
)

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