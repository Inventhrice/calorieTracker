package groups

import (
	"net/http"

	"example.com/m/v2/middlewares"
	"github.com/gin-gonic/gin"
)

var Settings map[string]string

func RefreshSettings(ctx *gin.Context) {
	rows, err := middlewares.Database.Query("SELECT keyName, value FROM settings")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}

	for rows.Next() {
		var keyname string
		var value string

		if err := rows.Scan(&keyname, &value); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}
		Settings[keyname] = value
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": Settings})
}
