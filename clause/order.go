package clause

import (
	"bytes"
	"fmt"
	"gorm.io/gorm"
)

type Order struct {
	Field string
	Desc  bool
}
type Orders []Order

func (os Orders) Build(db *gorm.DB) *gorm.DB {

	if os == nil || len(os) == 0 {
		return db
	}
	var buffer bytes.Buffer
	for index, order := range os {
		buffer.WriteString(fmt.Sprintf("%s ", order.Field))
		if order.Desc {
			buffer.WriteString("desc")
		}
		if index != len(os)-1 {
			buffer.WriteString(",")
		}
	}

	db = db.Order(buffer.String())
	return db
}
