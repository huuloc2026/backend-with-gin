package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper := viper.New()
	viper.AddConfigPath("config")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	//read config
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	fmt.Println("Server Port::", viper.GetInt("server.port"))
}
