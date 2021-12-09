package main

import (
	"fmt"
	"sort"
)

var bookStacks map[int]int

//TODO: sort desc settings.DiscountScaling or change data structure
func computeCost(settings Settings, basket Basket) Cost {
	settings.DiscountScaling = append(settings.DiscountScaling, Discount{
		NumBooks:           1,
		DiscountPercentage: 0,
	})

	var discountableBookIDs []int
	var notDiscountableBookIDs []int
	var notDiscountableGroupCost int
	sort.Ints(basket.BookIds)
	for _, bookID := range basket.BookIds {
		isBookDiscountable := false
		for _, discountableBookID := range settings.DiscountableBookIDs {
			if bookID == discountableBookID {
				isBookDiscountable = true
				break
			}
		}
		if isBookDiscountable {
			discountableBookIDs = append(discountableBookIDs, bookID)
		} else {
			notDiscountableBookIDs = append(notDiscountableBookIDs, bookID)
			notDiscountableGroupCost += settings.BookCost
		}
	}

	var res Cost
	res.BookGroups = append(res.BookGroups, BookGroup{
		BookIDs:            notDiscountableBookIDs,
		DiscountPercentage: 0,
		GroupTotalCost:     notDiscountableGroupCost,
	})
	res.TotalCost = notDiscountableGroupCost

	fmt.Printf("discountableBooks: %+v\n", discountableBookIDs)
	fmt.Printf("notDiscountableBooks: %+v\n", notDiscountableBookIDs)

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
		for _, discountScale := range settings.DiscountScaling {
			if discountScale.NumBooks <= len(bookStacks) {
				newBookGroup := createBookGroup(settings, discountScale)
				sort.Ints(newBookGroup.BookIDs)
				if newBookGroup.DiscountPercentage == 0 && res.BookGroups[0].DiscountPercentage == 0 {
					res.BookGroups[0] = mergeGroups(res.BookGroups[0], newBookGroup)
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

	return res
}

func mergeGroups(bookGroup1, bookGroup2 BookGroup) BookGroup {
	var newBookGroup BookGroup
	newBookGroup.BookIDs = append(newBookGroup.BookIDs, bookGroup1.BookIDs...)
	newBookGroup.BookIDs = append(newBookGroup.BookIDs, bookGroup2.BookIDs...)

	newBookGroup.DiscountPercentage = bookGroup1.DiscountPercentage
	newBookGroup.GroupTotalCost = bookGroup1.GroupTotalCost + bookGroup2.GroupTotalCost

	return newBookGroup
}

func createBookGroup(settings Settings, discountScale Discount) BookGroup {
	var bookIDs []int
	var groupTotalCost int
	numBooks := discountScale.NumBooks
	for stackBookID := range bookStacks {
		bookIDs = append(bookIDs, stackBookID)
		bookStacks[stackBookID]--
		if bookStacks[stackBookID] == 0 {
			delete(bookStacks, stackBookID)
		}
		groupTotalCost += int(float64(settings.BookCost) * (1. - float64(discountScale.DiscountPercentage)/100))
		numBooks--
		if numBooks == 0 {
			break
		}
	}

	return BookGroup{
		BookIDs:            bookIDs,
		DiscountPercentage: discountScale.DiscountPercentage,
		GroupTotalCost:     groupTotalCost,
	}
}
