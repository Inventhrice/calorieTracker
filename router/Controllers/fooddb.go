package controllers

import (
	"net/http"

	models "example.com/m/v2/Models"
	"github.com/gin-gonic/gin"
)

func updateFood(ctx *gin.Context) {
	var errmsg string
	if id, err := helper_getIntFromStr(ctx.Param("id")); err != nil {
		errmsg = "Invalid food info ID"
	} else {
		var food models.FoodInfo
		if err := ctx.BindJSON(&food); err != nil {
			errmsg = "Invalid JSON data: " + err.Error()
		} else {
			food.ID = id
			if err := models.UpdateFood(food); err != nil {
				errmsg = err.Error()
			} else {
				ctx.Status(http.StatusOK)
				return
			}
		}
	}

	Helper_ctx400(ctx, errmsg)
}

func addFood(ctx *gin.Context) {
	var errmsg string
	var food models.FoodInfo

	if err := ctx.BindJSON(&food); err != nil {
		errmsg = "Invalid JSON data: " + err.Error()
	} else {
		if userid, err := helper_GetUserID(ctx); err != nil {
			errmsg = err.Error()
		} else {
			food.UserID = userid
			if newID, err := models.AddFood(food); err != nil {
				errmsg = err.Error()
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Food info added successfully", "addedID": newID})
				return
			}
		}
	}
	Helper_ctx400(ctx, errmsg)
}

func getFood(ctx *gin.Context) {
	var errmsg string
	if userID, err := helper_GetUserID(ctx); err != nil {
		errmsg = err.Error()
	} else {
		if id, err := helper_getIntFromStr(ctx.Param("id")); err != nil {
			errmsg = "Invalid food info ID"
		} else {
			if food, err := models.GetFood(id, userID); err != nil {
				errmsg = err.Error()

			} else {
				ctx.JSON(http.StatusOK, gin.H{"data": food})
			}
		}
	}
	Helper_ctx400(ctx, errmsg)
}

func deleteFood(ctx *gin.Context) {
	var errmsg string

	if id, err := helper_getIntFromStr(ctx.Param("id")); err != nil {
		errmsg = "Invalid food info ID"
	} else {
		if err := models.DeleteFood(id); err != nil {
			errmsg = err.Error()
		} else {
			ctx.Status(http.StatusOK)
			return
		}
	}
	Helper_ctx400(ctx, errmsg)
}

func getAllFoods(ctx *gin.Context) {
	var errmsg string
	if userid, err := helper_GetUserID(ctx); err != nil {
		errmsg = err.Error()
	} else {
		if listFoods, err := models.GetAllFoods(userid); err != nil {
			errmsg = err.Error()
		} else {
			ctx.JSON(http.StatusOK, listFoods)
			return
		}
	}
	Helper_ctx400(ctx, errmsg)
}

func InitFoodDBApi(group *gin.RouterGroup) {
	group.POST("/", addFood)
	group.GET("/:id", getFood)
	group.GET("/all", getAllFoods)
	group.DELETE("/:id", deleteFood)
	group.PATCH("/:id", updateFood)
}
