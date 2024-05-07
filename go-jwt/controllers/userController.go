package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/MaaHiN15/go-practice/go-jwt/initializers"
	"github.com/MaaHiN15/go-practice/go-jwt/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func PingPong(con *gin.Context) { con.JSON(200, gin.H{"message": "pong"}) };

func SignUp(con *gin.Context) {
	// Get req body
	var PostBody struct { Email string; Password string }
	if con.Bind(&PostBody) != nil { con.JSON(http.StatusBadRequest, gin.H{ "error" : "failed to get req body"}); return }
	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(PostBody.Password), 10)
	if err != nil { con.JSON(http.StatusBadRequest, gin.H{ "error" : "failed to hash password"}); return }
	// Create user
	user := models.User{Email: PostBody.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)
	if result.Error != nil { con.JSON(http.StatusBadRequest, gin.H{ "error" : "failed to create user"}); return }
	con.JSON(http.StatusOK, gin.H{"Success" : PostBody.Email+" was created"})
};


func Login(con *gin.Context) {
	// Get req body
	var PostBody struct { Email string; Password string }
	if con.Bind(&PostBody) != nil { con.JSON(http.StatusBadRequest, gin.H{"error" : "failed to get req body"}); return }
	// find user
	var user models.User
	initializers.DB.First(&user, "Email = ?", PostBody.Email)
	if user.ID == 0 { con.JSON(http.StatusBadRequest, gin.H{"error" : "Invalid username & password"}); return }
	// Password comparison
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(PostBody.Password))
	if err != nil { con.JSON(http.StatusBadRequest, gin.H{"error" : "Invalid username & password"}); return }
	// create jwt token // making expiry as 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{ "sub": user.ID, "exp": time.Now().Add(time.Hour * 1).Unix() })
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil { con.JSON(http.StatusBadRequest, gin.H{"error" : "failed to create token" }); return }
	con.SetSameSite(http.SameSiteLaxMode)
	con.SetCookie("Authorization", tokenString, 3600 * 24, "", "", false, true)
	con.JSON(http.StatusOK, gin.H{"status" : "success"})
};

func Validate(con *gin.Context) {
	user, _ := con.Get("user")
	con.JSON(http.StatusOK, gin.H{"status" : "logged in", "user" : user})
};