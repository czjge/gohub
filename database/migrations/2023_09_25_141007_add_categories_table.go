package migrations

import (
	"database/sql"

	"github.com/czjge/gohub/app/models"
	"github.com/czjge/gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Category struct {
		models.BaseModel

		Name        string `gorm:"type:varchar(255);not null;default:'';index;"`
		Description string `gorm:"type:varchar(255);not null;default:'';"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Category{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Category{})
	}

	migrate.Add("2023_09_25_141007_add_categories_table", up, down)
}
