package atm

import (
	"context"
	"errors"
	"github.com/agxmaster/atm/clause"
	"gorm.io/gorm"
	gclause "gorm.io/gorm/clause"
	"reflect"
)

type ResultWithPageAny struct {
	Total int64 `json:"total"`
	Data  any   `json:"data"`
}

func Create(ctx context.Context, db *gorm.DB, model any, data any) error {
	var err error
	err = db.Clauses(gclause.Returning{}).Model(model).Create(data).Error
	return err
}

func BatchCreate(ctx context.Context, db *gorm.DB, model any, data any) error {
	var err error
	err = db.Clauses(gclause.Returning{}).Model(model).Create(data).Error
	return err
}

func Delete(ctx context.Context, db *gorm.DB, data any) error {
	err := db.Delete(data).Error
	return err
}

func DeleteBatch(ctx context.Context, db *gorm.DB, model any, ids []int64) error {
	err := db.Delete(model, ids).Error
	return err
}

func Update(ctx context.Context, db *gorm.DB, model any, ids int64, data any) error {
	err := db.Model(model).Where("id", ids).Updates(&data).Error
	return err
}

func BatchUpdates(ctx context.Context, db *gorm.DB, model any, ids []int64, data any) error {
	err := db.Model(model).Where("id in ?", ids).Updates(&data).Error
	return err
}

func Query(ctx context.Context, db *gorm.DB, modelType reflect.Type, clauses clause.Clauses) (any, error) {
	model := reflect.New(modelType).Interface()
	res := reflect.New(reflect.SliceOf(modelType)).Interface()

	db, err, _ := query(ctx, db, "", model, clauses)
	err = db.Find(&res).Error
	return res, err
}

func First(ctx context.Context, db *gorm.DB, model any, id int64, columns []string) (any, error) {

	db = db.Model(model).Where("id", id)

	if columns != nil {
		db = clause.Select(columns).Build(db)
	}
	err := db.First(&model).Error
	return model, err
}

func QueryPage(ctx context.Context, db *gorm.DB, modelType reflect.Type, clauses clause.Clauses, destStruct bool) (ResultWithPageAny, error) {
	var (
		resultWithPage ResultWithPageAny
		res            interface{}
	)

	model := reflect.New(modelType).Interface()

	if destStruct {
		res = reflect.New(reflect.SliceOf(modelType)).Interface()
	} else {
		res = new([]map[string]interface{})
	}

	db, err, count := query(ctx, db, "", model, clauses)
	if err != nil {
		return resultWithPage, err
	}

	err = db.Find(res).Error
	if err != nil {
		return resultWithPage, err
	}

	resultWithPage.Data = res
	resultWithPage.Total = count

	return resultWithPage, nil
}

func query(ctx context.Context, db *gorm.DB, table string, model any, clauses clause.Clauses) (scope *gorm.DB, err error, count int64) {

	if table != "" {
		scope = db.Table(table)
	}
	if reflect.TypeOf(model).Elem().Kind() == reflect.Struct ||
		reflect.TypeOf(model).Elem().Elem().Kind() == reflect.Struct {
		scope = db.Model(model)
	}
	if table != "" && scope.Statement.Model == nil {
		return db, errors.New("table and model not set"), 0
	}
	return clauses.Build(scope)
}
