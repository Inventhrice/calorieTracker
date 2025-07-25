package groups

import (
	"database/sql"
	"net/http"
	"time"

	"example.com/m/v2/middlewares"
	"github.com/gin-gonic/gin"
)

type WeightData struct {
	UserID string  `db:"userid"`
	Date   string  `json:"daterecord" db:"dateRecord"`
	KG     float32 `json:"kg" db:"kg"`
}

func getWeight(ctx *gin.Context) {
	userID := helper_GetUserID(ctx)
	weight := WeightData{}
	startDate, _ := time.Parse(time.DateOnly, ctx.Param("start"))

	if err := middlewares.Database.Get(&weight, "SELECT dateRecord, kg FROM weightTrack WHERE dateRecord=? AND userid=?", startDate.Format(time.DateOnly), userID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, weight)
}

func addWeight(ctx *gin.Context) {
	userID := helper_GetUserID(ctx)
	var weight WeightData
	mode := "PATCH"
	if err := ctx.BindJSON(&weight); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	var temp float32
	if err := middlewares.Database.QueryRow("SELECT kg FROM weightTrack WHERE dateRecord=? AND userid=?;", weight.Date, userID).Scan(&temp); err != nil {
		if err == sql.ErrNoRows {
			mode = "POST"
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}
	}
	weight.UserID = userID
	var stmt string
	if mode == "POST" {

		stmt = "INSERT INTO weightTrack (userid, dateRecord, kg) VALUES (:userid, :dateRecord, :kg)"
	} else if mode == "PATCH" {
		stmt = "UPDATE weightTrack SET kg=:kg WHERE dateRecord=:dateRecord AND userid=:userid"
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "How the hell did you even get this. mode wasn't POST or PATCH"})
	}

	// This is haram and im aware of it.
	result, err := middlewares.Database.NamedExec(stmt, weight)

	if _, err := Helper_ExecError(result, err, "Information was unable to be added"); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func InitWeightAPI(group *gin.RouterGroup) {
	group.GET("/:start", getWeight)
	group.POST("/", addWeight)
}
