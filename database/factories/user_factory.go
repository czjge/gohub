package factories

import (
	"github.com/czjge/gohub/app/models/user"
	"github.com/czjge/gohub/pkg/helpers"
	"github.com/go-faker/faker/v4"
)

func MakeUsers(times int) []user.User {

	var objs []user.User

	for i := 0; i < times; i++ {
		model := user.User{
			Name:     faker.Username(),
			Email:    faker.Email(),
			Phone:    helpers.RandomNumber(11),
			Password: "$2a$14$Epc8Uwjuus7rB1ctD3LKweVLvyUfbG0B1CRiPO0yCwtAwx9Nv/7dG",
		}
		objs = append(objs, model)

		// 不生成重复数据
		faker.ResetUnique()
	}

	return objs
}
