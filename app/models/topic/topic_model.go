// Package topic 模型
package topic

import (
	"github.com/czjge/gohub/app/models"
	"github.com/czjge/gohub/app/models/category"
	"github.com/czjge/gohub/app/models/user"
	"github.com/czjge/gohub/pkg/database"
)

type Topic struct {
	models.BaseModel

	Title      string `json:"title,omitempty"`
	Body       string `json:"body,omitempty"`
	UserID     string `json:"user_id,omitempty"`
	CategoryID string `json:"category_id,omitempty"`

	User     user.User         `json:"user"`
	Category category.Category `json:"category"`

	models.CommonTimestampsField
}

func (topic *Topic) Create() {
	database.DB().Create(&topic)
}

func (topic *Topic) Save() (rowsAffected int64) {
	result := database.DB().Save(&topic)
	return result.RowsAffected
}

func (topic *Topic) Delete() (rowsAffected int64) {
	result := database.DB().Delete(&topic)
	return result.RowsAffected
}
