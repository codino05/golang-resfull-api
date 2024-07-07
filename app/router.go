package app

import (
	"go-restfull-api/controller"
	"go-restfull-api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryContoller controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryContoller.FindAll)
	router.GET("/api/categories/:categoryId", categoryContoller.FindById)
	router.POST("/api/categories", categoryContoller.Create)
	router.PUT("/api/categories/:categoryId", categoryContoller.Update)
	router.DELETE("/api/categories/:categoryId", categoryContoller.Delete)

	router.PanicHandler = exception.ErrorHandler
	return router
}
