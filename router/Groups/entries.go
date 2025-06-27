package groups

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"example.com/m/v2/middlewares"
	"github.com/gin-gonic/gin"
)

type Entry struct {
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

// Helper function that executes Atoi, and automatically sends a BadRequest reponse if there is an error. Returns the parsed int or -1 on an error.
func helper_getIntFromStr(context *gin.Context, idParam string, errMsg string) int {
	if idParam == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "ID parameter is required"})
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return -1
	}
	return id
}

func Helper_ExecError(r sql.Result, initerr error, noRowsFound_errMsg string) (int64, error) {
	// CHecks if the initial call to Query has an error
	if initerr != nil {
		return 0, initerr
	}

	// Checks how many rows are affected, and returns nil if there's an error
	rowsAffected, err := r.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rowsAffected == 0 {
		return 0, errors.New(noRowsFound_errMsg)
	}

	return rowsAffected, nil
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

	result, err := middlewares.Database.NamedExec("INSERT INTO entries (dateRecord, meal, foodID, foodname, grams, cal, protein, fat, carbs) VALUES (:dateRecord, :meal, :foodID, :foodname, :grams, :cal, :protein, :fat, :carbs)", entry)

	if _, err := Helper_ExecError(result, err, "Entry was unable to be added"); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newId, _ := result.LastInsertId()
	ctx.JSON(http.StatusOK, gin.H{"message": "Entry added successfully", "addedID": newId})
}

func getAllEntries(ctx *gin.Context) {
	entries := []Entry{}
	if err := middlewares.Database.Select(&entries, "SELECT * FROM processed_entries"); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, entries)
}

func getEntriesByWeek(ctx *gin.Context) {
	entries := []Entry{}
	startDate, _ := time.Parse(time.DateOnly, ctx.Param("start"))
	endDate, _ := time.Parse(time.DateOnly, ctx.Param("end"))

	if err := middlewares.Database.Select(&entries, "SELECT * FROM processed_entries WHERE dateRecord BETWEEN ? AND ?", startDate.Format(time.DateOnly), endDate.Format(time.DateOnly)); err != nil {
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
