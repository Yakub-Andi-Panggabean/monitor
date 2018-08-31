package util

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func init() {

	viper.SetConfigType("toml")
	if os.Getenv("ENVIRONMENT") == "dev" {
		viper.SetConfigName("config.dev")
	} else {
		viper.SetConfigName("config")
	}

	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath("$HOME/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error when reading config %s ", err))
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {

		fmt.Println("configuration file changed: ", e.Name)

	})

}

//GetConfig get key configuration from configuration file
func GetConfig(key string) interface{} {

	return viper.Get(key)

}
