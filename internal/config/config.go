package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var flags = pflag.NewFlagSet("flags", pflag.ExitOnError)

func init() {
	// Start a logger
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file:", err)
	}
	log.SetOutput(logFile)

	// Define the flags and bind them to viper
	flags.StringP("ServerAddress", "a", ":8080", "HTTP server network address")

	// Parse the command-line flags
	err = flags.Parse(os.Args[1:])
	if err != nil {
		log.Printf("Error parsing flags: %v", err)
	}

	// Bind the flags to viper
	bindFlagToViper("ServerAddress")

	// Set the environment variable names
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	bindEnvToViper("ServerAddress", "TODO_PORT")

	// Read the environment variables
	viper.AutomaticEnv()
}

func bindFlagToViper(flagName string) {
	if err := viper.BindPFlag(flagName, flags.Lookup(flagName)); err != nil {
		log.Println(err)
	}
}

func bindEnvToViper(viperKey, envKey string) {
	if err := viper.BindEnv(viperKey, envKey); err != nil {
		log.Println(err)
	}
}

func Address() string {
	return viper.GetString("ServerAddress")
}
