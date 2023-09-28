package migrations

import (
	"database/sql"

	"github.com/czjge/gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type User struct {
		City         string `gorm:"type:varchar(10);not null;default:'';"`
		Introduction string `gorm:"type:varchar(255);not null;default:'';"`
		Avatar       string `gorm:"type:varchar(255);not null;default:'';"`
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&User{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropColumn(&User{}, "City")
		migrator.DropColumn(&User{}, "Introduction")
		migrator.DropColumn(&User{}, "Avatar")
	}

	migrate.Add("2023_09_28_091841_add_fields_to_user", up, down)
}
