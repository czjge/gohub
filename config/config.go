package config

import "github.com/BurntSushi/toml"

var cfg Configuration

type Configuration struct {
	App   AppConfig `toml:"app"`
	Mysql DBConfig  `toml:"mysql"`
	Log   LogConfig `toml:"log"`
}

type AppConfig struct {
	Env   string `toml:"env"`
	Debug bool   `toml:"debug"`
	Url   string `toml:"url"`
	Port  string `toml:"port"`
}

type DBConfig struct {
}

type LogConfig struct {
}

func InitConfig(path string) error {
	_, err := toml.DecodeFile(path, &cfg)
	return err
}

func GetConfig() Configuration {
	return cfg
}
