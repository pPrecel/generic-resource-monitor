package config

type Config struct {
	ServerPort    string `yaml:"serverPort"`
	ServerAddress string `yaml:"serverAddress"`

	Entries []MonitorEntry
}

type MonitorEntry struct {
	FileInfo `yaml:",inline"`
	Metrics  []MetricInfo
}

type MetricInfo struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Type        string `yaml:"type"`
	Line        int    `yaml:"line"`
	Column      int    `yaml:"column"`
}

type FileInfo struct {
	Filename  string `yaml:"filename"`
	TimeDelay string `yaml:"delay"`
}
