package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	Mailer = "mailer"
)

var (
	Config *viper.Viper
	Cfg    = map[string]*viper.Viper{}
)

func init() {
	configJson()
	cfgJson()
	fmt.Println("config all init successed...")
}

func configJson() {
	Config = viper.New()
	Config.AddConfigPath("./")
	Config.SetConfigName("config.json")
	Config.SetConfigType("json")
	if err := Config.ReadInConfig(); err != nil {
		fmt.Println("configJson err:", err)
	}
}

func cfgJson() {
	v := viper.New()
	Cfg[Mailer] = v
	v.AddConfigPath("./config/cfg")
	v.SetConfigName(Mailer + ".json")
	v.SetConfigType("json")
	if err := v.ReadInConfig(); err != nil {
		fmt.Println("cfgJson err:", err)
	}
}
