package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // or "json"
	viper.AddConfigPath(".")      // look in the current directory

	viper.AutomaticEnv() // read in environment variables

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("⚠️  No config file found: %v\n", err)
	} else {
		fmt.Println("✅ Loaded config:", viper.ConfigFileUsed())
	}
}
