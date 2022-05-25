package app

import "os"

func GetDatabaseName() string {
	return os.Getenv("DB_NAME")
}
