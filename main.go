package main

import (
	"bytes"
	"encoding/gob"
	"fetch/cache"
	"fetch/receipts"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Make a singleton cache with a lifespan of the runtime of the app
var receiptScoreCache *cache.Cache[string, int] = cache.New[string, int]()

func main() {
	r := gin.Default()
	r.POST("/receipts/process", processReceipt)
	r.GET("/receipts/:uuid/points", getReceiptPoints)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func getReceiptPoints(c *gin.Context) {
	var routeParams receipts.ReceiptScoreRouteParams
	bindingError := c.ShouldBindUri(&routeParams)
	if bindingError != nil {
		c.AbortWithError(http.StatusBadRequest, bindingError)
		return
	}

	receiptScore, valueFound := receiptScoreCache.Get(routeParams.Uuid)
	if valueFound == false {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"points": receiptScore,
	})
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

	// We should account for duplicate receipts, so we will encode our receipt into bytes
	var receiptBuffer bytes.Buffer
	encoder := gob.NewEncoder(&receiptBuffer)
	encodingError := encoder.Encode(receipt)
	if encodingError != nil {
		c.AbortWithError(http.StatusInternalServerError, encodingError)
		return
	}

	// And use those bytes above to generate a UUID
	uuid := uuid.NewSHA1(uuid.NameSpaceURL, receiptBuffer.Bytes())
	receiptScoreCache.Set(uuid.String(), receiptScore)

	c.IndentedJSON(http.StatusOK, gin.H{
		"id":     uuid.String(),
		"points": receiptScore,
	})
}
