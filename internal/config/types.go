package config

type Config struct {
	ServerPort    string `yaml:"serverPort"`
	ServerAddress string `yaml:"serverAddress"`

	Entries []MonitorEntry
}

type MonitorEntry struct {
	MetricInfo `yaml:",inline"`
	FileInfo   `yaml:",inline"`
}

type MetricInfo struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Timestamp   int64  `yaml:"timestamp"`
}

type FileInfo struct {
	Filename   string `yaml:"filename"`
	LineNumber int    `yaml:"lineNumber"`
	ElemNumber int    `yaml:"elemNumber"`
}
