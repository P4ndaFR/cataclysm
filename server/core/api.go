package core

import gin "gopkg.in/gin-gonic/gin.v1"

// APIStart start the main API
func APIStart(l string) (err error) {

	r := gin.Default()
	r.GET("/health", handler.Health)
	r.Run(l)

}
