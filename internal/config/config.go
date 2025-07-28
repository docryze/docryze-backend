package config

import "github.com/spf13/viper"

type Config struct {
	Application struct {
		Name string `yaml:"name"`
	} `yaml:"application"`
	Server struct {
		Addr string `yaml:"addr"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		DSN string `yaml:"dsn"`
	} `yaml:"database"`
}

func LoadConfig() (*Config, error) {
	// 读取YAML配置
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	return &config, nil
}
