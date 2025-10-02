package groups

import (
	"net/http"

	"example.com/m/v2/middlewares"
	"github.com/gin-gonic/gin"
)

type Goals struct {
	GoalLbs      float32 `json:"goalLbs" db:"goalLbs"`
	Multiplier   int     `json:"multiplier" db:"multiplier"`
	ErrMargin    float32 `json:"acceptablePercent" db:"acceptablePercent"`
	GoalsPerMeal string  `json:"goalsPerMeal" db:"goalsPerMeal"`
}

func goals(ctx *gin.Context) {
	userID := helper_GetUserID(ctx)
	var goals Goals
	if err := middlewares.Database.Get(&goals, "SELECT goalLbs, multiplier, acceptablePercent, goalsPerMeal FROM goals WHERE userID=?", userID); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, goals)
}

func InitGoalsAPI(group *gin.RouterGroup) {
	group.GET("/", CheckAuthenticated, goals)
}
