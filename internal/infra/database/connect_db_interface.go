package database

type DbConnectInterface interface {
	Connect(dsn string) (interface{}, error)
}
