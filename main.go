package main

import (
	"fetch/receipts"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/receipts/process", processReceipt)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func processReceipt(c *gin.Context) {
	var receipt receipts.Receipt
	bindingError := c.ShouldBindJSON(&receipt)
	if bindingError != nil {
		c.AbortWithError(http.StatusBadRequest, bindingError)
		return
	}

	receiptId := receipts.Process(receipt)
	c.JSON(http.StatusOK, gin.H{
		"id": receiptId,
	})
}
