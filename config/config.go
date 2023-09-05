package config

import "github.com/BurntSushi/toml"

var cfg Configuration

type Configuration struct {
	App   AppConfig           `toml:"app"`
	Mysql map[string]DBConfig `toml:"mysql"`
	Log   LogConfig           `toml:"log"`
	Redis RedisConfig         `toml:"redis"`
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
	Level     string `toml:"level"`
	Type      string `toml:"type"`
	Filename  string `toml:"filename"`
	MaxSize   int    `toml:"max_size"`
	MaxBackup int    `toml:"max_backup"`
	MaxAge    int    `toml:"max_age"`
	Compress  bool   `toml:"compress"`
}

type RedisConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	DB       int    `toml:"db"`
}

func InitConfig(path string) error {
	_, err := toml.DecodeFile(path, &cfg)
	return err
}

func GetConfig() Configuration {
	return cfg
}
