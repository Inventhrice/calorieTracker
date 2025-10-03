package controllers

import (
	"net/http"

	"example.com/m/v2/middlewares"
	"github.com/gin-gonic/gin"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Profile struct {
	ID        string `json:"id" db:"ID"`
	Firstname string `json:"firstname" db:"firstname"`
	Lastname  string `json:"lastname" db:"lastname"`
	Pronouns  string `json:"pronouns" db:"pronouns"`
	Username  string `json:"username" db:"username"`
}

func CheckAuthenticated(ctx *gin.Context) {
	token := ""

	if temp, err := ctx.Cookie("token"); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Error": err.Error()})
		return
	} else {
		token = temp
	}

	if userID, err := middlewares.CheckLoggedIn(token); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Error": "User is not logged in."})
		return
	} else {
		ctx.Set("loggedInUser", userID)
		ctx.Next()
	}
}

func Login(ctx *gin.Context) {
	var creds Credentials

	if err := ctx.BindJSON(&creds); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if token, err := middlewares.AuthenticateUser(creds.Username, creds.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	} else {
		ctx.SetCookie("token", token, 0, "", "", false, false)
		ctx.Status(http.StatusOK)
		return
	}
}

func updatePassword(ctx *gin.Context) {
	var errmsg string
	if userID, err := helper_GetUserID(ctx); err != nil {
		errmsg = err.Error()
	} else {
		creds := struct {
			Password string `json:"password"`
		}{""}

		if err := ctx.BindJSON(&creds); err != nil {
			errmsg = err.Error()
		} else {
			if err := middlewares.ChangePassword(userID, creds.Password); err != nil {
				errmsg = err.Error()
			} else {
				ctx.Status(http.StatusOK)
				return
			}
		}

	}
	Helper_ctx400(ctx, errmsg)
}

func Logout(ctx *gin.Context) {
	var errmsg string
	if userID, err := helper_GetUserID(ctx); err != nil {
		errmsg = err.Error()
	} else {
		middlewares.RemoveActiveSession(userID)
		ctx.Status(http.StatusOK)
	}
	Helper_ctx400(ctx, errmsg)
}

func profile(ctx *gin.Context) {
	var errmsg string
	if userID, err := helper_GetUserID(ctx); err != nil {
		errmsg = err.Error()
	} else {
		var user Profile
		if err := middlewares.Database.Get(&user, "SELECT firstname, lastname, pronouns, username FROM users WHERE ID=?", userID); err != nil {
			errmsg = err.Error()
		} else {
			ctx.JSON(http.StatusOK, user)
		}
	}
	Helper_ctx400(ctx, errmsg)
}

func updateProfile(ctx *gin.Context) {
	var errmsg string
	if userID, err := helper_GetUserID(ctx); err != nil {
		errmsg = err.Error()
	} else {
		var data Profile

		if err := ctx.BindJSON(&data); err != nil {
			errmsg = err.Error()
		} else {
			data.ID = userID
			if _, err := middlewares.Database.NamedExec("UPDATE users SET firstname=:firstname, lastname=:lastname, pronouns=:pronouns, username=:username WHERE ID=:ID", data); err != nil {
				errmsg = err.Error()
			} else {
				ctx.Status(http.StatusOK)
			}
		}
	}
	Helper_ctx400(ctx, errmsg)
}

func InitProfileAPI(group *gin.RouterGroup) {
	group.GET("/", CheckAuthenticated, profile)
	group.PATCH("/", CheckAuthenticated, updateProfile)
	group.PATCH("/password", CheckAuthenticated, updatePassword)
}
