package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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

type Settings struct {
	BookCost            int        `json:"bookCost" binding:"required"`
	DiscountableBookIDs []int      `json:"discountableBookIds" binding:"required"`
	DiscountScaling     []Discount `json:"discountScaling" binding:"required"`
}

type Discount struct {
	NumBooks           int `json:"nbOfBooks" binding:"required"`
	DiscountPercentage int `json:"discountPercentage" binding:"required"`
}

var currentSettings Settings

func setSettings(c *gin.Context) {
	var settings Settings
	if err := c.ShouldBindJSON(&settings); err != nil {
		fmt.Printf("err: %s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currentSettings = settings

	c.JSON(http.StatusOK, currentSettings)
}

type Basket struct {
	BookIds []int `json:"basketBookIds" binding:"required"`
}

type Cost struct {
	TotalCost  int         `json:"totalCost" binding:"required"`
	BookGroups []BookGroup `json:"bookGroups" binding:"required"`
}

type BookGroup struct {
	BookIDs            []int `json:"bookIds" binding:"required"`
	DiscountPercentage int   `json:"discountPercentage" binding:"required"`
	GroupTotalCost     int   `json:"groupTotalCost" binding:"required"`
}


func cost(c *gin.Context) {
	var basket Basket
	if err := c.ShouldBindJSON(&basket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("Basket: %+v\n", basket)
	fmt.Printf("CurrSettings: %+v\n", currentSettings)

	c.JSON(http.StatusOK, computeCost(currentSettings, basket))
}
