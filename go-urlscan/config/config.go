package config

import "github.com/spf13/viper"
import "log"

// ReadConfig reads the configuration file
func ReadConfig(file string) Configuration {
	viper.SetConfigFile(file)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading %s, %s", file, err)
	}
	var configuration Configuration
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	return configuration
}

// WriteUpdatedConfig updates the config
func WriteUpdatedConfig(config Configuration) {
	viper.Set("service", config.Service)
	viper.Set("apikeys", config.APIKeys)
	viper.WriteConfig()
}
