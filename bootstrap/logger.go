package bootstrap

import (
	"github.com/czjge/gohub/config"
	"github.com/czjge/gohub/pkg/logger"
)

func SetupLogger() {

	logger.InitLogger(
		config.GetConfig().Log.Filename,
		config.GetConfig().Log.MaxSize,
		config.GetConfig().Log.MaxBackup,
		config.GetConfig().Log.MaxAge,
		config.GetConfig().Log.Compress,
		config.GetConfig().Log.Type,
		config.GetConfig().Log.Level,
	)
}
