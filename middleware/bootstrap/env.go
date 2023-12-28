package bootstrap

import (
	"github.com/spf13/viper"

	"log"
)

type Env struct {
	Port     string `mapstructure:"PORT"`
	AuthPort string `mapstructure:"AUTH_PORT"`
	AppEnv   string `mapstructure:"APP_ENV"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env: ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Unable to decode into struct: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development mode")
	} else {
		log.Println("The App is running in production mode")
	}

	return &env
}
