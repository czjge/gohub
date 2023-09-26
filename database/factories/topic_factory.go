package factories

import (
	"github.com/czjge/gohub/app/models/topic"

	"github.com/go-faker/faker/v4"
)

func MakeTopics(count int) []topic.Topic {

	var objs []topic.Topic

	for i := 0; i < count; i++ {
		topicModel := topic.Topic{
			Title:      faker.Sentence(),
			Body:       faker.Paragraph(),
			CategoryID: "3",
			UserID:     "1",
		}
		objs = append(objs, topicModel)

		// 设置唯一性，如 Topic 模型的某个字段需要唯一，即可取消注释
		faker.ResetUnique()
	}

	return objs
}
