package routes

import (
	. "projects/log_files_obfuscator/controllers"

	"github.com/gin-gonic/gin"
)

func RunRoutes() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/user/:key/:value", User)
	router.GET("/admin/:key/:value", Admin)
	router.Run(":3000")
}
