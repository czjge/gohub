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
