package controllers

import (
	"net/http"

	"github.com/Naveenchand06/go-gin-jwt-auth/database"
	"github.com/Naveenchand06/go-gin-jwt-auth/models"
	"github.com/Naveenchand06/go-gin-jwt-auth/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func Signup(ctx *gin.Context) {
	var newUser models.User
	if err := ctx.BindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validate := validator.New()
	err := validate.Struct(newUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 12)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
	}
	newUser.Password = string(hashed)
	user, err := newUser.RegisterUser(database.GetDB())
	if err!= nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	token, err := utils.GetSignedJWTToken(user)
	if err!= nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
	response :=  map[string]interface{} {
		"user": user,
		"token": token,
	}
	ctx.JSON(http.StatusOK, response)
}

func Login(ctx *gin.Context) {
	var loginUser models.User
	if err := ctx.BindJSON(&loginUser); err!= nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	user, err := loginUser.LoginUser(database.GetDB())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	
	token, err := utils.GetSignedJWTToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := map[string]interface{}{
		"token": token,
		"user": user,
	}
	ctx.JSON(http.StatusOK, response)
}

