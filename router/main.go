package main

import (
	"fmt"
	"net/http"

	controllers "example.com/m/v2/Controllers"
	"example.com/m/v2/middlewares"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func initStaticRoutes(router *gin.Engine) {
	router.StaticFile("/", "/app/public/index.html")
	router.StaticFile("/favicon.ico", "/app/public/favicon.ico")
	router.StaticFile("/serviceworker.js", "/app/public/serviceworker.js")
	router.StaticFS("/assets", http.Dir("/app/public/assets"))
}

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/api/settings", controllers.RefreshSettings)
	initStaticRoutes(router)

	router.POST("/login", controllers.Login)
	router.POST("/logout", controllers.CheckAuthenticated, controllers.Logout)

	authorizedRoutes := router.Group("/api", controllers.CheckAuthenticated)

	controllers.InitFoodDBApi(authorizedRoutes.Group("/foodDB"))
	controllers.InitEntriesApi(authorizedRoutes.Group("/entries"))
	controllers.InitWeightAPI(authorizedRoutes.Group("/weight"))
	controllers.InitProfileAPI(authorizedRoutes.Group("/profile"))
	controllers.InitGoalsAPI(authorizedRoutes.Group("/goals"))

	return router
}

func main() {
	if err := middlewares.InitDB(); err != nil {
		fmt.Println(err)
		return
	}
	defer middlewares.Database.Close()
	controllers.Settings = make(map[string]string)
	router := InitRouter()
	router.Run()
}
