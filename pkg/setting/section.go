package setting

type Config struct {
	Server   ServerSetting   `mapstructure:"server"`
	Database DatabaseSetting `mapstructure:"databases"`
	Security SecuritySetting `mapstructure:"security"`
}
type DatabaseSetting struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DBName       string `mapstructure:"dbName"`
	SSLMode      string `mapstructure:"sslmode"`
	MaxOpenConns int    `mapstructure:"maxOpenConns"`
	MaxIdleConns int    `mapstructure:"maxIdleConns"`
	MaxLifeTime  string `mapstructure:"connectionMaxLifeTime"`
}
type ServerSetting struct {
	Port int `mapstructure:"port"`
}
type SecuritySetting struct {
	JWT JWTSetting `mapstructure:"jwt"`
}

type JWTSetting struct {
	Key string `mapstructure:"key"`
}
