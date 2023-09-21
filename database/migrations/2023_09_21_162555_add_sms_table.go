package migrations

import (
	"database/sql"

	"github.com/czjge/gohub/app/models"
	"github.com/czjge/gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Sms struct {
		models.BaseModel

		Phone         string `gorm:"type:char(11);comment:手机号;default:'';index:'idx_phone';"`
		SignName      string `gorm:"type:varchar(255);comment:短信签名;default:'';"`
		TemplateCode  string `gorm:"type:varchar(100);comment:短信模板;default:'';"`
		TemplateParam string `gorm:"type:text;comment:短信模板参数;default:null;"`
		RequestId     string `gorm:"type:varchar(100);comment:请求ID;default:'';uniqueIndex;"`
		BizId         string `gorm:"type:varchar(100);comment:发送回执ID;default:'';"`
		Code          string `gorm:"type:varchar(100);comment:请求状态码;default:'';"`
		Message       string `gorm:"type:varchar(100);comment:状态码描述;default:'';"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Sms{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Sms{})
	}

	migrate.Add("2023_09_21_162555_add_sms_table", up, down)
}
