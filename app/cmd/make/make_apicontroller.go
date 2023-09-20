package make

import (
	"fmt"
	"os"
	"strings"

	"github.com/czjge/gohub/pkg/console"
	"github.com/spf13/cobra"
)

var CmdMakeAPIController = &cobra.Command{
	Use:   "apicontroller",
	Short: "Create api controller，exmaple: make apicontroller v1/user",
	Run:   runMakeAPIController,
	Args:  cobra.ExactArgs(1),
}

func runMakeAPIController(cmd *cobra.Command, args []string) {

	// 处理参数，要求附带 API 版本（v1 或者 v2）
	array := strings.Split(args[0], "/")
	if len(array) != 2 {
		console.Exit("api controller name format: v1/user")
	}

	apiVersion, name := array[0], array[1]
	model := makeModelFromString(name)

	dirPath := fmt.Sprintf("app/http/controllers/api/%s/", apiVersion)
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		console.Exit(err.Error())
	}

	fileName := fmt.Sprintf("%s_controller.go", model.TableName)

	createFileFromStub(dirPath+fileName, "apicontroller", model, map[string]string{"{{version}}": apiVersion})
}
