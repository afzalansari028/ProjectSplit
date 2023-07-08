package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestAPI(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
func DoSplit(c *gin.Context) {
	fmt.Println("Split function called!!!")
}
