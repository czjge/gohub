package migrations

import (
	"database/sql"

	"github.com/czjge/gohub/app/models"
	"github.com/czjge/gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type User struct {
		models.BaseModel
	}

	type Category struct {
		models.BaseModel
	}

	type Topic struct {
		models.BaseModel

		Title      string `gorm:"type:varchar(255);not null;default:'';index;"`
		Body       string `gorm:"type:longtext;not null;"`
		UserID     string `gorm:"type:bigint;not null;default:0;index;"`
		CategoryID string `gorm:"type:bigint;not null;default:0;index;"`

		// 外键
		User     User
		Category Category

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Topic{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Topic{})
	}

	migrate.Add("2023_09_25_155116_add_topics_table", up, down)
}
