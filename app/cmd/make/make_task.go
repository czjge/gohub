package make

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CmdMakeTask = &cobra.Command{
	Use:   "task",
	Short: "Create task file, example: make task buy",
	Run:   runMakeTask,
	Args:  cobra.ExactArgs(1),
}

func runMakeTask(cmd *cobra.Command, args []string) {

	model := makeModelFromString(args[0])

	filePath := fmt.Sprintf("app/tasks/%s_task.go", model.PackageName)

	createFileFromStub(filePath, "task", model)
}
