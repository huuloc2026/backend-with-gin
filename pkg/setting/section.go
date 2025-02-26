package setting

type Config struct {
	Server   ServerSetting   `mapstructure:"server"`
	Database PostgresSetting `mapstructure:"databases"`
	Redis    RedisSetting    `mapstructure:"redis"`
	Security SecuritySetting `mapstructure:"security"`
	Logger   LoggerSetting   `mapstructure:"logger"`
}
type PostgresSetting struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
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
type RedisSetting struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Db       int    `mapstructure:"db"`
	MaxConns int    `mapstructure:"maxConns"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
type LoggerSetting struct {
	Level       string `mapstructure:"level"`
	File_log    string `mapstructure:"file"`
	Max_Size    int    `mapstructure:"maxSize"`
	Max_Backups int    `mapstructure:"maxBackups"`
	Max_Age     int    `mapstructure:"maxAge"`
	Compress    bool   `mapstructure:"compress"`
}
type SecuritySetting struct {
	JWT JWTSetting `mapstructure:"jwt"`
}

type JWTSetting struct {
	Key string `mapstructure:"key"`
}
