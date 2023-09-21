package make

import (
	"fmt"

	"github.com/czjge/gohub/pkg/app"
	"github.com/czjge/gohub/pkg/console"
	"github.com/spf13/cobra"
)

var CmdMakeMigration = &cobra.Command{
	Use:   "migration",
	Short: "Create a migration file, example: make migration add_users_table user",
	Run:   runMakeMigration,
	Args:  cobra.ExactArgs(2),
}

func runMakeMigration(cmd *cobra.Command, args []string) {

	timeStr := app.TimenowInTimezone().Format("2006_01_02_150405")

	model := makeModelFromString(args[0])
	modelName := makeModelFromString(args[1])
	fileName := timeStr + "_" + model.PackageName
	filePath := fmt.Sprintf("database/migrations/%s.go", fileName)
	createFileFromStub(filePath, "migration", model, map[string]string{"{{FileName}}": fileName, "{{ModelName}}": modelName.StructName})
	console.Success("Migration file created，after modify it, use `migrate up` to migrate database.")
}
