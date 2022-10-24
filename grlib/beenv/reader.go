package beenv

import (
	"github.com/spf13/viper"
	"os"
)

// read init viper instance with yml settings and map to env interface{}
func read(path string, env interface{}) error {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(path)
	v.AutomaticEnv()
	v.SetConfigType("yml")
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	err := v.Unmarshal(env)
	if err != nil {
		return err
	}
	return nil
}

// LoadConfig read yml config from default path and map to interface{}
func LoadConfig(env interface{}) error {
	return LoadConfigFromPath(EnvConfigPath, env)
}

// LoadConfigFromPath read yml config from path and map to interface{}
func LoadConfigFromPath(cfgPath string, env interface{}) error {
	configPath := os.Getenv(cfgPath)
	if configPath == "" {
		configPath = "configs"
	}
	err := read(configPath, env)
	if err != nil {
		return err
	}
	return nil
}
