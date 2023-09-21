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

func CurrentDatabase(name ...string) (dbname string) {
	if len(name) > 0 {
		dbname = DB(name[0]).Migrator().CurrentDatabase()
	} else {
		dbname = DB().Migrator().CurrentDatabase()
	}
	return
}

func DeleteAllTables(name ...string) error {
	var err error
	dbname := CurrentDatabase(name...)
	tables := []string{}

	// 读取所有数据表
	err = DB(name...).Table("information_schema.tables").
		Where("table_schema = ?", dbname).
		Pluck("table_name", &tables).
		Error
	if err != nil {
		return err
	}

	// 暂时关闭外键检测
	DB(name...).Exec("SET foreign_key_checks = 0;")

	// 删除所有表
	for _, table := range tables {
		err := DB(name...).Migrator().DropTable(table)
		if err != nil {
			return err
		}
	}

	// 开启 MySQL 外键检测
	DB(name...).Exec("SET foreign_key_checks = 1;")
	return nil
}
