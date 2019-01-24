package api

import (
	bdb "bytom-bdb/bigchaindb"
	"net/http"

	"github.com/gin-gonic/gin"
)

//create transaction
func CreateTransaction(c *gin.Context) {
	transaction, err := bdb.AssetCreationAndTransferE2E()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err})
	}
	//tran, _ := transaction.String()

	c.JSON(http.StatusOK, gin.H{"fileds": transaction})
	return
}

