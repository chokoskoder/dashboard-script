package config

import (
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
