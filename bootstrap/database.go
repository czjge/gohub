package bootstrap

import (
	"time"

	"github.com/czjge/gohub/app/models/user"
	"github.com/czjge/gohub/config"
	"github.com/czjge/gohub/pkg/database"
	"github.com/czjge/gohub/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDB() {

	var dbConfigs map[string]gorm.Dialector

	// prepare dialectors
	connectInfo := config.GetConfig().Mysql
	for name, config := range connectInfo {
		if dbConfigs == nil {
			dbConfigs = make(map[string]gorm.Dialector, len(connectInfo))
		}
		dbConfigs[name] = mysql.New(mysql.Config{
			DSN: config.DSN,
		})
	}

	// connect DB
	for k, v := range dbConfigs {
		if database.DBCollections == nil {
			database.DBCollections = make(map[string]*database.DBInfo, len(dbConfigs))
		}
		DB, SQLDB, err := database.Connect(v, logger.NewGormLogger())
		if err != nil {
			panic(err)
		}
		dbStruct := &database.DBInfo{
			DB:    DB,
			SQLDB: SQLDB,
		}
		database.DBCollections[k] = dbStruct
		// set DB parameters
		database.DBCollections[k].SQLDB.SetMaxOpenConns(config.GetConfig().Mysql[k].MaxIdleConns)
		database.DBCollections[k].SQLDB.SetMaxIdleConns(config.GetConfig().Mysql[k].MaxIdleConns)
		database.DBCollections[k].SQLDB.SetConnMaxLifetime(time.Duration(config.GetConfig().Mysql[k].ConnMaxLifetime))
	}

	// 数据库迁移
	database.DB().Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&user.User{})
}
