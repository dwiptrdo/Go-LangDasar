package database

var connection string

func init() {
	connection = "MongoDB"
}

func GetDatabase() string {
	return connection
}