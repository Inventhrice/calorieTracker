package groups

import (
	"database/sql"
	"net/http"

	"example.com/m/v2/middlewares"
	"github.com/gin-gonic/gin"
)

type FoodInfo struct {
	ID          int     `json:"id" db:"ID"`
	Name        string  `json:"name" db:"name"`
	CalPerG     float32 `json:"calperg" db:"calPerG"`
	ProteinPerG float32 `json:"proteinperg" db:"proteinPerG"`
	FatPerG     float32 `json:"fatperg" db:"fatPerG"`
	CarbPerG    float32 `json:"carbperg" db:"carbPerG"`
	Notes       string  `json:"notes"  db:"notes"`
	Source      string  `json:"source"  db:"source"`
}

func updateFood(ctx *gin.Context) {
	id := helper_getIntFromStr(ctx, ctx.Param("id"), "Invalid food info ID")

	var food FoodInfo
	if err := ctx.BindJSON(&food); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data: " + err.Error()})
		return
	}
	food.ID = id

	result, err := middlewares.Database.NamedExec("UPDATE food_info SET name=:name, calPerG=:calPerG, proteinPerG=:proteinPerG, fatPerG=:fatPerG, carbPerG=:carbPerG, notes=:notes WHERE id = :ID", food)
	if _, err := Helper_ExecError(result, err, "No food info with the provided ID found"); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func addFood(ctx *gin.Context) {
	var food FoodInfo
	if err := ctx.BindJSON(&food); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := middlewares.Database.NamedExec("INSERT INTO food_info (name, calPerG, proteinPerG, fatPerG, carbPerG, notes, source) VALUES (:name, :calPerG, :proteinPerG, :fatPerG, :carbPerG, :notes, :source)", food)
	if _, err := Helper_ExecError(result, err, "Food info was unable to be added"); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newId, _ := result.LastInsertId()
	ctx.JSON(http.StatusOK, gin.H{"message": "Food info added successfully", "addedID": newId})
}

func getFood(ctx *gin.Context) {
	id := helper_getIntFromStr(ctx, ctx.Param("id"), "Invalid food info ID")
	var food FoodInfo
	if err := middlewares.Database.Get(&food, "SELECT * FROM food_info WHERE id=?", id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": food})
}

func deleteFood(ctx *gin.Context) {
	id := helper_getIntFromStr(ctx, ctx.Param("id"), "Invalid food info ID")
	result, err := middlewares.Database.Exec("DELETE FROM food_info WHERE id = ?", id)
	if _, err := Helper_ExecError(result, err, "No food info with the provided ID found"); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func getAllFoodIDs() ([]sql.NullInt16, error) {
	listFoodIDs := make([]sql.NullInt16, 0)

	if err := middlewares.Database.Get(&listFoodIDs, "SELECT id FROM food_info"); err != nil {
		return nil, err
	}
	return listFoodIDs, nil
}

func getAllFoods(ctx *gin.Context) {
	listFoods := []FoodInfo{}
	if err := middlewares.Database.Select(&listFoods, "SELECT * FROM food_info"); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, listFoods)
}

func InitFoodDBApi(group *gin.RouterGroup) {
	group.POST("/", addFood)
	group.GET("/:id", getFood)
	group.GET("/all", getAllFoods)
	group.DELETE("/:id", deleteFood)
	group.PATCH("/:id", updateFood)
}
