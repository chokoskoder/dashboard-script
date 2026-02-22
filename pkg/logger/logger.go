package logger

import (
	"log/slog"
	"os"	
)

type AppEnv string

const (
	EnvLocal AppEnv = "local"
	EnvProd AppEnv = "prod"
)

func SetupLogger( level slog.Level, environemt string) *slog.Logger {
	var logLevel slog.Leveler = level
	//how are we passing a slog.Level value to a slog.Leveler value ? how does this wokr ?

	opts := &slog.HandlerOptions{
		Level: logLevel,
		AddSource: environemt == string(EnvLocal),
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			return a
			//write this logic 
			//we need to mask passwords and personal info here(what is something that I need to mask in this script)
			//replace ist with utc ? why shouldnt I use it in ist ?
		},
	}

	var handler slog.Handler
	if environemt == "Local"{
		handler = slog.NewTextHandler(os.Stdout, opts)
	} else {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	}

	logger := slog.New(handler)
	logger.With(
		slog.Float64("go_version" , 1.21),
	)
	return logger
}

//how do I initialise this config with the data from the original config and is that the way to go about it ?
// I will just pass the necessary variables ahead to the setup function and work with it 