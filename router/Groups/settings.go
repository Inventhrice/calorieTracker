package groups

import (
	"net/http"

	"example.com/m/v2/middlewares"
	"github.com/gin-gonic/gin"
)

var Settings map[string]string

type Goals struct {
	GoalLbs      float32 `json:"goalLbs" db:"goalLbs"`
	Multiplier   int     `json:"multiplier" db:"multiplier"`
	ErrMargin    float32 `json:"acceptablePercent" db:"acceptablePercent"`
	GoalsPerMeal string  `json:"goalsPerMeal" db:"goalsPerMeal"`
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

func goals(ctx *gin.Context){
	userID := helper_GetUserID(ctx)
	var goals Goals
	if err := middlewares.Database.Get(&goals, "SELECT goalLbs, multiplier, acceptablePercent, goalsPerMeal FROM goals WHERE userID=?",userID); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, goals)
}


func initSettingsAPI(group *gin.RouterGroup) {
	group.GET("/:setting") // if setting == all then return everything, else return value
	group.PATCH("/:setting")
	group.POST("/")
	group.DELETE("/:setting")
}

func InitGoalsAPI(group *gin.RouterGroup) {
	group.GET("/", goals)
}
