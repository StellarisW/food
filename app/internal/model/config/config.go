package config

type Config struct {
	App    *App    `mapstructure:"app"  yaml:"app"`
	Logger *Logger `mapstructure:"logger" yaml:"logger"`

	DataBase *Database `mapstructure:"database"  yaml:"database"`

	Server *Server `mapstructure:"server"  yaml:"server"`

	Cors CORS `mapstructure:"cors" yaml:"cors"`

	Auth Auth `mapstructure:"auth" yaml:"auth"`

	YelpApiKey string `mapstructure:"yelpApiKey" yaml:"yelpApiKey"`
}
