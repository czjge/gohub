package config

import "github.com/BurntSushi/toml"

var cfg Configuration

type Configuration struct {
	App   AppConfig           `toml:"app"`
	Mysql map[string]DBConfig `toml:"mysql"`
	Log   LogConfig           `toml:"log"`
}

type AppConfig struct {
	Env   string `toml:"env"`
	Debug bool   `toml:"debug"`
	Url   string `toml:"url"`
	Port  string `toml:"port"`
}

type DBConfig struct {
	DSN             string `toml:"dsn"`
	MaxOpenConns    int    `toml:"max_open_conns"`
	MaxIdleConns    int    `toml:"max_idle_conns"`
	ConnMaxLifetime int    `toml:"conn_max_lifetime"`
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
