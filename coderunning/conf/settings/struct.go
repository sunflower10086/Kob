package settings

import "fmt"

var Conf = new(AppConfig)

type AppConfig struct {
	AllServer *AllServer  `mapstructure:"server"`
	EtcdConf  *EtcdConfig `mapstructure:"etcd"`
	//LogConf   *LogConfig   `mapstructure:"log"`
	//MysqlConf *MysqlConfig `mapstructure:"mysql"`
}

type AllServer struct {
	BotRunningConfig *Server `mapstructure:"botrun"`
	SnakeConfig      *Server `mapstructure:"snake"`
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

type EtcdConfig struct {
	Endpoint string `mapstructure:"endpoint"`
}
