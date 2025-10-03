package groups

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ErrMsg struct {
	InvalidID string
	ExecError string
}

type DBStruct interface {
	GetIDKey() string
	GetID() string
}

func helper_GetUserID(ctx *gin.Context) string {
	return ctx.GetString("loggedInUser")
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
