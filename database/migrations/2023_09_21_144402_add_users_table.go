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

		Name     string `gorm:"type:varchar(50);not null;default:'';comment:姓名;"`
		Email    string `gorm:"type:varchar(50);not null;default:'';index;comment:邮箱;"`
		Phone    string `gorm:"type:varchar(50);not null;default:'';index;comment:手机号;"`
		Password string `gorm:"type:varchar(60);not null;default:'';comment:密码;"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&User{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&User{})
	}

	migrate.Add("2023_09_21_144402_add_users_table", up, down)
}
