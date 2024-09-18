package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/tiagoncardoso/golang-api/configs"
	_ "github.com/tiagoncardoso/golang-api/docs"
	"github.com/tiagoncardoso/golang-api/internal/entity"
	"github.com/tiagoncardoso/golang-api/internal/infra/database/sqlite_db"
	"github.com/tiagoncardoso/golang-api/internal/presenter/controller"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
)

// @title Golang API
// @version 1.0
// @description This is a simple API to manage products and users

// @contact.name Tiago Cardoso
// @contact.email tiago.mncardoso@gmail.com

// @host localhost:8000
// @BasePath /
// @securityDefinitions.jwt Bearer
// @in header
// @name Authorization
func main() {
	config, _ := configs.LoadConfig(".")

	db := initDb(config.DBHost)
	initWebServer(db, config.JWTTokenAuth, config.JWTTExpiresIn)
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

func initWebServer(db *gorm.DB, jwt *jwtauth.JWTAuth, jwtExpiresIn int) {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.WithValue("jwt", jwt))
	router.Use(middleware.WithValue("jwtExpiresIn", jwtExpiresIn))
	//router.Use(middleware.LogRequest) // This is my middleware

	productController := controller.NewProductController(db, router)
	productController.Register(jwt)

	userController := controller.NewUserController(db, router)
	userController.Register()

	router.Get("/api/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/api/doc.json")))

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		slog.Error("Error on server start", "msg", err)
	}
}
