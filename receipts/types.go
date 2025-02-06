package receipts

type Receipt struct {
	Retailer     string        `json:"retailer" binding:"required"`
	PurchaseDate string        `json:"purchaseDate" binding:"required"`
	PurchaseTime string        `json:"purchaseTime" binding:"required"`
	Total        string        `json:"total" binding:"numeric"`
	Items        []ReceiptItem `json:"items" binding:"min=1"`
}

type ReceiptItem struct {
	ShortDescription string `json:"shortDescription" binding:"required"`
	Price            string `json:"price" binding:"numeric"`
}

type ReceiptScoreRouteParams struct {
	Uuid string `uri:"uuid" binding:"required,uuid"`
}
