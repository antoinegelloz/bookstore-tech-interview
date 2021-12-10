package main

type Settings struct {
	BookCost            uint          `json:"bookCost" binding:"required"`
	DiscountableBookIDs []int         `json:"discountableBookIds" binding:"required"`
	DiscountScaling     DiscountSlice `json:"discountScaling" binding:"required"`
}

type Discount struct {
	NumBooks           uint `json:"nbOfBooks" binding:"required"`
	DiscountPercentage uint `json:"discountPercentage" binding:"required"`
}

type DiscountSlice []Discount

func (x DiscountSlice) Len() int           { return len(x) }
func (x DiscountSlice) Less(i, j int) bool { return x[i].NumBooks > x[j].NumBooks }
func (x DiscountSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

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
