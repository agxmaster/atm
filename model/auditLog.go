package model

import (
	"time"
)

type AuditLog struct {
	ID        uint      `gorm:"column:id"`
	Table     string    `json:"table" gorm:"column:table_name"`
	SqlText   string    `json:"sql" gorm:"column:sql_text"`
	UserId    int64     `json:"user_id" gorm:"column:user_id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime(0);autoUpdateTime" json:"created_at"`
}

func (u *AuditLog) TableName() string {
	return "auditlog"
}
