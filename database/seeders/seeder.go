package seeders

import "github.com/czjge/gohub/pkg/seed"

func Initialize() {

	// 触发加载本目录下其他文件中的 init 方法

	seed.SetRunOrder([]string{
		"SeedUsersTable",
	})
}
