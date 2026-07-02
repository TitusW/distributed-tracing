package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

var cfgPath string

type Config struct {
	Database DatabaseConfig `yaml:"database"`
}

type DatabaseConfig struct {
	Name     string `yaml:"name" env:"DB_NAME" env-description:"Database name"`
	Port     string `yaml:"port" env:"DB_PORT" env-description:"Database port"`
	Host     string `yaml:"host" env:"DB_HOST" env-description:"Database host"`
	Username string `yaml:"username" env:"DB_USERNAME" env-description:"Username of DB"`
	Password string `yaml:"password" env:"DB_PASSWORD" env-description:"Password for the corresponding Username DB"`
}

func InitializeConfig() Config {
	initFlag()

	var cfg Config

	err := cleanenv.ReadConfig(getConfigPath(), &cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	return cfg
}

func initFlag() {
	flag.StringVar(&cfgPath, "config", "files/app_config/config.dev.yaml", "location of config file")
	flag.Parse()
}

func getConfigPath() string {
	return cfgPath
}
