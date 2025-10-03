package groups

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"example.com/m/v2/middlewares"
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

// func delete(ctx *gin.Context, T DBStruct, errs ErrMsg, idParam string, tableName string) error {
// 	id := helper_getIntFromStr(ctx, idParam, "Invalid entry ID")
// 	result, err := middlewares.Database.Exec("DELETE FROM "+tableName+" WHERE id = ?", id)
// 	if _, err := Helper_ExecError(result, err, errs.ExecError); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
// 		return err
// 	}

// 	ctx.Status(http.StatusOK)
// 	return nil
// }

func addUpdate(ctx *gin.Context, T DBStruct, errs ErrMsg, query string, isPatch bool) error {
	id := helper_getIntFromStr(ctx, T.GetIDKey(), errs.InvalidID)
	if err := ctx.BindJSON(&T); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return err
	}

	if string(id) != T.GetID() && isPatch {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": fmt.Sprintf("IDs do not match, aborting request. obj id: %v | request id: %v", id, T.GetID())})
	}

	result, err := middlewares.Database.NamedExec(query, T)
	if _, err := Helper_ExecError(result, err, errs.ExecError); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return err
	}

	if isPatch {
		ctx.Status(http.StatusOK)
	} else {
		newId, _ := result.LastInsertId()
		ctx.JSON(http.StatusOK, gin.H{"message": "Added successfully", "addedID": newId})
	}

	return nil
}

func get(ctx *gin.Context, query string, T DBStruct) error {
	if err := middlewares.Database.Get(&T, query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return err
	}
	ctx.JSON(http.StatusOK, T)
	return nil
}

func getAll(ctx *gin.Context, query string, T []DBStruct) error {
	if err := middlewares.Database.Select(&T, query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return err
	}
	ctx.JSON(http.StatusOK, T)
	return nil
}
