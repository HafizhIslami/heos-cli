package config

type Config struct {
	Ollama   OllamaConfig   `yaml:"ollama"`
	Workers  WorkerConfig   `yaml:"workers"`
	Workspace WorkspaceConfig `yaml:"workspace"`
	HEOS     HEOSConfig     `yaml:"heos"`
}

type OllamaConfig struct {
	Host    string `yaml:"host"`
	Timeout int    `yaml:"timeout"`
}

type WorkerConfig struct {
	Architect string `yaml:"architect"`
	Engineer  string `yaml:"engineer"`
}

type WorkspaceConfig struct {
	TaskDir string `yaml:"task_dir"`
}

type HEOSConfig struct {
	Directory string `yaml:"directory"`
}