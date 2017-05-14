package model

// Config is the main Cataclysm configuration file structure
type Config struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
