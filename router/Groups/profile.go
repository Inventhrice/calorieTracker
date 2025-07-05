package groups

import (
	"net/http"
	"strings"

	"example.com/m/v2/middlewares"
	"github.com/gin-gonic/gin"
)

// type Credentials struct {
// 	Email string `json:"email" db:"email"`
// 	Password string `json:"password" db:"password"`
// }

type Credentials struct {
	Email string `json:"email"`
	Password string `json:"password"`
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

	if token, err := middlewares.AuthenticateUser(creds.Email, creds.Password); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	} else{
		ctx.JSON(http.StatusOK, gin.H{"email": creds.Email, "token": token})
		return
	}
}

func profile(ctx *gin.Context){
	if userID, exists := ctx.Get("loggedInUser"); !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"Error": "User is not authenticated."})
	} else{
		var user Profile
		if err := middlewares.Database.Get(&user, "SELECT id, firstname, lastname, pronouns FROM users WHERE id=?", userID); err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"Error": err.Error()})
			return 
		}

		ctx.JSON(http.StatusOK, user)
	}
}

func checkAuthenticated(ctx *gin.Context){
	header := ctx.GetHeader("Authorization")

	if(header == ""){
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Error": "Authorization header is missing"})
		return
	}

	authToken := strings.Split(header[1:len(header)-1], " ")
	if(len(authToken) != 2 || authToken[0] != "Bearer"){
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": "Invalid token format"})
		return
	}

	token := authToken[1]

	if userID, err := middlewares.CheckLoggedIn(token); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Error": "User is not logged in."})
		return
	} else {
		ctx.Set("loggedInUser", userID)
		ctx.Next()
	}
}

func InitProfileAPI(group *gin.RouterGroup){
	group.POST("/login", login)
	group.GET("/", checkAuthenticated, profile)
}

