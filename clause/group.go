package clause

import (
	"gorm.io/gorm"
	"strings"
)

type Groups []string

func (g Groups) Build(db *gorm.DB) *gorm.DB {
	if g == nil {
		return db
	}
	return db.Group(strings.Join(g, ","))
}
