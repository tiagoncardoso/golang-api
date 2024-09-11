package sqlite_db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ConnectDb struct {
	Dsn string
}

func NewConnect(dsn string) *ConnectDb {
	return &ConnectDb{
		Dsn: dsn,
	}
}

func (c *ConnectDb) Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(c.Dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
