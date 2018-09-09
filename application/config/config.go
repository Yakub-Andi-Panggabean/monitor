package config

import (
	"fmt"
	"os"
	"strings"

	"firstwap.com/sms-monitoring-api/application/parser"
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
func Get(key string) interface{} {

	return viper.Get(key)

}

func GetAllDataSources() []string {

	dbSlice := make([]string, 0)

	for _, ds := range viper.AllKeys() {

		if strings.HasPrefix(ds, "database") {

			databases := strings.Split(ds, ".")[0] + "." + strings.Split(ds, ".")[1]

			if !parser.Contains(dbSlice, databases) && !strings.Contains(databases, "connection") {
				dbSlice = append(dbSlice, databases)
			}

		}

	}

	fmt.Println(dbSlice)

	return dbSlice

}
