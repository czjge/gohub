package database

import (
	"database/sql"

	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

type DBInfo struct {
	DB    *gorm.DB
	SQLDB *sql.DB
}

var DBCollections map[string]*DBInfo

func Connect(dbConfig gorm.Dialector, _logger gormlogger.Interface) (*gorm.DB, *sql.DB, error) {
	DB, err := gorm.Open(dbConfig, &gorm.Config{
		Logger: _logger,
	})
	if err != nil {
		return nil, nil, err
	}

	SQLDB, err := DB.DB()
	if err != nil {
		return nil, nil, err
	}
	return DB, SQLDB, nil
}

func DB(name ...string) *gorm.DB {
	if len(name) > 0 {
		if collect, ok := DBCollections[name[0]]; ok {
			return collect.DB
		}
		return nil
	}
	return DBCollections["default"].DB
}

func SQLDB(name ...string) *sql.DB {
	if len(name) > 0 {
		if collect, ok := DBCollections[name[0]]; ok {
			return collect.SQLDB
		}
		return nil
	}
	return DBCollections["default"].SQLDB
}
