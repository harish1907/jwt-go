package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/harish1907/jwt-go/intializers"
	"github.com/harish1907/jwt-go/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	// Get the email and password in request body.
	var body struct {
		Email    string
		Password string
	}
	c.Bind(&body)
	if body.Email == "" || body.Password == "" {
		c.JSON(404, gin.H{
			"message": "Email and password is requried..",
		})
		return
	}

	//hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Failed to creating hash the password..",
		})
		return
	}
	// Creating a User
	jwtuser := models.JwtUser{Email: body.Email, Password: string(hash)}
	result := intializers.DB.Create(&jwtuser)

	if result.Error != nil {
		c.JSON(404, gin.H{
			"message": "Failed to creating user..",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User creating Successfully.",
	})
}

func Login(c *gin.Context) {
	// Get the email and password in request body.
	var body struct {
		Email    string
		Password string
	}

	c.Bind(&body)
	if body.Email == "" || body.Password == "" {
		c.JSON(404, gin.H{
			"message": "Email and password is requried..",
		})
		return
	}

	//Look up request body
	var user models.JwtUser
	intializers.DB.First(&user, "email=?", body.Email)

	if user.ID == 0 {
		c.JSON(404, gin.H{
			"message": "Invalied email or password.",
		})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Invalied email or password.",
		})
		return
	}

	// Generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenstring, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Failed to create token",
		})
		return
	}

	//set cokkies
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenstring, 3600*24*30, "", "", false, true)

	// response
	c.JSON(200, gin.H{
		// "token": tokenstring,
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(200, gin.H{
		"result": "I am logged in",
		"user":   user,
	})
}
