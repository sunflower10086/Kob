package settings

import "fmt"

type AppConfig struct {
	AllServer *AllServer `mapstructure:"server"`

	LogConf   *LogConfig   `mapstructure:"log"`
	MysqlConf *MysqlConfig `mapstructure:"mysql"`
	RedisConf *RedisConfig `mapstructure:"redis"`
}

type AllServer struct {
	HttpConfig   *Server `mapstructure:"http"`
	MatchConfig  *Server `mapstructure:"match"`
	SnakeConfig  *Server `mapstructure:"gameSnake"`
	ResultConfig *Server `mapstructure:"result"`
	BotRunConfig *Server `mapstructure:"botrun"`
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
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Dbname       string `mapstructure:"dbname"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
}
