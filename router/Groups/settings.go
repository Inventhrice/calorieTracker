package groups

import (
	"net/http"

	"example.com/m/v2/middlewares"
	"github.com/gin-gonic/gin"
)

var Settings map[string]string

type Goals struct{
	DateRecord string `json:"daterecord" db:"dateRecord"`
	GoalLbs float32 `json:"goallbs" db:"goallbs"`
	Multiplier int `json:"multiplier" db:"multiplier"`
	ErrMargin float32 `json:"acceptablePercent" db:"acceptablePercent"`
	goalsPerMeal string `json:"goalsPerMeal" db:"goalsPerMeal"`
}


func RefreshSettings(ctx *gin.Context) {
	rows, err := middlewares.Database.Query("SELECT keyName, value FROM settings")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}

	for rows.Next() {
		var keyname string
		var value string

		if err := rows.Scan(&keyname, &value); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}
		Settings[keyname] = value
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": Settings})
}

func initSettingsAPI(group *gin.RouterGroup){
	group.GET("/:setting") // if setting == all then return everything, else return value
	group.PATCH("/:setting")
	group.POST("/")
	group.DELETE("/:setting")
}

func initGoalsAPI(group *gin.RouterGroup){
	group.GET("/:daterecord") //QUERY goals table by date record, or get most recent entry
	group.PATCH("/:daterecord")
	group.POST("/")

}
