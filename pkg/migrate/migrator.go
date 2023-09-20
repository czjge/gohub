package migrate

import (
	"github.com/czjge/gohub/pkg/database"
	"gorm.io/gorm"
)

// 数据迁移操作类
type Migrator struct {
	Folder   string
	DB       *gorm.DB
	Migrator gorm.Migrator
}

// 对应数据的 migrations 表里的一条数据
type Migration struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;"`
	Migration string `gorm:"type:varchar(255);not null;unique;default:'';"`
	Batch     int
}

// 创建 Migrator 实例，用以执行迁移操作
func NewMigrator() *Migrator {

	migrator := &Migrator{
		Folder:   "database/migrations/",
		DB:       database.DB(),
		Migrator: database.DB().Migrator(),
	}

	migrator.createMigrationTable()

	return migrator
}

// 创建 migrations 表
func (migrator *Migrator) createMigrationTable() {

	migration := Migration{}

	if !migrator.Migrator.HasTable(&migration) {
		migrator.Migrator.CreateTable(&migration)
	}
}
