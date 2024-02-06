package clause

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type PageWithGroup struct {
	Page   Page
	Groups Groups
}

func (p PageWithGroup) Count(db *gorm.DB) (int64, error) {
	if !p.Page.NeedCount {
		return 0, nil
	}

	if p.Groups == nil || len(p.Groups) == 0 {
		return p.Page.Count(db)
	}
	var count int64
	err := db.Select(fmt.Sprintf("count(distinct %s) as count", strings.Join(p.Groups, ","))).Scan(&count).Error
	return count, err
}

func (p PageWithGroup) Build(db *gorm.DB) *gorm.DB {
	return p.Groups.Build(p.Page.Build(db))
}
