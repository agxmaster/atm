package clause

import "gorm.io/gorm"

type Logical string

const LogicalAnd Logical = "and"
const LogicalOr Logical = "or"

type GormWhere map[Logical]ColumnMap

func (gw GormWhere) Build(db *gorm.DB) *gorm.DB {
	if gw == nil {
		return db
	}

	for logical, condition := range gw {
		if logical == LogicalAnd {
			for key, val := range condition {
				db = db.Where(key, val)
			}
		}

		if logical == LogicalOr {
			for key, val := range condition {
				db = db.Or(key, val)
			}
		}
	}
	return db
}
