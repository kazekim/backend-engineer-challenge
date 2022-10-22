package besqlx

type Config struct {
	DatabaseDriver            string `json:"database_driver" mapstructure:"database_driver"`
	DatabaseUrl            string `json:"database_url" mapstructure:"database_url"`
	DatabasePort           string `json:"database_port" mapstructure:"database_port"`
	DatabaseName           string `json:"database_name" mapstructure:"database_name"`
	DatabaseUsername       string `json:"database_username" mapstructure:"database_username"`
	DatabasePassword       string `json:"database_password" mapstructure:"database_password"`
	DatabaseLogEnabled     bool   `json:"database_log_enabled" mapstructure:"database_log_enabled"`
	DatabaseDisableSSLMode bool   `json:"database_disable_ssl_mode" mapstructure:"database_disable_ssl_mode"`
}
