package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/Flash-Pass/flash-pass-server/config"
)

func InitMySQL(cfg config.MySQLConfig) (*gorm.DB, error) {
	return gorm.Open(
		mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)),
		&gorm.Config{
			PrepareStmt:                              true,
			SkipDefaultTransaction:                   true,
			DisableForeignKeyConstraintWhenMigrating: true,
		},
	)
}
