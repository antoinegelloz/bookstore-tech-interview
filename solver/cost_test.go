package solver

import (
	"bookstore/models"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestComputeCost(t *testing.T) {
	s := models.Settings{
		BookCost:            800,
		DiscountableBookIDs: []int{3, 47, 83, 133, 194},
		DiscountScaling: []models.Discount{
			{NumBooks: 2, DiscountPercentage: 5},
			{NumBooks: 3, DiscountPercentage: 10},
			{NumBooks: 4, DiscountPercentage: 20},
			{NumBooks: 5, DiscountPercentage: 25},
		},
	}

	b := models.Basket{
		BookIDs: []int{1, 3, 3, 47, 83, 133, 194},
	}

	expected := models.Cost{
		TotalCost: 4600,
		BookGroups: []models.BookGroup{
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

	res, err := ComputeCost(s, b)
	require.NoError(t, err)
	assert.Equal(t, res, expected)
}
