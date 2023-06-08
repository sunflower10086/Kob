package constants

var (
	BackendAddr = ":3000"
)

var (
	EtcdAddress     = "127.0.0.1:2379"
	RedisAddress    = "127.0.0.1:6379"
	MySQLDefaultDSN = "root:123456@tcp(localhost:3306)/kob?charset=utf8mb4&parseTime=True&loc=Local"
)

var (
	UserTableName = "user"
)

var (
	InfoPath  = "../log/info.log"
	ErrorPath = "../log/error.log"
)
