package config

type Config struct {
	App AppConfig `yaml:"app"`
	Database DBConfig `yaml:"database"`
}

type AppConfig struct {
	Name string `yaml:"name"`
	Port int `yaml:"port"`
}

type DBConfig struct {
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Password string `yaml:"password"`
	Name string `yaml:"name"`
	Port int `yaml:"port"`
	// MigrationsPath string `yaml:"migrations_path"`
} 