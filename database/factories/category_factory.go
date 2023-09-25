package factories

import (
	"github.com/czjge/gohub/app/models/category"

	"github.com/go-faker/faker/v4"
)

func MakeCategories(count int) []category.Category {

	var objs []category.Category

	for i := 0; i < count; i++ {
		categoryModel := category.Category{
			Name:        faker.Username(),
			Description: faker.Sentence(),
		}
		objs = append(objs, categoryModel)

		// 设置唯一性，如 Category 模型的某个字段需要唯一，即可取消注释
		faker.ResetUnique()
	}

	return objs
}
