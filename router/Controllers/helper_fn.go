package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Gets userID from the gin.Context.
//
// If loggedInUser is empty, returns error: "There is no logged in user."
func helper_GetUserID(ctx *gin.Context) (string, error) {
	var err error
	id := ctx.GetString("loggedInUser")
	if id == "" {
		err = errors.New("There is no logged in user.")
	}
	return id, err
}

// Helper function that executes Atoi, and automatically sends a BadRequest reponse if there is an error. Returns the parsed int or -1 on an error.
func helper_getIntFromStr(idParam string) (int, error) {
	if err := Helper_CheckIDParam(idParam); err != nil {
		return -1, err
	} else {
		if id, err := strconv.Atoi(idParam); err != nil {
			return -1, err
		} else {
			return id, nil
		}
	}
}

func Helper_ctx400(ctx *gin.Context, errMsg string) {
	if errMsg != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": errMsg})
	}
}

// Checks for ID parameter
func Helper_CheckIDParam(id string) error {
	if id == "" {
		return errors.New("ID parameter is required")
	} else {
		return nil
	}
}
