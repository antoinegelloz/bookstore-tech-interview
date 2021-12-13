package models

import (
	"fmt"
	"sort"
)

type Settings struct {
	BookCost            uint          `json:"bookCost" binding:"required"`
	DiscountableBookIDs []int         `json:"discountableBookIds" binding:"required"`
	DiscountScaling     DiscountSlice `json:"discountScaling" binding:"required"`
}

type DiscountSlice []Discount

type Discount struct {
	NumBooks           uint `json:"nbOfBooks" binding:"required"`
	DiscountPercentage uint `json:"discountPercentage" binding:"required"`
}

func (x DiscountSlice) Len() int           { return len(x) }
func (x DiscountSlice) Less(i, j int) bool { return x[i].NumBooks > x[j].NumBooks }
func (x DiscountSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func (s *Settings) Validate() error {
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

	return nil
}
