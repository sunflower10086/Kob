package main

import (
	"backend/conf/mysql"
	"backend/internal/models"

	"gorm.io/gen"
)

func init() {
	mysql.Init()
}

func main() {
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

	g.UseDB(mysql.DB) // reuse your gorm dal

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	//g.GenerateModelAs(constants.UserTableName, "User")
	g.GenerateAllTable()

	g.ApplyBasic(models.User{}, models.Bot{})
	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	//g.ApplyInterface(func(Querier) {}, model.User{})

	// Generate the code
	g.Execute()
}
