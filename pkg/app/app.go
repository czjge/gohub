package app

import "github.com/czjge/gohub/config"

func IsLocal() bool {
	return config.GetConfig().App.Env == "local"
}

func IsProduction() bool {
	return config.GetConfig().App.Env == "production"
}

func IsTesting() bool {
	return config.GetConfig().App.Env == "testing"
}
