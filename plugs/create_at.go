package plugs

import (
	"context"
	"github.com/agxmaster/atm"
	"github.com/agxmaster/atm/util"
	"gorm.io/gorm"
	"time"
)

var Now = time.Now()

type HookName string

type Plug interface {
	Callback(ctx context.Context, db *gorm.DB, config *atm.Config)
}

type PlugBase struct {
	Name HookName
}
type CreateAt struct {
}

func (p *CreateAt) Callback(ctx context.Context, db *gorm.DB, config *atm.Config) {
	switch data := db.Statement.Dest.(type) {
	case *atm.RowsMap:
		if config.AutoCreatedAt && util.Contains(db.Statement.Table, config.AutoCreatedAtIgnoreTables) {
			if _, ok := (*data)["created_at"]; !ok {
				(*data)["created_at"] = time.Now()
			}

		}
		if config.AutoUpdatedAt && util.Contains(db.Statement.Table, config.AutoUpdatedAtIgnoreTables) {
			if _, ok := (*data)["updated_at"]; !ok {
				(*data)["updated_at"] = time.Now()
			}

		}

	case *[]atm.RowsMap:
		for k := range *data {
			if config.AutoCreatedAt && util.Contains(db.Statement.Table, config.AutoCreatedAtIgnoreTables) {
				if _, ok := (*data)[k]["created_at"]; !ok {
					(*data)[k]["created_at"] = Now
				}
			}
			if config.AutoUpdatedAt && util.Contains(db.Statement.Table, config.AutoUpdatedAtIgnoreTables) {
				if _, ok := (*data)[k]["updated_at"]; !ok {
					(*data)[k]["updated_at"] = Now
				}
			}

		}
	}
}

type UpdateAt struct {
}

func (p *UpdateAt) Callback(ctx context.Context, db *gorm.DB, config *atm.Config) {

	if !util.Contains(db.Statement.Table, config.AutoUpdatedAtIgnoreTables) {
		return
	}

	switch data := db.Statement.Dest.(type) {
	case *atm.RowsMap:
		if _, ok := (*data)["updated_at"]; !ok {
			(*data)["updated_at"] = Now
		}

	case atm.RowsMap:
		if _, ok := data["updated_at"]; !ok {
			data["updated_at"] = Now
		}
	}
}
