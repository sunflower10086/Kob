package mysql

import (
	"fmt"
	"snake/conf/settings"
	"snake/internal/dao/query"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Q *query.Query
)

func Init(conf *settings.AppConfig) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.MysqlConf.User,
		conf.MysqlConf.Password,
		conf.MysqlConf.Host,
		conf.MysqlConf.Port,
		conf.MysqlConf.Dbname,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, _ := db.DB()
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(conf.MysqlConf.MaxIdleConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(conf.MysqlConf.MaxOpenConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil {
		return err
	}

	Q = query.Use(db)
	return nil

}
