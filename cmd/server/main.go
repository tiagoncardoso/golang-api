package main

import (
	"github.com/tiagoncardoso/golang-api/configs"
	"github.com/tiagoncardoso/golang-api/internal/entity"
	"github.com/tiagoncardoso/golang-api/internal/infra/database/sqlite_db"
	"github.com/tiagoncardoso/golang-api/internal/presenter/controller"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
)

func main() {
	config, _ := configs.LoadConfig(".")

	db := initDb(config.DBHost)
	initWebServer(db)
}

func initDb(dsn string) *gorm.DB {
	con := sqlite_db.NewConnect(dsn)
	db, err := con.Connect()
	if err != nil {
		panic(err)
	}

	_ = db.AutoMigrate(&entity.Product{}, &entity.User{})

	return db
}

func initWebServer(db *gorm.DB) {
	controller.NewProductController(db)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		slog.Error("Error on server start", "msg", err)
	}
}
