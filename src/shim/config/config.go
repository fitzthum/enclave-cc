package config

type Containerd struct {
	AgentContainerInstance string `toml:"agent_container_instance"`
	AgentContainerRootDir  string `toml:"agent_container_root_dir"`
}

type Config struct {
	RuntimeClass string     `toml:"runtime_class"`
	LogLevel     string     `toml:"log_level"`
	Containerd   Containerd `toml:"containerd"`
}