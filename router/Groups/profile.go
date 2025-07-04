package groups;

import (
	"database/sql"
	"net/http"
	"time"

	"example.com/m/v2/middlewares"
	"github.com/gin-gonic/gin"
)


type Credentials struct {
	Email string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type Profile struct {
	ID string `json:"id" db:"id"`
	Firstname string `json:"firstname" db:"firstname"`
	Lastname string `json:"lastname" db:"lastname"`
	Pronouns string `json:"pronouns" db:"pronouns"`
}

func login(ctx *gin.Context){
	var creds Credentials

	if err := ctx.BindJSON(&creds); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	if err := middlewares.AuthenticateUser(creds.Email, creds.Password); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	ctx.Set("loggedInUser", creds.Email)
}

func profile(ctx *gin.Context){
	var user Profile
	email := 
	if err := middlewares.Database.Get(&user, "SELECT id, firstname, lastname, pronouns WHERE email=?", email); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Error": err.Error()})
		return 
	}

	ctx.JSON(http.StatusOK, user)
}

func initProfileAPI(group *gin.RouterGroup){
	group.POST("/login", login)
	group.GET("/profile", profile)
	group.PATCH("/profile")
}

