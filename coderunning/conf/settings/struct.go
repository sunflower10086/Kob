package settings

var Conf = new(AppConfig)

type AppConfig struct {
	AllServer *AllServer `mapstructure:"server"`
	//LogConf   *LogConfig   `mapstructure:"log"`
	//MysqlConf *MysqlConfig `mapstructure:"mysql"`
}

type AllServer struct {
	BotRunningConfig *Server `mapstructure:"botrun"`
	//ResultConfig *Server `mapstructure:"result"`
}

type Server struct {
	Name string `mapstructure:"name"`
	Mode string `mapstructure:"mode"`
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}
