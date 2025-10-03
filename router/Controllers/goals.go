package controllers

import (
	"net/http"

	"example.com/m/v2/middlewares"
	"github.com/gin-gonic/gin"
)

type Goals struct {
	GoalLbs        float32 `json:"goalLbs" db:"goalLbs"`
	Multiplier     int     `json:"multiplier" db:"multiplier"`
	ErrMargin      float32 `json:"acceptablePercent" db:"acceptablePercent"`
	GoalsPerMeal   string  `json:"goalsPerMeal" db:"goalsPerMeal"`
	ProteinGPerLbs float32 `json:"ProteinGPerLBS" db:"proteinGPerLBS"`
	FatGPerLbs     float32 `json:"FatGPerLBS" db:"fatGPerLBS"`
	UserID         string  `db:"userid"`
}

func getGoal(ctx *gin.Context) {
	userID := helper_GetUserID(ctx)
	var goals Goals
	if err := middlewares.Database.Get(&goals, "SELECT goalLbs, multiplier, acceptablePercent, goalsPerMeal FROM goals WHERE userID=?", userID); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, goals)
}

func addGoal(ctx *gin.Context) {
	userID := helper_GetUserID(ctx)
	var goals Goals
	goals.UserID = userID

	if err := ctx.BindJSON(&goals); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	result, err := middlewares.Database.NamedExec("INSERT INTO goals (goalLbs, multiplier, acceptablePercent, goalsPerMeal, proteinGPerLBS, fatGPerLBS, userid) VALUES (:goalLbs, :multiplier, :acceptablePercent, :goalsPerMeal, :proteinGPerLBS, :fatGPerLBS, :userid)", goals)

	if _, err := Helper_ExecError(result, err, "Goals Entry/Info was unable to be added"); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Goal added successfully"})

}

func InitGoalsAPI(group *gin.RouterGroup) {
	group.GET("/", CheckAuthenticated, getGoal)
	group.POST("/", CheckAuthenticated, addGoal)
}
