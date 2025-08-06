package config

import (
	"fmt"
	"log"
	"os"
	"github.com/spf13/viper"
)

type AppEnv struct {
	MONGO_URI  string
	DB_NAME    string
	PORT       string
	JWT_SECRET string
	JWT_EXPIRES string
}

var Env AppEnv

func Init() {

	developmentEnv := os.Getenv("GO_ENV")
	if developmentEnv == "" {
		developmentEnv = "local"
	}

	configFileName := fmt.Sprintf("config.%v", developmentEnv)
	viper.SetConfigName(configFileName) // looks for config.yaml
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config: %s", err)
	}

	Env = AppEnv{
		MONGO_URI:  viper.GetString("mongodb.uri"),
		DB_NAME:    viper.GetString("mongodb.name"),
		PORT:      viper.GetString("server.port"),
		JWT_SECRET: viper.GetString("jwt.secret"),
		JWT_EXPIRES: viper.GetString("jwt.expires"),
	}
}
