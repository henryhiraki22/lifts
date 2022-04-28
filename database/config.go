package database

import "fmt"

type Config struct {
	Host         string
	UserName     string
	Password     string
	DatabaseName string
}

var GetConnectionString = func(config Config) string {
	//connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", config.UserName, config.Password, config.Host, config.DatabaseName)
	connectionString := fmt.Sprintf(config.UserName + ":" + config.Password + "@tcp(" + config.Host + ")/" + config.DatabaseName)
	// "lift:lift1234@tcp(db:3306)/lifts?charset=utf8mb4&parseTime=True&loc=Local")
	return connectionString
}
