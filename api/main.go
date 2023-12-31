package main

import (
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/datagenx/license-generator/api/controllers"
	"github.com/datagenx/license-generator/internal/initializers"
	"github.com/gin-gonic/gin"
)

func main() {

	// initilization
	runtime.GOMAXPROCS(runtime.NumCPU())
	initializers.LoadEnvVar()

	if strings.ToUpper(os.Getenv("ENV_TYPE")) == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}
	//

	router := gin.Default()

	router.GET("/", welcome)
	router.POST("/", welcome)
	router.POST("/generate", controllers.PostGenerate)
	router.POST("/validate", controllers.PostValidate)

	router.Run()
}

// welcome msg
func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to License-GO",
	})
}
