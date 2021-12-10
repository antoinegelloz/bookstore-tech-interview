package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
)

var currentSettings Settings

func setSettings(c *gin.Context) {
	var settings Settings
	if err := c.ShouldBindJSON(&settings); err != nil {
		fmt.Printf("err: %s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validateAndSetSettings(settings); err != nil {
		fmt.Printf("err: %s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, currentSettings)
}

func validateAndSetSettings(s Settings) error {
	sort.Sort(s.DiscountScaling)
	for i, d := range s.DiscountScaling {
		if d.NumBooks < 2 {
			return fmt.Errorf("invalid settings: nbOfBooks has to be greater than 2: %+v\n", s)
		}
		if i > 0 {
			if s.DiscountScaling[i].DiscountPercentage >= s.DiscountScaling[i-1].DiscountPercentage {
				return fmt.Errorf("invalid settings: discount has to increase with nbOfBooks: %+v\n", s)
			}
			if s.DiscountScaling[i].NumBooks == s.DiscountScaling[i-1].NumBooks {
				return fmt.Errorf("invalid settings: all nbOfBooks need to be different: %+v\n", s)
			}
		}
	}

	currentSettings = s
	return nil
}

func cost(c *gin.Context) {
	var basket Basket
	if err := c.ShouldBindJSON(&basket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("Basket: %+v\n", basket)
	fmt.Printf("CurrentSettings: %+v\n", currentSettings)

	c.JSON(http.StatusOK, computeCost(basket))
}
