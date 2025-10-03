package groups

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func helper_GetUserID(ctx *gin.Context) string {
	return ctx.GetString("loggedInUser")
}

// Helper function that executes Atoi, and automatically sends a BadRequest reponse if there is an error. Returns the parsed int or -1 on an error.
func helper_getIntFromStr(context *gin.Context, idParam string, errMsg string) int {
	if idParam = helper_CheckIDParam(context, idParam); idParam != "" {
		id, err := strconv.Atoi(idParam)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
			return -1
		}
		return id
	}
	return -1
}

func helper_CheckIDParam(context *gin.Context, idParam string) string {
	if id := context.Param(idParam); id == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "ID parameter is required"})
		return ""
	} else {
		return id
	}

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

func Helper_BindObj(ctx *gin.Context, T any) error {
	if err := ctx.BindJSON(T); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return err
	}
	return nil
}
