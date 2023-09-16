package mysql

import (
	"fmt"
	"github.com/zakariawahyu/go-api-product/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDBConnection(cfg *config.Config) (*gorm.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.Mysql.User,
		cfg.Mysql.Password,
		cfg.Mysql.Host,
		cfg.Mysql.Port,
		cfg.Mysql.DbName,
	)

	db, err := gorm.Open(mysql.Open(dataSourceName))
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(cfg.Mysql.MaxIdleConnection)
	sqlDB.SetMaxOpenConns(cfg.Mysql.MaxOpenConnection)

	return db, nil
}
