package config

import "github.com/spf13/viper"

var Config Conf

type Conf struct {
	Name  string
	Usage string
}

func Setup() {
	viper.SetConfigName("application")
	viper.SetConfigType("properties")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(err)
	}
}
