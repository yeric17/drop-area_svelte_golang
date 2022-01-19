package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.MaxMultipartMemory = 1 << 20 // 1 MiB
	router.SetTrustedProxies([]string{"http://localhost:8080"})
	router.Static("/images", "./images")
	router.POST("/images", handleImage)

	router.Run(":7070")
}

func handleImage(g *gin.Context) {
	file, err := g.FormFile("file")
	if err != nil {
		g.String(400, "Bad request")
		fmt.Printf("Error: %v\n", err)
		return
	}

	err = g.SaveUploadedFile(file, "./images/"+file.Filename)
	if err != nil {
		g.String(500, "Server error")
		return
	}

	g.String(200, "File uploaded successfully")
}
