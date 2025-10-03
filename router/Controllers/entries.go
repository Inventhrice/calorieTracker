package groups

import (
	"fmt"
	"net/http"

	models "example.com/m/v2/Models"
	"example.com/m/v2/middlewares"
	"github.com/gin-gonic/gin"
)

func updateEntries(ctx *gin.Context) {
	var errmsg string

	if id, err := helper_getIntFromStr(ctx.Param("id")); err != nil {
		errmsg = "Invalid entry ID"
	} else {
		var entry models.Entry
		if err := ctx.BindJSON(&entry); err != nil {
			errmsg = err.Error()
		} else {
			if id == entry.ID {
				if err := models.M_updateEntry(entry); err != nil {
					errmsg = err.Error()
				} else {
					ctx.Status(http.StatusOK)
					return
				}
			} else {
				errmsg = fmt.Sprintf("IDs do not match, aborting request. id: %v, entry.ID: %v", id, entry.ID)
			}
		}
	}
	Helper_ctx400(ctx, errmsg)
}

func addEntries(ctx *gin.Context) {
	var errmsg string
	var entry models.Entry
	if err := ctx.BindJSON(&entry); err != nil {
		errmsg = err.Error()
	} else {
		if entry.UserID, err = helper_GetUserID(ctx); err != nil {
			errmsg = err.Error()
		} else {
			if newID, err := models.M_addEntry(entry); err != nil {
				errmsg = err.Error()
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Entry added successfully", "addedID": newID})
				return
			}
		}
	}
	Helper_ctx400(ctx, errmsg)
}

func getAllEntries(ctx *gin.Context) {
	var errmsg string
	if userID, err := helper_GetUserID(ctx); err != nil {
		errmsg = err.Error()
	} else {
		entries := []models.Entry{}
		if err := middlewares.Database.Select(&entries, "SELECT ID, dateRecord, meal, foodname, foodID, grams, cal, carbs, protein, fat, notes FROM processed_entries WHERE userid=?", userID); err != nil {
			errmsg = err.Error()
		} else {
			ctx.JSON(http.StatusOK, entries)
			return
		}
	}

	Helper_ctx400(ctx, errmsg)
}

func getEntriesByWeek(ctx *gin.Context) {
	var errmsg string
	if userID, err := helper_GetUserID(ctx); err != nil {
		errmsg = err.Error()
	} else {
		if entries, err := models.M_getEntriesByWeek(ctx.Param("start"), ctx.Param("end"), userID); err != nil {
			errmsg = err.Error()
		} else {
			ctx.JSON(http.StatusOK, entries)
			return
		}
	}
	Helper_ctx400(ctx, errmsg)
}

func deleteEntries(ctx *gin.Context) {
	var errmsg string
	if id, err := helper_getIntFromStr(ctx.Param("id")); err != nil {
		errmsg = "Invalid entry ID"
	} else {
		if err := models.M_deleteEntry(id); err != nil {
			errmsg = err.Error()
		} else {
			ctx.Status(http.StatusOK)
			return
		}
	}
	Helper_ctx400(ctx, errmsg)
}

func InitEntriesApi(group *gin.RouterGroup) {
	group.POST("/", addEntries)
	group.DELETE("/:id", deleteEntries)
	group.PATCH("/:id", updateEntries)
	group.GET("/all", getAllEntries)
	group.GET("/:start/:end", getEntriesByWeek)

}
