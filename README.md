# Book Store

## Introduction

To try and encourage more sales of different books from a popular 5 book series, an online bookshop has decided to offer
discounts on multiple book purchases.

One copy of any of the five books costs $8.

If, however, you buy two different books, you get a 5% discount on those two books.

If you buy 3 different books, you get a 10% discount.

If you buy 4 different books, you get a 20% discount.

If you buy all 5, you get a 25% discount.

Note: that if you buy four books, of which 3 are different titles, you get a 10% discount on the 3 that form part of a
set, but the fourth book still costs $8.

Your mission is to write a system to calculate the price of any conceivable shopping basket, giving as big a discount as
possible.

For example, how much does this basket of books costs?

* 2 copies of the 1st book
* 2 copies of the 2nd book
* 2 copies of the 3rd book
* 1 copy of the 4th book
* 1 copy of the 5th book

One way of grouping these 8 books is:

* 1 group of 5 -> 25% discount (1st, 2nd, 3rd, 4th, 5th)
* 1 group of 3 -> 10% discount (1st, 2nd, 3rd)

This would give a total of:

* 5 books at a 25% discount
* 3 books at a 10% discount

Resulting in:

* 5 * (8 - 2.00) = 5 * 6.00 = $30.00
* 3 * (8 - 0.80) = 3 * 7.20 = $21.60

For a total of $51.60

However, a different way to group these 8 books is:

* 1 group of 4 books -> 20% discount (1st, 2nd, 3rd, 4th)
* 1 group of 4 books -> 20% discount (1st, 2nd, 3rd, 5th)

This would give a total of:

* 4 books at a 20% discount
* 4 books at a 20% discount

Resulting in:

* 4 * (8 - 1.60) = 4 * 6.40 = $25.60
* 4 * (8 - 1.60) = 4 * 6.40 = $25.60

For a total of $51.20

And $51.20 is the price with the biggest discount.

# Instructions

Design and implement a REST API using whatever libraries you like in Go. It must contain two endpoints:

## POST /discountSettings

It receives the settings for today's discounts and would be called by the shop owners to control their shop sales.

Each call to this endpoint will replace any previous one.

Example of received body (based on previous examples):

	{
		"bookCost": 800,
		"discountableBookIds": [3, 47, 83, 133, 194],
		"discountScaling": [
			{ "nbOfBooks": 2, "discountPercentage": 5},
			{ "nbOfBooks": 3, "discountPercentage": 10},
			{ "nbOfBooks": 4, "discountPercentage": 20},
			{ "nbOfBooks": 5, "discountPercentage": 25}
		] 
	}

## POST /cost

It receives a basket and would be called during a book sale, obviously way more often than the settings endpoint.

Example of received body:

	{
		"basketBookIds": [1, 3, 3, 47, 83, 133, 194]
	}

Example of resulting body:

	{
		"bookGroups": [
			{
				"bookIds": [1, 3],
				"discountPercentage": 0,
				"groupTotalCost": 1600
			},
			{
				"bookIds": [3, 47, 83, 133, 194],
				"discountPercentage": 25,
				"groupTotalCost": 3000
			},
		],
		"totalCost": 4600
	}

This endpoint will return the total cost (after discounts) in cents and the different book groups. For example, for a
single book, the cost is 800 cents, which equals $8.00. Only integer calculations are necessary for this exercise.

Greatly inspired by https://exercism.org/tracks/go/exercises/book-store
