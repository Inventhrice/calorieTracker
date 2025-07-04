package groups;

import (
	"database/sql"
	"net/http"
	"time"

	"example.com/m/v2/middlewares"
	"github.com/gin-gonic/gin"
)

type Profile struct {
	ID string `json:"id" db:"id"`
	Firstname string `json:"firstname" db:"firstname"`
	Lastname string `json:"lastname" db:"lastname"`
	Pronouns string `json:"pronouns" db:"pronouns"`
}

type Settings struct {

}

type Goals struct{
	DateRecord string `json:"daterecord" db:"dateRecord"`
	GoalLbs float32 `json:"goallbs" db:"goallbs"`
	Multiplier int `json:"multiplier" db:"multiplier"`
	ErrMargin float32 `json:"acceptablePercent" db:"acceptablePercent"`
	goalsPerMeal string `json:"goalsPerMeal" db:"goalsPerMeal"`
}

func initProfileAPI(group *gin.RouterGroup){
	group.GET("/")
	group.POST("/login")
	group.PATCH("/update")
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
