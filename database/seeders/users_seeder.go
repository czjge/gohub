package seeders

import (
	"fmt"

	"github.com/czjge/gohub/database/factories"
	"github.com/czjge/gohub/pkg/console"
	"github.com/czjge/gohub/pkg/logger"
	"github.com/czjge/gohub/pkg/seed"
	"gorm.io/gorm"
)

func init() {

	seed.Add("SeedUsersTable", func(db *gorm.DB) {

		users := factories.MakeUsers(10)

		result := db.Table("users").Create(&users)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
