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
	router.StaticFS("/app", http.Dir("/app/public/"))
	router.StaticFile("/favicon.ico", "/app/public/favicon.ico")

	router.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "/app/login.html")
	})

	router.POST("/login", groups.Login)

	authorizedRoutes := router.Group("/api", groups.CheckAuthenticated)

	foodDBAPI := authorizedRoutes.Group("/foodDB")
	groups.InitFoodDBApi(foodDBAPI)

	entriesAPI := authorizedRoutes.Group("/entries")
	groups.InitEntriesApi(entriesAPI)

	weightAPI := authorizedRoutes.Group("/weight")
	groups.InitWeightAPI(weightAPI)

	profileAPI := authorizedRoutes.Group("/profile")
	groups.InitProfileAPI(profileAPI)
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
