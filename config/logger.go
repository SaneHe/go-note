package config

type Logger struct {
	Debug              string `yaml:"debug"`
	ConsoleColor       bool   `yaml:"consoleColor"`
	LogFileName        string `yaml:"logFileName"`
	ReportCaller       bool   `yaml:"reportCaller"`
	LogTimestampFormat string `yaml:"logTimestampFormat"`
}
