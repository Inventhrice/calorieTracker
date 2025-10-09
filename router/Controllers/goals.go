package controllers

import (
	"net/http"

	models "example.com/m/v2/Models"
	"example.com/m/v2/middlewares"
	"github.com/gin-gonic/gin"
)

type Goals struct {
	GoalLbs        float32 `json:"goalLbs" db:"goalLbs"`
	Multiplier     int     `json:"multiplier" db:"multiplier"`
	ErrMargin      float32 `json:"acceptablePercent" db:"acceptablePercent"`
	GoalsPerMeal   string  `json:"goalsPerMeal" db:"goalsPerMeal"`
	ProteinGPerLbs float32 `json:"proteinGPerLBS" db:"proteinGPerLBS"`
	FatGPerLbs     float32 `json:"fatGPerLBS" db:"fatGPerLBS"`
	UserID         string  `db:"userid"`
}

func getGoal(ctx *gin.Context) {
	var errmsg string
	if userID, err := helper_GetUserID(ctx); err != nil {
		errmsg = err.Error()
	} else {
		var goals Goals
		if err := middlewares.Database.Get(&goals, "SELECT goalLbs, multiplier, acceptablePercent, goalsPerMeal, proteinGPerLBS, fatGPerLBS FROM goals WHERE userID=?  ORDER BY ID DESC LIMIT 1;", userID); err != nil {
			errmsg = err.Error()
		} else {
			ctx.JSON(http.StatusOK, goals)
			return
		}
	}
	Helper_ctx400(ctx, errmsg)
}

func addGoal(ctx *gin.Context) {
	var errmsg string
	if userID, err := helper_GetUserID(ctx); err != nil {
		errmsg = err.Error()
	} else {
		var goals Goals
		goals.UserID = userID

		if err := ctx.BindJSON(&goals); err != nil {
			errmsg = err.Error()
		} else {
			result, err := middlewares.Database.NamedExec("INSERT INTO goals (goalLbs, multiplier, acceptablePercent, goalsPerMeal, proteinGPerLBS, fatGPerLBS, userid) VALUES (:goalLbs, :multiplier, :acceptablePercent, :goalsPerMeal, :proteinGPerLBS, :fatGPerLBS, :userid)", goals)

			if _, err := models.Helper_ExecError(result, err, "Goals Entry/Info was unable to be added"); err != nil {
				errmsg = err.Error()
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Goal added successfully"})
				return
			}
		}
	}
	Helper_ctx400(ctx, errmsg)

}

func InitGoalsAPI(group *gin.RouterGroup) {
	group.GET("/", CheckAuthenticated, getGoal)
	group.POST("/", CheckAuthenticated, addGoal)
}
