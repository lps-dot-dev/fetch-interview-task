package main

import (
	"bytes"
	"encoding/gob"
	"fetch/receipts"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	receiptScore, processingError := receipts.Process(receipt)
	if processingError != nil {
		c.AbortWithError(http.StatusBadRequest, processingError)
		return
	}

	// We should account for duplicate receipts
	// If the same receipt is given twice, we should return the same UUID
	var receiptBuffer bytes.Buffer
	encoder := gob.NewEncoder(&receiptBuffer)
	encodingError := encoder.Encode(receipt)
	if encodingError != nil {
		c.AbortWithError(http.StatusInternalServerError, encodingError)
		return
	}

	uuid := uuid.NewSHA1(uuid.NameSpaceURL, receiptBuffer.Bytes())

	c.JSON(http.StatusOK, gin.H{
		"id":    uuid.String(),
		"score": receiptScore,
	})
}
