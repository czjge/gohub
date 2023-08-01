package bootstrap

import (
	"flag"
	"fmt"

	"github.com/czjge/gohub/config"
)

func SetupConfig() error {
	configPath := flag.String("c", "config/config.toml", "config file path")
	if err := config.InitConfig(*configPath); err != nil {
		return fmt.Errorf("load config file error, %v", err)
	}
	return nil
}
