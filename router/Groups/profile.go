package groups

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
	ID        string `json:"id" db:"id"`
	Firstname string `json:"firstname" db:"firstname"`
	Lastname  string `json:"lastname" db:"lastname"`
	Pronouns  string `json:"pronouns" db:"pronouns"`
	Username string `json:"username" db:"username"`
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

func changepwd(ctx *gin.Context){
	userID := helper_GetUserID(ctx)
	creds := struct{Password string `json:"password"`}{""}

	if err := ctx.BindJSON(&creds); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if err := middlewares.ChangePassword(userID, creds.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	} else{
		ctx.Status(http.StatusOK)
	}
}

func profile(ctx *gin.Context) {
	userID := helper_GetUserID(ctx)
	var user Profile
	if err := middlewares.Database.Get(&user, "SELECT id, firstname, lastname, pronouns, username FROM users WHERE id=?", userID); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
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


func Logout(ctx *gin.Context){
	userID := helper_GetUserID(ctx)
	middlewares.RemoveActiveSession(userID)
	ctx.Status(http.StatusOK)
}

func InitProfileAPI(group *gin.RouterGroup) {
	group.GET("/", CheckAuthenticated, profile)
	group.PATCH("/password", CheckAuthenticated, changepwd)
}
