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
		// default to dev environment for local development
		in = "config.dev"
	}

	v.SetConfigName(in)
	v.SetConfigType("yaml")
	// search order: parent dir first, then current dir
	v.AddConfigPath("../")
	v.AddConfigPath("./")
	// also check configs subdirectory
	v.AddConfigPath("./configs")
	// support configs at project root level
	v.AddConfigPath("../../")

	// allow environment variables to override config file values
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("failed to read config '%s': %v", in, err))
	}

	return v
}`
