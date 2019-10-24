package openapi

import "github.com/gin-gonic/gin"

func Init() {
	router := NewRouter()
	router.Use(gin.Recovery())
	if err := router.Run(); err != nil {
		panic(err)
	}
}
