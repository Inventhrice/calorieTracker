package groups

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
	if idParam, err := Helper_CheckIDParam(idParam); err != nil {
		id, err := strconv.Atoi(idParam)
		if err != nil {
			return -1, err
		}
		return id, nil
	} else {
		return -1, err
	}
}

func Helper_ctx400(ctx *gin.Context, errMsg string) {
	ctx.JSON(http.StatusBadRequest, gin.H{"Error": errMsg})
}

func Helper_CheckIDParam(id string) (string, error) {
	if id == "" {
		return "", errors.New("ID parameter is required")
	} else {
		return id, nil
	}

}
