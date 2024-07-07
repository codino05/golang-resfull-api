package main

import (
	"go-restfull-api/app"
	"go-restfull-api/controller"
	"go-restfull-api/helper"
	"go-restfull-api/middleware"
	"go-restfull-api/repository"
	"go-restfull-api/service"
	"net/http"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryContoller := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryContoller)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicHandler(err)
}
