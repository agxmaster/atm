package clause

import "gorm.io/gorm"

const DefaultPageSize = 10

type Page struct {
	PageNum   int
	PageSize  int
	NeedCount bool
}

func (p Page) Build(db *gorm.DB) *gorm.DB {

	if p.PageSize <= 0 {
		p.PageSize = DefaultPageSize
	}

	if p.PageNum <= 0 {
		p.PageNum = 1
	}

	return db.Offset((p.PageNum - 1) * p.PageSize).Limit(p.PageSize)
}

func (p Page) Count(db *gorm.DB) (int64, error) {
	if !p.NeedCount {
		return 0, nil
	}
	var count int64
	err := db.Count(&count).Error
	return count, err
}
