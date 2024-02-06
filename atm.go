package atm

import (
	"context"
	"errors"
	"fmt"
	"github.com/agxmaster/atm/clause"
	"github.com/agxmaster/atm/util"
	"gorm.io/gorm"
	gclause "gorm.io/gorm/clause"
	"reflect"
)

var Conf *Config

type Config struct {
	AutoUpdatedAt             bool
	AutoUpdatedAtIgnoreTables []string

	AutoCreatedAt             bool
	AutoCreatedAtIgnoreTables []string

	AutoAuditLog       bool
	AutoAuditLogTables []string

	//AutoOperatorLog       bool
	//AutoOperatorLogTables []string
}

type Atm[T any] struct {
	Db *gorm.DB
}

func DefaultConfig() *Config {
	return &Config{
		AutoUpdatedAt: true,
		AutoCreatedAt: true,
		AutoAuditLog:  true,
	}
}

func (gorm *Atm[T]) Create(ctx context.Context, table string, data RowsMap) error {

	var err error
	err = gorm.Db.Table(table).Create(&data).Error
	return err
}

func (gorm *Atm[T]) BatchCreate(ctx context.Context, table string, data []RowsMap) error {
	var err error
	err = gorm.Db.Table(table).Create(&data).Error
	return err
}

func (gorm *Atm[T]) SCreate(ctx context.Context, data T) error {
	var err error
	err = gorm.Db.Clauses(gclause.Returning{}).Create(data).Error
	return err
}

func (gorm *Atm[T]) SBatchCreate(ctx context.Context, data []T) error {
	var err error
	err = gorm.Db.Clauses(gclause.Returning{}).Create(data).Error
	return err
}

func (gorm *Atm[T]) Delete(ctx context.Context, table string, id int64) error {
	err := gorm.Db.Exec(fmt.Sprintf("delete from %s where id = %d", table, id)).Error
	return err
}

func (gorm *Atm[T]) DeleteBatch(ctx context.Context, table string, ids []int64) error {
	err := gorm.Db.Exec(fmt.Sprintf("delete from %s where id in (%s)", table, util.ArrayToString(ids, ","))).Error
	return err
}

func (gorm *Atm[T]) SDelete(ctx context.Context, data T) error {
	err := gorm.Db.Delete(data).Error
	return err
}

func (gorm *Atm[T]) SDeleteBatch(ctx context.Context, ids []int64) error {
	err := gorm.Db.Delete(new(T), ids).Error
	return err
}

func (gorm *Atm[T]) Update(ctx context.Context, table string, id int64, data RowsMap) error {
	err := gorm.Db.Table(table).Where("id", id).Updates(data).Error
	return err
}

func (gorm *Atm[T]) BatchUpdates(ctx context.Context, table string, ids []int64, data RowsMap) error {
	err := gorm.Db.Table(table).Where("id in ?", ids).Updates(data).Error
	return err
}

func (gorm *Atm[T]) SUpdate(ctx context.Context, id int64, data T) error {
	err := gorm.Db.Model(new(T)).Where("id", id).Updates(&data).Error
	return err
}

func (gorm *Atm[T]) SBatchUpdates(ctx context.Context, ids []int64, data T) error {
	err := gorm.Db.Model(new(T)).Where("id in ?", ids).Save(&data).Error
	return err
}

func (gorm *Atm[T]) First(ctx context.Context, table string, id int64) (data RowsMap, err error) {
	err = gorm.Db.Table(table).Where("id", id).Take(&data).Error
	return data, err
}

func (gorm *Atm[T]) QueryPage(ctx context.Context, table string, clauses clause.Clauses) (ResultWithPage[T], error) {
	var data = new([]T)
	var resultWithPage ResultWithPage[T]

	db, err, count := gorm.query(ctx, table, clauses)

	if err != nil {
		return resultWithPage, err
	}
	err = db.Find(data).Error
	resultWithPage.Total = count
	resultWithPage.Data = *data

	return resultWithPage, err
}

func (gorm *Atm[T]) Query(ctx context.Context, table string, clauses clause.Clauses) ([]RowsMap, error) {
	var data []RowsMap

	db, err, _ := gorm.query(ctx, table, clauses)
	err = db.Scan(&data).Error
	return data, err
}

func (gorm *Atm[T]) SQueryPage(ctx context.Context, clauses clause.Clauses) (ResultWithPage[T], error) {
	var data []T
	var resultWithPage ResultWithPage[T]
	db, err, count := gorm.query(ctx, "", clauses)

	if err != nil {
		return resultWithPage, err
	}

	err = db.Find(&data).Error
	resultWithPage.Data = data
	resultWithPage.Total = count

	return resultWithPage, err
}

func (gorm *Atm[T]) SQuery(ctx context.Context, clauses clause.Clauses) ([]T, error) {
	var data []T

	db, err, _ := gorm.query(ctx, "", clauses)
	err = db.Find(&data).Error
	return data, err
}

func (gorm *Atm[T]) SFirst(ctx context.Context, id int64) (clause.ColumnMap, error) {
	var data clause.ColumnMap
	err := gorm.Db.Model(new(T)).Where("id", id).First(&data).Error
	return data, err
}

func (gorm *Atm[T]) query(ctx context.Context, table string, clauses clause.Clauses) (db *gorm.DB, err error, count int64) {

	db = gorm.Db
	if table != "" {
		db = db.Table(table)
	} else if reflect.TypeOf(new(T)).Elem().Kind() == reflect.Struct ||
		reflect.TypeOf(new(T)).Elem().Elem().Kind() == reflect.Struct {
		db = db.Model(new(T))
	} else {
		return db, errors.New("table and model not set"), 0
	}
	return clauses.Build(db)

}
