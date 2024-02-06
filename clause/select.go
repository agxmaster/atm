package clause

import (
	"gorm.io/gorm"
	"strings"
)

type Select []string

func (s Select) Build(db *gorm.DB) *gorm.DB {
	if s == nil {
		return db
	}
	return db.Select(strings.Join(s, ","))
}
