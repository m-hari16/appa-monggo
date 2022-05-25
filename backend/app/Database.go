package app

type Database interface {
	NewDB() interface{}
}

func NewDatabase() Database {
	return MongoDB{}
}
