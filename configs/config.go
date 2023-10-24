package configs

import (
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Config struct
type Config struct {
	App      `mapstructure:"app"`
	Postgres `mapstructure:"postgres"`
}

// App struct
type App struct {
	Debug bool   `mapstructure:"debug"`
	Env   string `mapstructure:"env"`
	Port  string `mapstructure:"port"`
}

// Postgres struct
type Postgres struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"database"`
	SSLMode  bool   `mapstructure:"sslmode"`
}

var config Config

// InitViper func
func InitViper(path, env string) {
	getConfig(path, env)
}

// GetViper func
func GetViper() *Config {
	return &config
}

func getConfig(path, env string) {
	viper.SetConfigName("config")
	viper.AddConfigPath(path)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file has changed: ", e.Name)
	})
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalln(err)
	}
}
