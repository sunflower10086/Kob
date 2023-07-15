package settings

var Conf = new(AppConfig)

type AppConfig struct {
	AllServer *AllServer   `mapstructure:"server"`
	LogConf   *LogConfig   `mapstructure:"log"`
	MysqlConf *MysqlConfig `mapstructure:"mysql"`
}

type AllServer struct {
	MatchConfig  *Server `mapstructure:"match"`
	ResultConfig *Server `mapstructure:"result"`
}

type Server struct {
	Name string `mapstructure:"name"`
	Mode string `mapstructure:"mode"`
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type LogConfig struct {
	Level        string `mapstructure:"level"`
	InfoFilename string `mapstructure:"infoFilename"`
	ErrFilename  string `mapstructure:"errFilename"`
	MaxSize      int    `mapstructure:"max_size"`
	MaxAge       int    `mapstructure:"max_age"`
	MaxBackups   int    `mapstructure:"max_backups"`
}

type MysqlConfig struct {
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Dbname       string `mapstructure:"dbname"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
}
