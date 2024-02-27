package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var imageContainer *ImageContainer

func init() {
	imageContainer = NewImageContainer("poloxue/public_images", "main")
}

func GetRandomImage(c *gin.Context) {
	category := c.Param("category")
	fmt.Println(category)
	imageURL, err := imageContainer.RandomImage(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching image URL"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"image": imageURL})
}

func main() {
	router := gin.Default()
	router.GET("/image/random/:category", GetRandomImage)
	_ = router.Run(":8080")
}
