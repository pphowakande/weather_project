package configuration

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// GetConfig ... This function gets configuration from .yaml files
func GetConfig() Configuration {
	env := os.Getenv("TIER")
	fmt.Println("env : ", env)
	if env == "" {
		env = "local"
	}
	if env != "production" && env != "local" {
		fmt.Println("invalid env for loading the configuration set TIER value as local or production")
	}
	viper.SetConfigName(env)
	viper.AddConfigPath("./src/config/tier")
	var configuration Configuration
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file : ", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Println("unable to decode into struct : ", err)
	}
	return configuration
}
