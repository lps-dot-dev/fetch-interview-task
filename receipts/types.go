package receipts

type Receipt struct {
	Retailer     string        `json:"retailer"`
	PurchaseDate string        `json:"purchaseDate"`
	PurchaseTime string        `json:"purchaseTime"`
	Total        float32       `json:"total"`
	Items        []ReceiptItem `json:"items"`
}

type ReceiptItem struct {
	ShortDescription string  `json:"shortDescription"`
	Price            float32 `json:"price"`
}
