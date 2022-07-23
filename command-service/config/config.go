package config

import "github.com/spf13/viper"

type DBConfig struct {
	MySql MySql `yaml:"mysql" mapstructure:"mysql"`
}

type MySql struct {
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
	DbName   string `yaml:"dbname"`
	Password string `yaml:"password"`
	User     string `yaml:"user"`
}

func (c *DBConfig) ReadConfigFile() {
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}
	if err := viper.Unmarshal(&c); err != nil {
		panic(err.Error())
	}
}
