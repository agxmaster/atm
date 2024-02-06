package callback

import (
	"context"
	"github.com/agxmaster/atm"
	"github.com/agxmaster/atm/plugs"
	"gorm.io/gorm"
)

type HookName string

const CreateBefore HookName = "atm:CreateBefore"
const UpdateBefore HookName = "atm:UpdateBefore"
const CreateAfter HookName = "atm:CreateAfter"
const UpdateAfter HookName = "atm:UpdateAfter"
const InsertAfter HookName = "atm:InsertAfter"
const QueryAfter HookName = "atm:QueryAfter"
const DeleteAfter HookName = "atm:DeleteAfter"
const RawAfter HookName = "atm:RawAfter"
const RowAfter HookName = "atm:RowAfter"

var callbacks = map[HookName][]plugs.Plug{
	CreateBefore: {},
	UpdateBefore: {},
	CreateAfter:  {},
	UpdateAfter:  {},
	InsertAfter:  {},
	DeleteAfter:  {},
	QueryAfter:   {},
	RawAfter:     {},
	RowAfter:     {},
}

func Regist(ctx context.Context, gormDb gorm.DB, config *atm.Config) {

	if config.AutoCreatedAt || config.AutoUpdatedAt {
		callbacks[CreateBefore] = append(callbacks[CreateBefore], &plugs.CreateAt{})
	}

	if config.AutoUpdatedAt {
		callbacks[UpdateBefore] = append(callbacks[UpdateBefore], &plugs.UpdateAt{})
	}

	if config.AutoAuditLog {
		callbacks[CreateAfter] = append(callbacks[CreateAfter], &plugs.AuditLog{&plugs.AuditLogUserIdStoreInstance})
		callbacks[UpdateAfter] = append(callbacks[UpdateAfter], &plugs.AuditLog{&plugs.AuditLogUserIdStoreInstance})
		callbacks[InsertAfter] = append(callbacks[InsertAfter], &plugs.AuditLog{&plugs.AuditLogUserIdStoreInstance})
		callbacks[DeleteAfter] = append(callbacks[DeleteAfter], &plugs.AuditLog{&plugs.AuditLogUserIdStoreInstance})
		callbacks[QueryAfter] = append(callbacks[QueryAfter], &plugs.AuditLog{&plugs.AuditLogUserIdStoreInstance})
		callbacks[RawAfter] = append(callbacks[RawAfter], &plugs.AuditLog{&plugs.AuditLogUserIdStoreInstance})
		callbacks[RowAfter] = append(callbacks[RowAfter], &plugs.AuditLog{&plugs.AuditLogUserIdStoreInstance})
	}

	gormDb.Callback().Create().Before("gorm:create").Register(string(CreateBefore), func(db *gorm.DB) {
		for _, f := range callbacks[CreateBefore] {
			f.Callback(ctx, db, config)
		}
	})

	gormDb.Callback().Update().Before("gorm:update").Register(string(UpdateBefore), func(db *gorm.DB) {
		for _, f := range callbacks[UpdateBefore] {
			f.Callback(ctx, db, config)
		}
	})

	gormDb.Callback().Update().After("gorm:update").Register(string(UpdateAfter), func(db *gorm.DB) {
		for _, f := range callbacks[UpdateAfter] {
			f.Callback(ctx, db, config)
		}
	})

	gormDb.Callback().Create().After("gorm:create").Register(string(CreateAfter), func(db *gorm.DB) {
		for _, f := range callbacks[CreateAfter] {
			f.Callback(ctx, db, config)
		}
	})

	gormDb.Callback().Query().After("gorm:query").Register(string(QueryAfter), func(db *gorm.DB) {
		for _, f := range callbacks[QueryAfter] {
			f.Callback(ctx, db, config)
		}
	})

	gormDb.Callback().Delete().After("gorm:delete").Register(string(DeleteAfter), func(db *gorm.DB) {
		for _, f := range callbacks[DeleteAfter] {
			f.Callback(ctx, db, config)
		}
	})

	gormDb.Callback().Raw().After("gorm:raw").Register(string(RawAfter), func(db *gorm.DB) {
		for _, f := range callbacks[RawAfter] {
			f.Callback(ctx, db, config)
		}
	})

	gormDb.Callback().Row().After("gorm:row").Register(string(RowAfter), func(db *gorm.DB) {
		for _, f := range callbacks[RowAfter] {
			f.Callback(ctx, db, config)
		}
	})
}
