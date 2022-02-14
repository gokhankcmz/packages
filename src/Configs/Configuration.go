package Configs

var configurations AppConfig

func GetConfigs() *AppConfig {
	return &configurations
}

func SetDefault() {
	configurations = DefaultConfigs["dev"]
}
