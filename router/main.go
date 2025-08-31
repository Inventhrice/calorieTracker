package main

import (
	"fmt"
	"net/http"

	groups "example.com/m/v2/Groups"
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
	router.GET("/api/settings", groups.RefreshSettings)
	initStaticRoutes(router)

	router.POST("/login", groups.Login)
	router.POST("/logout", groups.CheckAuthenticated, groups.Logout)

	authorizedRoutes := router.Group("/api", groups.CheckAuthenticated)

	groups.InitFoodDBApi(authorizedRoutes.Group("/foodDB"))
	groups.InitEntriesApi(authorizedRoutes.Group("/entries"))
	groups.InitWeightAPI(authorizedRoutes.Group("/weight"))
	groups.InitProfileAPI(authorizedRoutes.Group("/profile"))
	groups.InitGoalsAPI(authorizedRoutes.Group("/goals"))

	return router
}

func main() {
	if err := middlewares.InitDB(); err != nil {
		fmt.Println(err)
		return
	}
	defer middlewares.Database.Close()
	groups.Settings = make(map[string]string)
	router := InitRouter()
	router.Run()
}
