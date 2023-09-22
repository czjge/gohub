package app

import (
	"time"

	"github.com/czjge/gohub/config"
)

func IsLocal() bool {
	return config.GetConfig().App.Env == "local"
}

func IsProduction() bool {
	return config.GetConfig().App.Env == "production"
}

func IsTesting() bool {
	return config.GetConfig().App.Env == "testing"
}

func TimenowInTimezone() time.Time {
	chinaTimezone, _ := time.LoadLocation(config.GetConfig().App.Timezone)
	return time.Now().In(chinaTimezone)
}

func URL(path string) string {
	return config.GetConfig().App.Url + path
}

func V1URL(path string) string {
	return URL("/v1/" + path)
}
