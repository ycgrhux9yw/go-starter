package code

var ViperConfiguration = `package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func NewViper(in string) *viper.Viper {
	v := viper.New()

	switch in {
	case "prod":
		in = "config.prod"
	case "dev":
		in = "config.dev"
	case "staging":
		in = "config.staging"
	default:
		in = "config.dev"
	}

	v.SetConfigName(in)
	v.SetConfigType("yaml")
	v.AddConfigPath("../")
	v.AddConfigPath("./")

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("failed to read config '%s': %v", in, err))
	}

	return v
}`
