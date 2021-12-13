package models

type Basket struct {
	BookIDs []int `json:"basketBookIds" binding:"required"`
}

type Cost struct {
	TotalCost  uint        `json:"totalCost" binding:"required"`
	BookGroups []BookGroup `json:"bookGroups" binding:"required"`
}

type BookGroup struct {
	BookIDs            []int `json:"bookIds" binding:"required"`
	DiscountPercentage uint  `json:"discountPercentage" binding:"required"`
	GroupTotalCost     uint  `json:"groupTotalCost" binding:"required"`
}
