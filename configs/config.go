package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	PortServer    string
	DBString      string
	ServerAddress string
}

func ReadConfigFile(path string) (config Config) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	config.PortServer = viper.GetString("PORT_SERVER")
	config.DBString = viper.GetString("DATABASE_URL")
	config.ServerAddress = viper.GetString("SERVER_ADDRESS")

	return
}
