package groups

import (
	"fmt"
	"net/http"

	models "example.com/m/v2/Models"
	"example.com/m/v2/middlewares"
	"github.com/gin-gonic/gin"
)

func updateEntries(ctx *gin.Context) {
	if id, err := helper_getIntFromStr(ctx.Param("id")); err != nil {
		Helper_ctx400(ctx, "Invalid entry ID")
	} else {
		var entry models.Entry
		if err := ctx.BindJSON(&entry); err != nil {
			Helper_ctx400(ctx, err.Error())
		} else {
			if id == entry.ID {
				if err := models.M_updateEntry(entry); err != nil {
					Helper_ctx400(ctx, err.Error())
				} else {
					ctx.Status(http.StatusOK)
				}
			} else {
				Helper_ctx400(ctx, fmt.Sprintf("IDs do not match, aborting request. id: %v, entry.ID: %v", id, entry.ID))
			}
		}
	}
}

func addEntries(ctx *gin.Context) {
	var entry models.Entry
	if err := ctx.BindJSON(&entry); err != nil {
		Helper_ctx400(ctx, err.Error())
	} else {
		if entry.UserID, err = helper_GetUserID(ctx); err != nil {
			Helper_ctx400(ctx, err.Error())
		} else {
			if newID, err := models.M_addEntry(entry); err != nil {
				Helper_ctx400(ctx, err.Error())
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Entry added successfully", "addedID": newID})
			}
		}
	}
}

func getAllEntries(ctx *gin.Context) {
	if userID, err := helper_GetUserID(ctx); err != nil {
		Helper_ctx400(ctx, err.Error())
	} else {
		entries := []models.Entry{}
		if err := middlewares.Database.Select(&entries, "SELECT ID, dateRecord, meal, foodname, foodID, grams, cal, carbs, protein, fat, notes FROM processed_entries WHERE userid=?", userID); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, entries)
		}
	}
}

func getEntriesByWeek(ctx *gin.Context) {
	if userID, err := helper_GetUserID(ctx); err != nil {
		Helper_ctx400(ctx, err.Error())
	} else {
		if entries, err := models.M_getEntriesByWeek(ctx.Param("start"), ctx.Param("end"), userID); err != nil {
			Helper_ctx400(ctx, err.Error())
		} else {
			ctx.JSON(http.StatusOK, entries)
		}
	}
}

func deleteEntries(ctx *gin.Context) {
	if id, err := helper_getIntFromStr(ctx.Param("id")); err != nil {
		Helper_ctx400(ctx, "Invalid entry ID")
	} else {
		if err := models.M_deleteEntry(id); err != nil {
			Helper_ctx400(ctx, err.Error())
		} else {
			ctx.Status(http.StatusOK)
		}
	}
}

func InitEntriesApi(group *gin.RouterGroup) {
	group.POST("/", addEntries)
	group.DELETE("/:id", deleteEntries)
	group.PATCH("/:id", updateEntries)
	group.GET("/all", getAllEntries)
	group.GET("/:start/:end", getEntriesByWeek)

}
