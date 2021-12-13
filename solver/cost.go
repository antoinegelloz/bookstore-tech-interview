package solver

import (
	"bookstore/models"
	"fmt"
	"sort"
)

var bookStacks map[int]int

func ComputeCost(s models.Settings, basket models.Basket) (res models.Cost, err error) {
	if err = s.Validate(); err != nil {
		return res, fmt.Errorf("invalid settings: %w", err)
	}

	s.DiscountScaling = append(s.DiscountScaling, models.Discount{
		NumBooks:           1,
		DiscountPercentage: 0,
	})

	// Process not discountable books
	var discountableBookIDs []int
	var notDiscountableBookIDs []int
	var notDiscountableGroupCost uint
	sort.Ints(basket.BookIDs)
	for _, bookID := range basket.BookIDs {
		isBookDiscountable := false
		for _, discountableBookID := range s.DiscountableBookIDs {
			if bookID == discountableBookID {
				isBookDiscountable = true
				break
			}
		}
		if isBookDiscountable {
			discountableBookIDs = append(discountableBookIDs, bookID)
		} else {
			notDiscountableBookIDs = append(notDiscountableBookIDs, bookID)
			notDiscountableGroupCost += s.BookCost
		}
	}

	res.BookGroups = append(res.BookGroups, models.BookGroup{
		BookIDs:            notDiscountableBookIDs,
		DiscountPercentage: 0,
		GroupTotalCost:     notDiscountableGroupCost,
	})
	res.TotalCost = notDiscountableGroupCost

	fmt.Printf("discountableBooks: %+v\n", discountableBookIDs)
	fmt.Printf("notDiscountableBooks: %+v\n", notDiscountableBookIDs)

	// Process discountable books
	bookStacks = map[int]int{
		discountableBookIDs[0]: 1,
	}
	for i, bookID := range discountableBookIDs {
		if i == 0 {
			continue
		}
		if _, ok := bookStacks[bookID]; !ok {
			bookStacks[bookID] = 1
		} else {
			bookStacks[bookID] += 1
		}
	}
	fmt.Printf("bookStacks: %+v\n", bookStacks)

	for len(bookStacks) > 0 {
		for _, discountScale := range s.DiscountScaling {
			if int(discountScale.NumBooks) <= len(bookStacks) {
				newBookGroup := createBookGroup(s, discountScale)
				sort.Ints(newBookGroup.BookIDs)
				if newBookGroup.DiscountPercentage == 0 && res.BookGroups[0].DiscountPercentage == 0 {
					res.BookGroups[0] = mergeBookGroups(res.BookGroups[0], newBookGroup)
					sort.Ints(res.BookGroups[0].BookIDs)
				} else {
					res.BookGroups = append(res.BookGroups, newBookGroup)
				}
				res.TotalCost += newBookGroup.GroupTotalCost
				fmt.Printf("bookStacks: %+v\n", bookStacks)
			}
		}
		fmt.Printf("bookGroups: %+v\n", res.BookGroups)
	}

	return res, nil
}

func createBookGroup(s models.Settings, discountScale models.Discount) models.BookGroup {
	var bookIDs []int
	var groupTotalCost float64
	numBooks := discountScale.NumBooks
	for stackBookID := range bookStacks {
		bookIDs = append(bookIDs, stackBookID)
		bookStacks[stackBookID]--
		if bookStacks[stackBookID] == 0 {
			delete(bookStacks, stackBookID)
		}
		groupTotalCost += float64(s.BookCost) *
			(1. - float64(discountScale.DiscountPercentage)/100)
		numBooks--
		if numBooks == 0 {
			break
		}
	}

	return models.BookGroup{
		BookIDs:            bookIDs,
		DiscountPercentage: discountScale.DiscountPercentage,
		GroupTotalCost:     uint(groupTotalCost),
	}
}

func mergeBookGroups(bookGroup1, bookGroup2 models.BookGroup) models.BookGroup {
	var newBookGroup models.BookGroup
	newBookGroup.BookIDs = append(newBookGroup.BookIDs, bookGroup1.BookIDs...)
	newBookGroup.BookIDs = append(newBookGroup.BookIDs, bookGroup2.BookIDs...)

	newBookGroup.DiscountPercentage = bookGroup1.DiscountPercentage
	newBookGroup.GroupTotalCost = bookGroup1.GroupTotalCost + bookGroup2.GroupTotalCost

	return newBookGroup
}
