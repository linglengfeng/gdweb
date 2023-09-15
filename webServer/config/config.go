package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var Config *viper.Viper

// var ConfigMap map[string]interface{}
func init() {
	configJson()
	fmt.Println("config init configJson successed")
}

func configJson() {
	Config = viper.New()
	Config.AddConfigPath("./")
	Config.SetConfigName("config.json")
	Config.SetConfigType("json")
	if err := Config.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
}
