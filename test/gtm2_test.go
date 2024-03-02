package test

import (
	"encoding/json"
	"fmt"
	"github.com/agxmaster/atm"
	"github.com/agxmaster/atm/clause"
	"gorm.io/gorm"
	"reflect"
	"testing"
	"time"
)

var r reflect.Type

func init() {
	r = reflect.TypeOf(Student{})
}

func TestBatchCreate2(t *testing.T) {
	u := []Student{
		{
			Name:      "1",
			Gender:    1,
			Age:       1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Class:     "aaa",
		},
		{
			Name:      "2",
			Gender:    1,
			Age:       1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Class:     "aaa",
		},
	}
	d, _ := json.Marshal(u)
	rt := reflect.TypeOf(Student{})
	i := reflect.New(reflect.SliceOf(rt)).Interface()
	m := reflect.New(rt).Interface()

	//it := reflect.New(rt).Interface()

	json.Unmarshal(d, &i)
	fmt.Printf("%#v", i)
	err := atm.BatchCreate(ctx, gormdb, m, i)

	if err != nil {
		t.Errorf("%+v", err)
	}
}

func TestCreate2(t *testing.T) {
	u := Student{

		Name:      "1",
		Gender:    1,
		Age:       1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Class:     "aaa",
	}
	d, _ := json.Marshal(u)

	rt := reflect.TypeOf(Student{})
	it := reflect.New(rt).Interface()
	json.Unmarshal(d, &it)
	fmt.Printf("%#v", it)
	err := atm.Create(ctx, gormdb, it, it)

	if err != nil {
		t.Errorf("%+v", err)
	}
}

func TestCreateFunc(t *testing.T) {
	m := reflect.New(r).Interface()
	err := atm.Create(ctx, gormdb, m, m)
	if err != nil {
		t.Errorf("TestCreateFunc %+v", err)
	}
}

func TestQueryPageFunc(t *testing.T) {
	data, err := atm.QueryPage(ctx, gormdb, r, []clause.Clause{clause.Page{NeedCount: true}}, false)
	fmt.Printf("%+v", data)
	if err != nil {
		t.Errorf("TestCreateFunc %+v", err)
	}
}

func TestFirstFunc(t *testing.T) {
	model := reflect.New(r).Interface()
	data, err := atm.First(ctx, gormdb, model, 1, []string{"id"})
	fmt.Printf("%+v", data)
	if err != nil {
		t.Errorf("TestCreateFunc %+v", err)
	}
}

func TestPageWithGroup(t *testing.T) {
	data, err := atm.QueryPage(ctx, gormdb, r, []clause.Clause{
		clause.PageWithGroup{
			Page:   clause.Page{NeedCount: true, PageSize: 3, PageNum: 1},
			Groups: []string{"name", "age"},
		},
		clause.Scope(func(db *gorm.DB) *gorm.DB {
			return db.Select("name,age")
		}),
	}, true)
	fmt.Printf("%+v", data.Data)
	if err != nil {
		t.Errorf("TestCreateFunc %+v", err)
	}
}
func TestPageWithGroup2(t *testing.T) {
	var count int64
	err := gormdb.Model(&Student{}).Select("count(distinct name,age)").Group("name").Scan(&count).Error

	fmt.Printf("%+v ,%+v", count, err)

	sql := gormdb.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return gormdb.Model(&Student{}).Select("count(distinct name,age)").Group("name").Scan(&count)
	})
	fmt.Printf("sss %s", sql)

	if err != nil {
		t.Errorf("TestCreateFunc %+v", err)
	}
}

func TestQueryPage3(t *testing.T) {
	db := &atm.Atm[atm.RowsMap]{Db: gormdb}

	res, err := db.QueryPage(ctx, TableUser,
		[]clause.Clause{
			clause.ColumnMap{"class": 1},
			clause.Page{PageNum: 1, PageSize: 2, NeedCount: true},
		})
	fmt.Printf("%+v \n", res)
	if err != nil {
		t.Errorf("QueryPage error %+v", err)
	}
}
