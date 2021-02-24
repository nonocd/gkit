package config

import "github.com/spf13/viper"

// Configer efines how to get and set value from configuration.
type Configer interface {
	Scan(v interface{}) error
	Get(key string) interface{}
}

// Config is a yaml config parser and implements Configer interface.
type Config struct {
	vp *viper.Viper
}

func NewConfig(filename string) Config {
	vp := viper.New()
	vp.SetConfigFile(filename)
	return NewConfigWithViper(vp)
}

func NewConfigWithViper(vp *viper.Viper) Config {
	return Config{
		vp: vp,
	}
}

func (c *Config) Load() error {
	return c.vp.ReadInConfig()
}

func (c *Config) Scan(v interface{}) error {
	return c.vp.Unmarshal(v)
}

func (c *Config) Get(key string) interface{} {
	return c.vp.Get(key)
}
