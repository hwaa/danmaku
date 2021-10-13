package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	//uses setHeader-middleware
	r.Use()
	// Get all public video's meta
	r.GET("/videoMetas", controllers.VideoC.GetPublicVideoMeta)

	// Upload video's metadata
	r.POST("/videoMetas", controllers.VideoC.PostVideoMetaData)

	// Upload video chunk content
	r.POST("/videoChunks", controllers.VideoC.PostVideoChunk)

	// Merge video chunks if all chunks are received, otherwise return the hashes of missing contents
	r.POST("/merge", controllers.VideoC.Merge)

	// Serving static storage file
	r.Static("/storage", "storage")

	log.Fatalln(r.Run()) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}