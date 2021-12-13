package handlers

import (
	"bookstore/global"
	"bookstore/models"
	"bookstore/solver"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SettingsHandler(c *gin.Context) {
	var s models.Settings
	if err := c.ShouldBindJSON(&s); err != nil {
		fmt.Printf("err: %s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := s.Validate(); err != nil {
		fmt.Printf("err: %s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	global.Settings = s

	c.JSON(http.StatusOK, global.Settings)
}

func CostHandler(c *gin.Context) {
	var (
		b   models.Basket
		res models.Cost
		err error
	)

	if err = c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("Basket: %+v\n", b)
	fmt.Printf("Current Settings: %+v\n", global.Settings)

	if res, err = solver.ComputeCost(global.Settings, b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
