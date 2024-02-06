package clause

import "gorm.io/gorm"

type ColumnMap map[string]interface{}

func (cm ColumnMap) Build(db *gorm.DB) *gorm.DB {
	if cm == nil {
		return db
	}
	return db.Where(map[string]interface{}(cm))
}
