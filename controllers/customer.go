package Controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/20-VIGNESH-K/EnergyAuditing/models"
	"github.com/20-VIGNESH-K/EnergyAuditing/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.CreateUser
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println("123")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := services.CreateUser(user)
	if err != nil {
		if err.Error() == "invalid Email" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		if err.Error() == "invalid Name" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		if err.Error() == "invalid Phone Number" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		if err.Error() == "user already exists" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
        }
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var user models.LogIn

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.LogIn(user)
	if err != nil {
		if err.Error() == "user does not exist" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		} else if err.Error() == "password does not match" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Password does not match"})
		} else if err.Error() == "please provide a password" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Please provide a password"}) 
		}else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func Weaving(c *gin.Context) {
    var weave models.Weaving

    if err := c.ShouldBindJSON(&weave); err != nil {
		log.Println(err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    result,err := services.Weaving(weave)
    if err != nil {
        if err.Error() == "string value is not allowed" {
            c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
        }
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": result})
}

func Textile(c *gin.Context) {
    var textile models.Textile

    if err := c.ShouldBindJSON(&textile); err != nil {
		log.Println(err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    result,err := services.Textile(textile)
    if err != nil {
        if err.Error() == "string value is not allowed" {
            c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
        }
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": result})
}


func IT(c *gin.Context) {
    var it models.IT

    if err := c.ShouldBindJSON(&it); err != nil {
		log.Println(err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    result,err := services.IT(it)
    if err != nil {
        if err.Error() == "string value is not allowed" {
            c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
        }
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": result})
}

func GetUser(c *gin.Context){
	var user models.GetUser
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	data,message,err := services.GetUser(user)
	if err != nil{
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"error": message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": data})
}
