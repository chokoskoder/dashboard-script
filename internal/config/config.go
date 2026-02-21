package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct{
	DBURI string
	DBname string
	Timeout time.Duration
	IsDryRun bool
	LogLevel string
}

//we will need function which will help us populate the config struct which will then be used all across our project FROM HERE
//a constructor function to load all the config variables 
func Load() *Config{
	return &Config{
		DBURI: getEnvAsString("DB_URI" , ""),
		DBname: getEnvAsString("DB_NAME",""),
		Timeout: getEnvDuration("DB_TIMEOUT" , 30*time.Millisecond),
		IsDryRun: getEnvAsString("DRY_RUN" , "false") == "true",
		LogLevel: getEnvAsString("LOG_LEVEL" , "info"),
	}
}
//we need to setup a logger and fix this logging method -> no we dont need to , 

//helper functions to read them as their intended types
func getEnvAsString(key , defaultVal string) string {
	if value,exists := os.LookupEnv(key); exists{
		return value
	} 
	if defaultVal == ""{
		//log here that there is something wrong with what we are making 
	}
	return defaultVal
}

func getEnvAsInt(key string, defaultVal int) int {
	valueStr := getEnvAsString(key , "")
	if valueStr == ""{
		return defaultVal
	}
	value,err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultVal
	}
	return value
}

func getEnvDuration(key string, defaultVal time.Duration) time.Duration {
	valueStr := getEnvAsString(key , "")
	if valueStr == ""{
		return defaultVal
	}
	duration,err := time.ParseDuration(valueStr)
	if err != nil {
		return defaultVal
	}
	return duration

}

