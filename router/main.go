package main

import (
	"fmt"
	"net/http"

	groups "example.com/m/v2/Groups"
	"example.com/m/v2/middlewares"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/api/settings", groups.RefreshSettings)
	router.StaticFile("/", "/app/public/index.html")
	router.StaticFile("/favicon.ico", "/app/public/favicon.ico")
	router.StaticFS("/assets", http.Dir("/app/public/assets"))

	router.POST("/login", groups.Login)
	router.POST("/logout", groups.CheckAuthenticated, groups.Logout)

	authorizedRoutes := router.Group("/api", groups.CheckAuthenticated)

	foodDBAPI := authorizedRoutes.Group("/foodDB")
	groups.InitFoodDBApi(foodDBAPI)

	entriesAPI := authorizedRoutes.Group("/entries")
	groups.InitEntriesApi(entriesAPI)

	weightAPI := authorizedRoutes.Group("/weight")
	groups.InitWeightAPI(weightAPI)

	profileAPI := authorizedRoutes.Group("/profile")
	groups.InitProfileAPI(profileAPI)

	goalsAPI := authorizedRoutes.Group("/goals")
	groups.InitGoalsAPI(goalsAPI)

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
