package config

type RedisConfig struct {
	Host     string `yaml:"host"`
	Password string `yaml:"passwd"`
	DB       int    `yaml:"database"`
}
