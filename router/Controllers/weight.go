package controllers

import (
	"net/http"

	models "example.com/m/v2/Models"
	"github.com/gin-gonic/gin"
)

func getWeight(ctx *gin.Context) {
	var errmsg string
	if userID, err := helper_GetUserID(ctx); err != nil {
		errmsg = err.Error()
	} else {
		if weight, err := models.GetWeight(ctx.Param("start"), userID); err != nil {
			errmsg = err.Error()
		} else {
			ctx.JSON(http.StatusOK, weight)
			return
		}
	}
	Helper_ctx400(ctx, errmsg)
}

func addWeight(ctx *gin.Context) {
	var errmsg string
	if userID, err := helper_GetUserID(ctx); err != nil {
		errmsg = err.Error()
	} else {
		var weight models.WeightData
		if err := ctx.BindJSON(&weight); err != nil {
			errmsg = err.Error()
		} else {
			if err := models.AddUpdateWeight(userID); err != nil {
				errmsg = err.Error()
			} else {
				ctx.Status(http.StatusOK)
				return
			}
		}
	}

	Helper_ctx400(ctx, errmsg)
}

func InitWeightAPI(group *gin.RouterGroup) {
	group.GET("/:start", getWeight)
	group.POST("/", addWeight)
}
