package initalize

import (
	"fmt"

	"github.com/huuloc2026/go-backend/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("config")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	//read config
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("unable to decode into struct, %w", err))
	}
}
