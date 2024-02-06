package plugs

import (
	"context"
	"fmt"
	"github.com/agxmaster/atm"
	"github.com/agxmaster/atm/util"
	"gorm.io/gorm"
)

const UserIdKey = "userIdKey"

type UserIdStore interface {
	GetUserId(ctx context.Context, db *gorm.DB) int64
	SetUserId(ctx context.Context, db *gorm.DB, userId int64) *gorm.DB
}

var AuditLogUserIdStoreInstance = AuditLogUserIdStore{}

type AuditLogUserIdStore struct {
}

type AuditLog struct {
	UserIdStore
}

func (s *AuditLogUserIdStore) GetUserId(ctx context.Context, db *gorm.DB) int64 {
	if userIdStr, ok := db.Get(UserIdKey); ok {
		if userId, ok := userIdStr.(int64); ok {
			return userId
		}
	}
	return -1
}

func (s *AuditLogUserIdStore) SetUserId(ctx context.Context, db *gorm.DB, userId int64) *gorm.DB {
	return db.Set(UserIdKey, userId)
}

func (p *AuditLog) Callback(ctx context.Context, db *gorm.DB, config *atm.Config) {

	if !util.Contains(db.Statement.Table, config.AutoAuditLogTables) {
		return
	}

	sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return db
	})
	goDb, err := db.DB()
	if err != nil {
		db.Logger.Error(ctx, "insert audit log  get db error: %+v", err)
		return
	}
	_, err = goDb.Exec(fmt.Sprintf("insert into auditlog(table_name,sql_text,user_id)  values(\"%s\", \"%s\", %d)", db.Statement.Table, sql, p.GetUserId(ctx, db)))
	if err != nil {
		db.Logger.Error(ctx, "insert audit log error: %+v ", err)
	}
}

/**
CREATE TABLE `auditlog` (
  `id` int NOT NULL AUTO_INCREMENT,
  `table_name` varchar(255) NOT NULL DEFAULT '',
  `sql_text` text,
  `user_id` int NOT NULL DEFAULT '0',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
*/
