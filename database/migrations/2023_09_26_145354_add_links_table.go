package migrations

import (
	"database/sql"

	"github.com/czjge/gohub/app/models"
	"github.com/czjge/gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Link struct {
		models.BaseModel

		Name string `gorm:"type:varchar(255);not null;default:'';"`
		URL  string `gorm:"type:varchar(255);not null;default:'';"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Link{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Link{})
	}

	migrate.Add("2023_09_26_145354_add_links_table", up, down)
}
