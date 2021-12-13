package solver

import (
	"bookstore/models"
	"fmt"
	"sort"
)

var discountableBookStacks map[int]int

func ComputeCost(s models.Settings, basket models.Basket) (res models.Cost, err error) {
	if err = s.Validate(); err != nil {
		return res, fmt.Errorf("invalid settings: %w", err)
	}

	s.DiscountScaling = append(s.DiscountScaling, models.Discount{
		NumBooks:           1,
		DiscountPercentage: 0,
	})
	sort.Ints(basket.BookIDs)

	// Split discountable and not discountable books in two slices
	var discountableBookIDs, notDiscountableBookIDs []int
	var notDiscountableGroupTotalCost uint
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
			notDiscountableGroupTotalCost += s.BookCost
		}
	}

	// Create the first group with the not discountable books
	if len(notDiscountableBookIDs) > 0 {
		res.BookGroups = append(res.BookGroups, models.BookGroup{
			BookIDs:            notDiscountableBookIDs,
			DiscountPercentage: 0,
			GroupTotalCost:     notDiscountableGroupTotalCost,
		})
		res.TotalCost = notDiscountableGroupTotalCost
	}

	fmt.Printf("discountableBooks: %+v\n", discountableBookIDs)
	fmt.Printf("notDiscountableBooks: %+v\n", notDiscountableBookIDs)

	if len(discountableBookIDs) == 0 {
		return
	}

	// Store the discountable books as one stack per book ID
	discountableBookStacks = map[int]int{
		discountableBookIDs[0]: 1,
	}
	for i, bookID := range discountableBookIDs {
		if i == 0 {
			continue
		}
		if _, ok := discountableBookStacks[bookID]; !ok {
			discountableBookStacks[bookID] = 1
		} else {
			discountableBookStacks[bookID] += 1
		}
	}
	fmt.Printf("discountableBookStacks: %+v\n", discountableBookStacks)

	// Create the book groups by using the stacks with a naive algorithm, starting with the largest discounts.
	// The books are removed from the stack when a group is created.
	for len(discountableBookStacks) > 0 {
		for _, discountScale := range s.DiscountScaling {
			if int(discountScale.NumBooks) <= len(discountableBookStacks) {
				newBookGroup := createBookGroup(s, discountScale)
				sort.Ints(newBookGroup.BookIDs)

				// If the new group doesn't have a discount, merge it to the not discountable books group
				if newBookGroup.DiscountPercentage == 0 && res.BookGroups[0].DiscountPercentage == 0 {
					res.BookGroups[0] = mergeBookGroups(res.BookGroups[0], newBookGroup)
					sort.Ints(res.BookGroups[0].BookIDs)
				} else {
					res.BookGroups = append(res.BookGroups, newBookGroup)
				}
				res.TotalCost += newBookGroup.GroupTotalCost
				fmt.Printf("discountableBookStacks: %+v\n", discountableBookStacks)
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
	for stackBookID := range discountableBookStacks {
		bookIDs = append(bookIDs, stackBookID)
		discountableBookStacks[stackBookID]--
		if discountableBookStacks[stackBookID] == 0 {
			delete(discountableBookStacks, stackBookID)
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
