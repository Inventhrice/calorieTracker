package groups

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"example.com/m/v2/middlewares"
	"github.com/gin-gonic/gin"
)

type Entry struct {
	UserID     string        `db:"userid"`
	ID         int           `json:"id" db:"ID"`
	Meal       string        `json:"meal" db:"meal"`
	DateRecord string        `json:"daterecord" db:"dateRecord"`
	FoodID     sql.NullInt32 `json:"foodID" db:"foodID"`
	Foodname   string        `json:"foodname" db:"foodname"`
	Quantity   float32       `json:"quantity" db:"grams"`
	Cal        float32       `json:"cal" db:"cal"`
	Protein    float32       `json:"protein" db:"protein"`
	Fat        float32       `json:"fat" db:"fat"`
	Carbs      float32       `json:"carbs" db:"carbs"`
	Notes      string        `json:"notes" db:"notes"`
}

func updateEntries(ctx *gin.Context) {
	id := helper_getIntFromStr(ctx, ctx.Param("id"), "Invalid entry ID")
	var entry Entry
	if err := ctx.BindJSON(&entry); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if id == entry.ID {
		result, err := middlewares.Database.NamedExec("UPDATE entries SET dateRecord=:dateRecord,meal=:meal,foodID=:foodID,grams=:grams,foodname=:foodname,cal=:cal,protein=:protein,fat=:fat,carbs=:carbs WHERE ID=:ID", entry)
		if _, err := Helper_ExecError(result, err, "No entry with the provided ID found"); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}
		ctx.Status(http.StatusOK)
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": fmt.Sprintf("IDs do not match, aborting request. id: %v, entry.ID: %v", id, entry.ID)})
	}
}

func addEntries(ctx *gin.Context) {
	var entry Entry
	if err := ctx.BindJSON(&entry); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	entry.UserID = helper_GetUserID(ctx)

	result, err := middlewares.Database.NamedExec("INSERT INTO entries (userid, dateRecord, meal, foodID, foodname, grams, cal, protein, fat, carbs) VALUES (:userid, :dateRecord, :meal, :foodID, :foodname, :grams, :cal, :protein, :fat, :carbs)", entry)

	if _, err := Helper_ExecError(result, err, "Entry was unable to be added"); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newId, _ := result.LastInsertId()
	ctx.JSON(http.StatusOK, gin.H{"message": "Entry added successfully", "addedID": newId})
}

func getAllEntries(ctx *gin.Context) {
	userID := helper_GetUserID(ctx)
	entries := []Entry{}
	if err := middlewares.Database.Select(&entries, "SELECT ID, dateRecord, meal, foodname, foodID, grams, cal, carbs, protein, fat, notes FROM processed_entries WHERE userid=?", userID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, entries)
}

func getEntriesByWeek(ctx *gin.Context) {
	userID := helper_GetUserID(ctx)
	entries := []Entry{}
	startDate, _ := time.Parse(time.DateOnly, ctx.Param("start"))
	endDate, _ := time.Parse(time.DateOnly, ctx.Param("end"))

	if err := middlewares.Database.Select(&entries, "SELECT ID, dateRecord, meal, foodname, foodID, grams, cal, carbs, protein, fat, notes FROM processed_entries WHERE dateRecord BETWEEN ? AND ? AND userid=?", startDate.Format(time.DateOnly), endDate.Format(time.DateOnly), userID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, entries)
}

func deleteEntries(ctx *gin.Context) {
	id := helper_getIntFromStr(ctx, ctx.Param("id"), "Invalid entry ID")
	result, err := middlewares.Database.Exec("DELETE FROM entries WHERE id = ?", id)
	if _, err := Helper_ExecError(result, err, "No entry with the provided ID found"); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func InitEntriesApi(group *gin.RouterGroup) {
	group.POST("/", addEntries)
	group.DELETE("/:id", deleteEntries)
	group.PATCH("/:id", updateEntries)
	group.GET("/all", getAllEntries)
	group.GET("/:start/:end", getEntriesByWeek)

}
