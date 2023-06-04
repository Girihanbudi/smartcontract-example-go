package env

import (
	"fmt"
	"smartm2m-technical-test/internal/pkg/env/config"
	"smartm2m-technical-test/internal/pkg/log"

	"github.com/spf13/viper"
)

const Instance string = "Env"

// global env declaration
var CONFIG config.Config

type Options struct {
	Path     string
	FileName string
	Ext      string
}

func NewDefaultOptions() Options {
	return Options{
		Path:     "./env",
		FileName: "config",
		Ext:      "yaml",
	}
}

func InitEnv(options Options) {
	log.Print(Instance, "reading config...")

	viper.AddConfigPath(options.Path)
	viper.SetConfigName(options.FileName)
	viper.SetConfigType(options.Ext)

	env := config.Config{}
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(Instance, "failed to read config", err)
	}

	if err := viper.Unmarshal(&env); err != nil {
		log.Fatal(Instance, "failed to unmarshal config", err)
	}

	log.Print(Instance, fmt.Sprintf("using %s stage mode", env.Stage))
	CONFIG = env
}

func ProvideEnv() config.Config {
	return CONFIG
}
