package main

import (
	"bytom-bdb/api"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("ginblock.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	// Use the following code if you need to write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()
	router.POST("/create_transaction", api.CreateTransaction)
	//router.POST("/block", api.TestBytomBlock)
	router.Run(":8090")
}

//bdb.SendTransaction()
//bdb.AssetCreationAndTransferE2E()
// hight, err := bdb.NewClientCreatesConnectionLocalhost()
// if err != nil {
// 	fmt.Println("请求错误", err)
// }
// fmt.Println(hight, "返回交易高度")
// //
// bdb.NewClientCreatesConnectionToTestNetwork()
// bdb.PostCreateTransaction()
