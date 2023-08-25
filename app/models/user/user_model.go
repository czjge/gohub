package user

import "github.com/czjge/gohub/app/models"

type User struct {
	models.BaseModel

	Name     string `json:"name,omitempty"` // omitempty:在结构体转json过程中，把没有赋值的结构体属性不在json中输出
	Email    string `json:"-"`              // json解析器忽略该字段，不输出给用户
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}
