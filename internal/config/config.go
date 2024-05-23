package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var flags = pflag.NewFlagSet("flags", pflag.ExitOnError)

// init инициализирует конфигурацию приложения
func init() {
	// запускаем логирование в файл
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		log.Fatalln("Failed to open log file:", err)
	}
	// устанавливаем logFile как стандартный вывод
	log.SetOutput(logFile)

	// подключем флаги
	flags.StringP("ServerAddress", "a", ":8080", "HTTP server network address")

	err = flags.Parse(os.Args[1:])
	if err != nil {
		log.Printf("Error parsing flags: %v", err)
	}

	bindFlagToViper("ServerAddress")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	bindEnvToViper("ServerAddress", "TODO_PORT")

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

// Address возвращает адрес сервера, в нашем случае порт!
func Address() string {
	return viper.GetString("ServerAddress")
}
