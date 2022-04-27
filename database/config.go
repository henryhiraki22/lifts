package database

import "fmt"

type Config struct {
	Host         string
	UserName     string
	Password     string
	DatabaseName string
}

var GetConnectionString = func(config Config) string {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", config.UserName, config.Password, config.Host, config.DatabaseName)
	return connectionString
}
