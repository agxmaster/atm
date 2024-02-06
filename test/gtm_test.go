package test

import (
	"context"
	"fmt"
	"github.com/agxmaster/atm"
	"github.com/agxmaster/atm/callback"
	"github.com/agxmaster/atm/clause"
	"github.com/agxmaster/atm/model"
	"github.com/agxmaster/atm/plugs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"testing"
	"time"
)

const TableUser = "student"

const UserIdKey = "userIdKey"

var gormdb *gorm.DB
var ctx = context.Background()
var err error

type Student struct {
	Id        int64     `gorm:"column:id" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Gender    int       `gorm:"column:gender" json:"gender"`
	Age       int       `gorm:"column:age" json:"age"`
	Class     string    `gorm:"column:class" json:"class"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime(0);autoUpdateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime(0);autoUpdateTime" json:"updated_at"`
}

func (s *Student) TableName() string {
	return "student"
}

type Student2 struct {
	*model.OperatorTime
	Id     int64  `gorm:"column:id" json:"id"`
	Name   string `gorm:"column:name" json:"name"`
	Gender int    `gorm:"column:gender" json:"gender"`
	Age    int    `gorm:"column:age" json:"age"`
	Class  string `gorm:"column:class" json:"class"`
}

var u = Student2{
	OperatorTime: &model.OperatorTime{CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

func TestGetSql(T *testing.T) {
	sql := `
	-- run this sql to db gorm
	CREATE TABLE student (
	  id bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
	  name varchar(128) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' ,
	  gender int NOT NULL DEFAULT '0' ,
	  age int NOT NULL DEFAULT '0',
	  class varchar(128) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
	  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ,
	  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	  PRIMARY KEY (id)
	) ENGINE=InnoDB AUTO_INCREMENT=41751 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
	`
	fmt.Println(sql)
}

func TestCreate(t *testing.T) {

	data := map[string]interface{}{
		//"id":         nil,
		"name":       "name1",
		"gender":     100,
		"age":        100,
		"class":      "aaa",
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}
	//需要提前设置userid auditlog用
	gormdb = plugs.AuditLogUserIdStoreInstance.SetUserId(ctx, gormdb, 23)

	//gormdb = gormdb.Set(UserIdKey, 22)

	db := &atm.Atm[interface{}]{Db: gormdb}
	err := db.Create(ctx, TableUser, data)
	if err != nil {
		t.Errorf("Create error %+v", err)
	}
}

func TestBatchCreate(t *testing.T) {
	data := []map[string]interface{}{
		{
			//"id":         nil,
			"name":       "name1",
			"gender":     100,
			"age":        100,
			"class":      "aaa",
			"created_at": time.DateTime,
			"updated_at": time.DateTime,
		},
		{
			//"id":         nil,
			"name":       "name2",
			"gender":     100,
			"age":        100,
			"class":      "aaa",
			"created_at": time.DateTime,
			"updated_at": time.DateTime,
		},
	}

	db := &atm.Atm[interface{}]{Db: gormdb}
	err := db.BatchCreate(ctx, TableUser, data)
	if err != nil {
		t.Errorf("BatchCreate error %+v", err)
	}
}

func TestSCreate(t *testing.T) {
	u := Student{
		Name:      "1",
		Gender:    1,
		Age:       1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Class:     "aaa",
	}

	db := &atm.Atm[Student]{Db: gormdb}
	err := db.SCreate(ctx, u)
	if err != nil {
		t.Errorf("SCreate error %+v", err)
	}
}

func TestSBatchCreate(t *testing.T) {
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

	db := &atm.Atm[Student]{Db: gormdb}
	err := db.SBatchCreate(ctx, u)
	if err != nil {
		t.Errorf("SBatchCreate error %+v", err)
	}
}

func TestDelete(t *testing.T) {

	db := &atm.Atm[interface{}]{Db: gormdb}
	err := db.Delete(ctx, TableUser, 10)
	if err != nil {
		t.Errorf("Delete error %+v", err)
	}
}

func TestDeleteBatch(t *testing.T) {
	db := &atm.Atm[interface{}]{Db: gormdb}
	err := db.DeleteBatch(ctx, TableUser, []int64{10, 11})
	if err != nil {
		t.Errorf("DeleteBatch error %+v", err)
	}
}

func TestSDelete(t *testing.T) {
	db := &atm.Atm[Student]{Db: gormdb}
	err := db.SDelete(ctx, Student{Name: "sss"})
	if err != nil {
		t.Errorf("SDelete error %+v", err)
	}
}

func TestSDeleteBatch(t *testing.T) {
	db := &atm.Atm[Student]{Db: gormdb}
	err := db.SDeleteBatch(ctx, []int64{1, 2})
	if err != nil {
		t.Errorf("SDeleteBatch error %+v", err)
	}
}

func TestUpdate(t *testing.T) {
	db := &atm.Atm[interface{}]{Db: gormdb}
	err := db.Update(ctx, TableUser, 68, atm.RowsMap{"name": "merry"})
	if err != nil {
		t.Errorf("Update error %+v", err)
	}
}

func TestBatchUpdates(t *testing.T) {
	db := &atm.Atm[interface{}]{Db: gormdb}
	err := db.BatchUpdates(ctx, TableUser, []int64{67}, atm.RowsMap{"name": "merry"})
	if err != nil {
		t.Errorf("BatchUpdates error %+v", err)
	}
}

func TestSUpdate(t *testing.T) {
	db := &atm.Atm[Student]{Db: gormdb}
	err := db.SUpdate(ctx, 1, Student{Name: "cherry"})
	if err != nil {
		t.Errorf("SUpdate error %+v", err)
	}
}

func TestSBatchUpdates(t *testing.T) {
	db := &atm.Atm[Student]{Db: gormdb}
	err := db.SBatchUpdates(ctx, []int64{1, 2}, Student{Id: 69, Name: "sorry"})
	if err != nil {
		t.Errorf("SBatchUpdates error %+v", err)
	}
}

func TestFirst(t *testing.T) {
	db := &atm.Atm[atm.RowsMap]{Db: gormdb}
	res, err := db.First(ctx, TableUser, 69)
	fmt.Printf("%+v \n", res)
	if err != nil {
		t.Errorf("First error %+v", err)
	}
}

func TestQuery(t *testing.T) {
	db := &atm.Atm[atm.RowsMap]{Db: gormdb}

	//SELECT * FROM `Student` WHERE `id` = 1 AND id > 1 OR id < 20 GROUP BY `id` ORDER BY id desc
	res, err := db.Query(ctx, TableUser,
		[]clause.Clause{
			clause.ColumnMap{"id": 1},
			clause.GormWhere{clause.LogicalAnd: {"id > ?": 1}, clause.LogicalOr: {"id < ?": 20}},
			clause.Orders{clause.Order{Field: "id", Desc: true}},
			clause.Groups{"id"},
		})
	fmt.Printf("%+v \n", res)
	if err != nil {
		t.Errorf("Query error %+v", err)
	}
}

func TestQuery2(t *testing.T) {
	db := &atm.Atm[atm.RowsMap]{Db: gormdb}

	//SELECT * FROM `Student` WHERE `id` = 1 AND id > 1 OR id < 20 AND id != 1 GROUP BY `id` ORDER BY id desc
	res, err := db.Query(ctx, TableUser,
		[]clause.Clause{
			clause.ColumnMap{"id": 1},
			clause.GormWhere{clause.LogicalAnd: {"id > ?": 1}, clause.LogicalOr: {"id < ?": 20}},
			clause.Scope(func(db *gorm.DB) *gorm.DB {
				return db.Where("id != ?", 1)
			}),
			clause.Orders{clause.Order{Field: "id", Desc: true}},
			clause.Groups{"id"},
		})
	fmt.Printf("%+v", res)
	if err != nil {
		t.Errorf("Query error %+v", err)
	}
}

func TestQueryPage(t *testing.T) {
	db := &atm.Atm[atm.RowsMap]{Db: gormdb}

	res, err := db.QueryPage(ctx, TableUser,
		[]clause.Clause{
			clause.ColumnMap{"id": 1},
			clause.GormWhere{clause.LogicalAnd: {"id > ?": 1}, clause.LogicalOr: {"id < ?": 20}},
			clause.Scope(func(db *gorm.DB) *gorm.DB {
				return db.Where("id != ?", 1)
			}),
			clause.Orders{clause.Order{Field: "id", Desc: true}},
			clause.Groups{"id"},
			clause.Page{PageNum: 1, PageSize: 2, NeedCount: true},
		})
	fmt.Printf("%+v \n", res)
	if err != nil {
		t.Errorf("QueryPage error %+v", err)
	}
}

func TestQueryPage2(t *testing.T) {
	db := &atm.Atm[atm.RowsMap]{Db: gormdb}

	res, err := db.QueryPage(ctx, TableUser,
		[]clause.Clause{
			clause.Page{PageNum: 1, PageSize: 2},
		})
	fmt.Printf("%+v \n", res)
	if err != nil {
		t.Errorf("QueryPage error %+v", err)
	}
}

// query中 page.NeedCount 无效 因为返回的结构中不带总条数
func TestSQuery(t *testing.T) {
	db := &atm.Atm[Student]{Db: gormdb}

	res, err := db.SQuery(ctx,
		[]clause.Clause{
			clause.ColumnMap{"id": 1},
			clause.GormWhere{clause.LogicalAnd: {"id > ?": 1}, clause.LogicalOr: {"id < ?": 20}},
			clause.Scope(func(db *gorm.DB) *gorm.DB {
				return db.Where("id != ?", 1)
			}),
			clause.Orders{clause.Order{Field: "id", Desc: true}},
			clause.Groups{"id"},
			clause.Page{PageNum: 1, PageSize: 2, NeedCount: true},
		})
	fmt.Printf("%+v \n", res)
	if err != nil {
		t.Errorf("SQuery error %+v", err)
	}
}

func TestSQueryPage(t *testing.T) {
	db := &atm.Atm[Student]{Db: gormdb}

	res, err := db.SQueryPage(ctx,
		[]clause.Clause{
			clause.ColumnMap{"id": 1},
			clause.GormWhere{clause.LogicalAnd: {"id > ?": 1}, clause.LogicalOr: {"id < ?": 20}},
			clause.Scope(func(db *gorm.DB) *gorm.DB {
				return db.Where("id != ?", 1)
			}),
			clause.Orders{clause.Order{Field: "id", Desc: true}},
			clause.Groups{"id"},
			clause.Page{PageNum: 1, PageSize: 2, NeedCount: true},
		})
	fmt.Printf("%+v \n", res)
	if err != nil {
		t.Errorf("SQueryPage error %+v", err)
	}
}

func TestAddCreateAt(t *testing.T) {

	data := map[string]interface{}{
		//"id":         nil,
		"name":   "name1",
		"gender": 100,
		"age":    100,
		"class":  "aaa",
		//"created_at": time.DateTime,
		//"updated_at": time.DateTime,
	}

	gormdb.Callback().Create().Before("gorm:before_create").Register("atmCreateBefore", func(db *gorm.DB) {

		switch data := db.Statement.Dest.(type) {
		case *atm.RowsMap:
			if _, ok := (*data)["created_at"]; !ok {
				(*data)["created_at"] = time.Now()
			}

			if _, ok := (*data)["updated_at"]; !ok {
				(*data)["updated_at"] = time.Now()
			}
		}
	})

	db := &atm.Atm[interface{}]{Db: gormdb}
	err := db.Create(ctx, TableUser, data)
	if err != nil {
		t.Errorf("Create error %+v", err)
	}
}

func TestAddCreateAtBatch(t *testing.T) {

	data := []map[string]interface{}{
		{
			//"id":         nil,
			"name":   "name1",
			"gender": 100,
			"age":    100,
			"class":  "aaa",
			//"created_at": time.DateTime,
			//"updated_at": time.DateTime,
		},
		{
			//"id":         nil,
			"name":   "name2",
			"gender": 100,
			"age":    100,
			"class":  "aaa",
			//"created_at": time.DateTime,
			//"updated_at": time.DateTime,
		},
	}

	gormdb.Callback().Create().Before("gorm:before_create").Register("atmCreateBefore", func(db *gorm.DB) {

		switch data := db.Statement.Dest.(type) {
		case *[]atm.RowsMap:
			for k := range *data {
				if _, ok := (*data)[k]["created_at"]; !ok {
					(*data)[k]["created_at"] = time.Now()
				}
				if _, ok := (*data)[k]["updated_at"]; !ok {
					(*data)[k]["updated_at"] = time.Now()
				}
			}
		}
	})

	db := &atm.Atm[interface{}]{Db: gormdb}
	err := db.BatchCreate(ctx, TableUser, data)
	if err != nil {
		t.Errorf("Create error %+v", err)
	}
}

func TestAddUpdateAt(t *testing.T) {
	db := &atm.Atm[interface{}]{Db: gormdb}

	gormdb.Callback().Update().Before("gorm:before_update").Register("atmUpdateBefore", func(db *gorm.DB) {

		switch data := db.Statement.Dest.(type) {
		case *atm.RowsMap:
			if _, ok := (*data)["updated_at"]; !ok {
				(*data)["updated_at"] = time.Now()
			}

		case atm.RowsMap:
			if _, ok := data["updated_at"]; !ok {
				data["updated_at"] = time.Now()
			}
		}

	})

	err := db.BatchUpdates(ctx, TableUser, []int64{67}, atm.RowsMap{"name": "merry"})
	if err != nil {
		t.Errorf("Create error %+v", err)
	}
}

func TestSqlLog(t *testing.T) {
	db := &atm.Atm[interface{}]{Db: gormdb}

	gormdb.Callback().Update().Before("gorm:after_update").Register("atmUpdateBefore", func(db *gorm.DB) {
		sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
			return db
		})
		fmt.Println(sql)
	})

	err := db.BatchUpdates(ctx, TableUser, []int64{67}, atm.RowsMap{"name": "merry"})
	if err != nil {
		t.Errorf("Create error %+v", err)
	}
}

func setup() {
	var dsn = "root:12345678@tcp(localhost:3306)/atm?charset=utf8&parseTime=True&loc=Local"
	gormdb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})

	atm.Conf = atm.DefaultConfig()

	callback.Regist(context.Background(), *gormdb, atm.Conf)

	if err != nil {
		panic(err)
	}
}

func teardown() {
	fmt.Println("test over ")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
