package config

type App struct {
	LogLevel string `mapstructure:"log_level"`

	HTTPServer struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	} `mapstructure:"http_server"`
}
