package database

import "fmt"

type Config struct {
	Host         string
	UserName     string
	Password     string
	DatabaseName string
}

var GetConnectionString = func(config Config) string {
	connectionString := fmt.Sprintf(config.UserName + ":" + config.Password + "@tcp(" + config.Host + ")/" + config.DatabaseName)
	return connectionString
}
