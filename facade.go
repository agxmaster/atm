package atm

import (
	"context"
	"github.com/agxmaster/atm/clause"
	"gorm.io/gorm"
)

type M struct {
	DB *gorm.DB
}

func (m *M) Create(ctx context.Context, table string, data map[string]interface{}) error {
	db := &Atm[interface{}]{Db: m.DB}
	return db.Create(ctx, table, data)
}

func (m *M) BatchCreate(ctx context.Context, table string, data []map[string]interface{}) error {
	db := &Atm[interface{}]{Db: m.DB}
	return db.BatchCreate(ctx, table, data)
}

func (m *M) Delete(ctx context.Context, table string, id int64) error {
	db := &Atm[interface{}]{Db: m.DB}
	return db.Delete(ctx, table, id)
}

func (m *M) DeleteBatch(ctx context.Context, table string, ids []int64) error {
	db := &Atm[interface{}]{Db: m.DB}
	return db.DeleteBatch(ctx, table, ids)
}

func (m *M) Update(ctx context.Context, table string, id int64, data clause.ColumnMap) error {
	db := &Atm[interface{}]{Db: m.DB}
	return db.Update(ctx, table, id, data)
}

func (m *M) BatchUpdates(ctx context.Context, table string, ids []int64, data clause.ColumnMap) error {
	db := &Atm[interface{}]{Db: m.DB}
	return db.BatchUpdates(ctx, table, ids, data)
}

func (m *M) First(ctx context.Context, table string, id int64) (RowsMap, error) {
	db := &Atm[RowsMap]{Db: m.DB}
	return db.First(ctx, table, id, nil)
}

func (m *M) Query(ctx context.Context, table string, clauses clause.Clauses) ([]RowsMap, error) {
	db := &Atm[RowsMap]{Db: m.DB}
	return db.Query(ctx, table, clauses)
}

func (m *M) QueryPage(ctx context.Context, table string, clauses clause.Clauses) (ResultWithPage[RowsMap], error) {
	db := &Atm[RowsMap]{Db: m.DB}
	return db.QueryPage(ctx, table, clauses)
}
