package config

type App struct {
	Domain    string `mapstructure:"domain" yaml:"domain"`
	PrefixUrl string `mapstructure:"prefixUrl" yaml:"prefixUrl"`
}
