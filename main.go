package main

import (
	"fetch/receipts"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/receipts/process", func(c *gin.Context) {
		var receiptId = receipts.Process()
		c.JSON(http.StatusOK, gin.H{
			"id": receiptId,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
