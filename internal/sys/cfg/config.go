package cfg

import "flag"

type (
	Config struct {
		Log    *LogConfig    `json:"log"`
		Server *ServerConfig `json:"server"`
		DB     *DBConfig     `json:"db"`
		Prop   *PropConfig   `json:"prop"`
	}

	LogConfig struct {
		Level string `json:"server"`
	}

	ServerConfig struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	}

	DBConfig struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
		Schema   string `json:"schema"`
		Name     string `json:"name"`
		SSL      bool   `json:"ssl"`
	}

	PropConfig struct {
		PageSize int `json:"page-size"`
	}
)

func Load() *Config {
	config := &Config{
		Log:    &LogConfig{},
		Server: &ServerConfig{},
		DB:     &DBConfig{},
		Prop:   &PropConfig{},
	}

	flag.StringVar(&config.Log.Level, "log-level", "info", "Log level")
	flag.StringVar(&config.Server.Host, "server-host", "localhost", "Server host")
	flag.IntVar(&config.Server.Port, "server-port", 8080, "Server port")
	flag.StringVar(&config.DB.Host, "db-host", "localhost", "Database host")
	flag.IntVar(&config.DB.Port, "db-port", 5432, "Database port")
	flag.StringVar(&config.DB.Username, "db-username", "admin", "Database username")
	flag.StringVar(&config.DB.Password, "db-password", "password", "Database password")
	flag.StringVar(&config.DB.Schema, "db-schema", "ak", "Database schema")
	flag.StringVar(&config.DB.Name, "db-name", "ak", "Database name")
	flag.BoolVar(&config.DB.SSL, "db-ssl", true, "Database use SSL")
	flag.IntVar(&config.Prop.PageSize, "prop-page-size", 12, "Pagination page size")

	flag.Parse()

	return config
}
