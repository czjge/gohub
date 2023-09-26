package factories

import (
	"github.com/czjge/gohub/app/models/link"

	"github.com/go-faker/faker/v4"
)

func MakeLinks(count int) []link.Link {

	var objs []link.Link

	for i := 0; i < count; i++ {
		linkModel := link.Link{
			Name: faker.Username(),
			URL:  faker.URL(),
		}
		objs = append(objs, linkModel)

		// 设置唯一性，如 Link 模型的某个字段需要唯一，即可取消注释
		faker.ResetUnique()
	}

	return objs
}
