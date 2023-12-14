package endpoints

import (
	"dictionary/client"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LookupDictionary(c *gin.Context) {
	symbol := c.Query("symbol")
	if symbol == "" {
		c.Status(http.StatusNotFound)
		c.Abort()
		return
	}

	translation := c.Query("translation")
	if translation == "" {
		translation = "chinese-traditional"
	}

	cambirdge := client.NewCambridge()
	lookup, _ := cambirdge.Lookup(symbol, translation)
	if lookup.Symbol == "" {
		c.Status(http.StatusNotFound)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, lookup)
}
