package configuration

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Env     string
	BaseUrl string
}

func LoadConfig() (conf Config, err error) {
	log.Println("Load Config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./configuration")
	viper.SetConfigName("config")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalln("err in load config >>", err.Error())
		return
	}
	fmt.Println(">>>>", viper.GetString("app.env"))
	return
}
