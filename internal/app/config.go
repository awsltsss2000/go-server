package app

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var config = new(Config)

type Config struct {
	RunMode string `yaml:"runMode"`

	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`

	Gorm struct {
		MaxIdleConns int `yaml:"maxIdleConns"`
		MaxOpenConns int `yaml:"maxOpenConns"`
		MaxLifetime  int `yaml:"maxLifetime"`
	} `yaml:"gorm"`

	Postgres struct {
		Driver string `yaml:"driver"`
		Dsn    string `yaml:"dsn"`
	} `yaml:"postgres"`

	Redis struct {
		Address     string `yaml:"address"`
		Password    string `yaml:"password"`
		MaxIdle     int    `yaml:"maxIdle"`
		MaxActive   int    `yaml:"maxActive"`
		IdleTimeout int    `yaml:"idleTimeout"`
	} `yaml:"redis"`
}

func InitConfig() {
	configFile := fmt.Sprintf("configs/%s.yaml", os.Getenv("DEPLOY_ENV"))
	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	// viper.AutomaticEnv()
	// viper.WatchConfig()
	// viper.OnConfigChange(func(e fsnotify.Event) {
	// 	if err := viper.Unmarshal(config); err != nil {
	// 		panic(err)
	// 	}
	// })
}

func GetConfig() Config {
	return *config
}
