package main

import (
	"matching/internal/models"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDb() {
	dsn := "root:123456@tcp(localhost:3306)/kob?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	sqlDB, _ := db.DB()
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(50)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	//if err != nil {
	//	panic(err)
	//}

}

func main() {
	initDb()

	g := gen.NewGenerator(gen.Config{
		// OutPath 生成代码的路径
		OutPath: "./internal/dao/query",
		// 生成Model的路径
		ModelPkgPath: "./internal/models",
		// 生成器模式
		// gen.WithDefaultQuery  是否生成全局变量Q作为DAO接口，如果开启，你可以通过这样的方式查询数据dal.Q.User.First()
		// gen.WithQueryInterface	生成查询API代码，而不是struct结构体。通常用来MOCK测试
		// gen.WithoutContext	生成无需传入context参数的代码
		Mode: gen.WithoutContext | gen.WithQueryInterface, // generate mode

		FieldCoverable:    true,
		FieldWithIndexTag: true,
	})

	g.UseDB(db) // reuse your gorm dal

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.GenerateModelAs("user", "User")
	//g.GenerateAllTable()

	//g.ApplyBasic(models.User{}, models.Bot{})
	g.ApplyBasic(models.User{})
	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	//g.ApplyInterface(func(Querier) {}, models.User{})

	// Generate the code
	g.Execute()
}
