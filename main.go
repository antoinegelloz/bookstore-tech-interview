package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/discountSettings", setSettings)
	router.POST("/cost", cost)
	if err := router.Run(); err != nil {
		fmt.Printf("api error: %s\n", err)
		return
	}
}
