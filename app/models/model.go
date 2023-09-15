package models

import (
	"time"

	"github.com/spf13/cast"
)

// 模型基类
type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
}

// gorm 默认使用 createAt 和 updateAt 追踪创建时间和更新时间
type CommonTimestampsField struct {
	CreatedAt time.Time `gorm:"column:created_at;index;" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;index;" json:"updated_at,omitempty"`
}

func (a BaseModel) GetStringID() string {
	return cast.ToString(a.ID)
}
