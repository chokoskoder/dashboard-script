package config

import (
	"log/slog"
	"os"
	"strconv"
	"time"
)

type Config struct{
	DBURI string
	DBname string
	Timeout time.Duration
	IsDryRun bool
	LogLevel slog.Level
	Environment string

}

//we will need function which will help us populate the config struct which will then be used all across our project FROM HERE
//a constructor function to load all the config variables 
func Load() *Config{
	return &Config{
		DBURI: getEnvAsString("DB_URI" , ""),
		DBname: getEnvAsString("DB_NAME",""),
		Timeout: getEnvDuration("DB_TIMEOUT" , 30*time.Millisecond),
		IsDryRun: getEnvAsString("DRY_RUN" , "false") == "true",
		LogLevel: getEnvAsSlogLogLevel("LOG_LEVEL" , slog.LevelInfo),
		Environment: getEnvAsString("ENV" , "local"),
	}
}
//we need to setup a logger and fix this logging method -> no we dont need to , 

//helper functions to read them as their intended types
func getEnvAsString(key , defaultVal string) string {
	if value,exists := os.LookupEnv(key); exists{
		return value
	} 
	if defaultVal == ""{
		
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

func getEnvAsSlogLogLevel(key string , defaultVal slog.Level) slog.Level{
	valueStr := getEnvAsString(key , "")
	if valueStr ==  "" {
		return defaultVal
	}
	var level slog.Level
	err := level.UnmarshalText([]byte(valueStr))
	if err != nil {
		return defaultVal
	}
	return level
	//need to write logic here to conver valueStr to slog.Level like info and debug and all , I have no custom logging level for now
}

