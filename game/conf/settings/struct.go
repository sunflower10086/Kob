package settings

import "fmt"

var Conf = new(AppConfig)

type AppConfig struct {
	AllServer *AllServer   `mapstructure:"server"`
	LogConf   *LogConfig   `mapstructure:"log"`
	MysqlConf *MysqlConfig `mapstructure:"mysql"`
	EtcdConf  *EtcdConfig  `mapstructure:"etcd"`
}

type AllServer struct {
	SnakeConfig      *Server `mapstructure:"snake"`
	ResultConfig     *Server `mapstructure:"result"`
	BotRunningConfig *Server `mapstructure:"botrun"`
}

type Server struct {
	Name string `mapstructure:"name"`
	Mode string `mapstructure:"mode"`
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

func (s *Server) GetAddr() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
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

type EtcdConfig struct {
	Endpoint string `mapstructure:"endpoint"`
}
