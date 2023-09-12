package config

import "github.com/BurntSushi/toml"

var cfg Configuration

type Configuration struct {
	App        AppConfig              `toml:"app"`
	Mysql      map[string]DBConfig    `toml:"mysql"`
	Log        LogConfig              `toml:"log"`
	Redis      map[string]RedisConfig `toml:"redis"`
	Captcha    CaptchaConfig          `toml:"captcha"`
	Sms        SmsConfig              `toml:"sms"`
	Verifycode VerifycodeConfig       `toml:"verifycode"`
}

type AppConfig struct {
	Name  string `toml:"name"`
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

type CaptchaConfig struct {
	Height          int     `toml:"height"`
	Width           int     `toml:"width"`
	Length          int     `toml:"length"`
	Maxskew         float64 `toml:"maxskew"`
	Dotcount        int     `toml:"dotcount"`
	ExpireTime      int     `toml:"expire_time"`
	DebugExpireTime int     `toml:"debug_expire_time"`
	TestingKey      string  `toml:"testing_key"`
}

type SmsConfig struct {
	AccessKeyId     string `toml:"access_key_id"`
	AccessKeySecret string `toml:"access_key_secret"`
	SignName        string `toml:"sign_name"`
	TemplateCode    string `toml:"template_code"`
}

type VerifycodeConfig struct {
	CodeLength       int    `toml:"code_length"`
	ExpireTime       int    `toml:"expire_time"`
	DebugExpireTime  int    `toml:"debug_expire_time"`
	DebugCode        string `toml:"debug_code"`
	DebugPhonePrefix string `toml:"debug_phone_prefix"`
	DebugEmailSuffix string `toml:"debug_email_suffix"`
}

func InitConfig(path string) error {
	_, err := toml.DecodeFile(path, &cfg)
	return err
}

func GetConfig() Configuration {
	return cfg
}
