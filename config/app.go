package config

import "github.com/czjge/gohub/pkg/config"

func init() {
	config.Add("app", func() map[string]any {
		// 避免包初始化中即时求值
		return map[string]any{

			"name":     config.Env("APP_NAME", "Gohub"),
			"env":      config.Env("APP_ENV", "production"),
			"debug":    config.Env("APP_DEBUG", false),
			"port":     config.Env("APP_PORT", "3000"),
			"key":      config.Env("APP_KEY", "33446a9dcf9ea060a0b6532b166da32f304af0de"),
			"url":      config.Env("APP_URL", "http://localhost:3000"),
			"timezone": config.Env("TIMEZONE", "Asia/Shanghai"),
		}
	})
}
