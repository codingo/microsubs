package helper

import "time"
import "log"
import "securitytrails/config"
// GetCurrentTime returns the current time
func  GetCurrentTime() string{
	dt :=  time.Now().Format(time.RFC3339)
	return dt
}
//ParseTime parsed the time from string to time
func ParseTime(datetime string) time.Time{
	time, err := time.Parse(time.RFC3339,datetime);
	if err!=nil {
		log.Println("Error while parsing date :", err);
	}
	return time
}

// GetOldestKey returns the oldest key used and update it usage
func GetOldestKey(cnfg *config.Configuration,paid bool) (string,int){
	oldestOne:=0
	minTime:=time.Now()
	
	keys:=cnfg.APIKeys
	
	if paid{
		keys=cnfg.PaidAPIKeys
	}
	
	for i, key := range keys{
		time:=ParseTime(key.LastUsed)
		if(time.Before(minTime)){
			oldestOne=i
			minTime=time
		}
	}
	// updating the access time
	keys[oldestOne].LastUsed=GetCurrentTime()
	// Updating configuration file
	config.WriteUpdatedConfig(*cnfg)

	return keys[oldestOne].Key,len(keys)
}