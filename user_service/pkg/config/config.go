package config

import (
	_ "github.com/EmreZURNACI/ProdFullReadyApp_User/pkg/log"
	"github.com/spf13/viper"
)

type Appconfig struct {
	SrvConfig   serverconfig   `mapstructure:"server"`
	PgConfig    postgresconfig `mapstructure:"aws_postgres"`
	MongoConfig mongoconfig    `mapstructure:"mongo"`
}

type serverconfig struct {
	Port int `mapstructure:"port"`
}

type postgresconfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	SSLMode  string `mapstructure:"sslmode"`
}

type mongoconfig struct {
	User       string `mapstructure:"user"`
	Password   string `mapstructure:"password"`
	AppName    string `mapstructure:"appname"`
	DbName     string `mapstructure:"name"`
	Collection string `mapstructure:"collection"`
}

// Golang'de dosya pathleri main.go'dan itibaren verilir.
func ReadConfig() *Appconfig {
	viper.SetConfigName("config")                                 // name of config file (without extension)
	viper.SetConfigType("yaml")                                   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./internal/.config")                     // path to look for the config file in
	viper.AddConfigPath("$HOME/ProdFullReayApp/internal/.config") // call multiple times to add many search paths
	viper.AddConfigPath(".")                                      // optionally look for config in the working directory
	err := viper.ReadInConfig()                                   // Find and read the config file
	if err != nil {                                               // Handle errors reading the config file
		return nil
	}
	var config Appconfig
	if err := viper.Unmarshal(&config); err != nil {
		return nil
	}
	return &config
}
