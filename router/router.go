package router

import (
	"bookstore/handlers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/discountSettings", handlers.SettingsHandler)
	router.POST("/cost", handlers.CostHandler)
	return router
}
