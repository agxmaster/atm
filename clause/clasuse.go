package clause

import (
	"gorm.io/gorm"
	"reflect"
	"sort"
)

type Clause interface {
	Build(db *gorm.DB) *gorm.DB
}

type CountAble interface {
	Count(db *gorm.DB) (int64, error)
	Clause
}

type Scope func(*gorm.DB) *gorm.DB

func (s Scope) Build(db *gorm.DB) *gorm.DB {
	return db.Scopes(s)
}

type Clauses []Clause

var clauseSort = map[reflect.Type]int{
	reflect.TypeOf(Scope(nil)):      1,
	reflect.TypeOf(ColumnMap{}):     2,
	reflect.TypeOf(GormWhere{}):     3,
	reflect.TypeOf(PageWithGroup{}): 5,
	reflect.TypeOf(Page{}):          6,
	reflect.TypeOf(Groups{}):        4,
	reflect.TypeOf(Orders{}):        7,
	reflect.TypeOf(Select{}):        8,
}

func (cs Clauses) Len() int {
	return len(cs)
}

func (cs Clauses) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}
func (cs Clauses) Less(i, j int) bool {
	weightLeft, _ := clauseSort[reflect.TypeOf(cs[i])]
	weightRight, _ := clauseSort[reflect.TypeOf(cs[j])]
	return weightLeft < weightRight
}

func (cs Clauses) Build(db *gorm.DB) (scope *gorm.DB, err error, count int64) {
	scope = db
	sort.Sort(cs)

	for _, c := range cs {
		if countAble, ok := c.(CountAble); ok {
			count, err = countAble.Count(scope)
		}
		scope = c.Build(scope)
	}
	return scope, err, count
}
