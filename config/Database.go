package config

import "fmt"

const MYSQL_URL_PATTERN = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local"

type DBConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

func BuildDBConfig(host string, port string, user string, dbName string, password string) *DBConfig {
	config := DBConfig{
		Host:     host,
		Port:     port,
		User:     user,
		DBName:   dbName,
		Password: password,
	}
	return &config
}

func (config *DBConfig) MysqlUrl() string {
	return fmt.Sprintf(
		MYSQL_URL_PATTERN,
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
	)
}
