package bootstrap

import (
	"time"

	"github.com/czjge/gohub/config"
	"github.com/czjge/gohub/pkg/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupDB() {

	// var dbConfig gorm.Dialector

	// prepare dialector
	var dbConfig gorm.Dialector = mysql.New(mysql.Config{
		DSN: config.GetConfig().Mysql.DSN,
	})

	// connect DB
	database.Connect(dbConfig, logger.Default.LogMode(logger.Info))

	// set DB parameters
	database.SQLDB.SetMaxOpenConns(config.GetConfig().Mysql.MaxOpenConns)
	database.SQLDB.SetMaxIdleConns(config.GetConfig().Mysql.MaxIdleConns)
	database.SQLDB.SetConnMaxLifetime(time.Duration(config.GetConfig().Mysql.ConnMaxLifetime))
}
