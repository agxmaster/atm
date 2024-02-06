package model

import "time"

type OperatorTime struct {
	CreatedAt time.Time `gorm:"column:created_at;type:datetime(0);autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime(0);autoUpdateTime" json:"updated_at"`
}
