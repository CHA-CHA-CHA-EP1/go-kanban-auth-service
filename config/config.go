package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	Env    string `mapstructure:"env" validate:"required,oneof=dev prod"`
	App    App    `mapstructure:"app" validate:"required"`
	Server Server `mapstructure:"server" validate:"required"`
	Redis  Redis  `mapstructure:"redis" validate:"required"`
}

type App struct {
	Name string `mapstructure:"name" validate:"required"`
}

type Server struct {
	Port    int `mapstructure:"port" validate:"required,gte=1,lte=65535"`
	Timeout int `mapstructure:"timeout" validate:"required,gte=1"`
}

type Redis struct {
	Host     string `mapstructure:"host" validate:"required,hostname"`
	Port     int    `mapstructure:"port" validate:"required,gte=1,lte=65535"`
	DB       int    `mapstructure:"db"`
	Password string `mapstructure:"password"`
}

func (c *Config) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

func InitialConfig() *Config {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	err = config.Validate()
	if err != nil {
		panic(err)
	}
	return &config
}
