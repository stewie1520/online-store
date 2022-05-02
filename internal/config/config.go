package config

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/jessevdk/go-flags"
	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		DriverName string `mapstructure:"driver_name"`
		Url        string `mapstructure:"url"`
	} `mapstructure:"database"`
	Port        int `mapstructure:"port"`
	Environment string
}

func (c Config) IsDev() bool {
	return c.Environment == "development"
}

func (c Config) IsProd() bool {
	return c.Environment == "production"
}

func (c Config) IsStaging() bool {
	return c.Environment == "staging"
}

var AppConfig *Config

func Init() {
	var cli struct {
		Config      string `short:"c" long:"config" description:"config file"`
		Environment string `short:"e" long:"environment" description:"app environment" default:"development"`
	}

	parser := flags.NewParser(&cli, flags.Default)
	if _, err := parser.Parse(); err != nil {
		log.Panicln(err)
	}

	if cli.Config == "" {
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Panicln(err)
		}

		cli.Config = path.Join(dir, "..", "internal/config/config.yaml")
	}

	viper.SetConfigFile(cli.Config)
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	var configs map[string]Config
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("error reading config file %s", err)
	}

	err := viper.Unmarshal(&configs)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	config := configs[cli.Environment]
	config.Environment = cli.Environment
	AppConfig = &config
}
