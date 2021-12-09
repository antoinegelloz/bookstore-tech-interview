package main

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestComputeCost(t *testing.T) {
	testSettings := Settings{
		BookCost:            800,
		DiscountableBookIDs: []int{3, 47, 83, 133, 194},
		DiscountScaling: []Discount{
			{5, 25},
			{4, 20},
			{3, 10},
			{2, 5},
		},
	}

	testBasket := Basket{
		BookIds: []int{1, 3, 3, 47, 83, 133, 194},
	}

	expectedCost := Cost{
		TotalCost: 4600,
		BookGroups: []BookGroup{
			{
				BookIDs:            []int{1, 3},
				DiscountPercentage: 0,
				GroupTotalCost:     1600,
			},
			{
				BookIDs:            []int{3, 47, 83, 133, 194},
				DiscountPercentage: 25,
				GroupTotalCost:     3000,
			},
		},
	}

	computedCost := computeCost(testSettings, testBasket)
	assert.Equal(t, computedCost, expectedCost)
}
