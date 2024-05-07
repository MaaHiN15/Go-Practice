package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/MaaHiN15/go-practice/go-jwt/initializers"
	"github.com/MaaHiN15/go-practice/go-jwt/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(con *gin.Context) {
	unauthorized := func() {con.AbortWithStatus(http.StatusUnauthorized)}
	// Get Cookie
	tokenString, err := con.Cookie("Authorization")
	if err != nil {unauthorized()}
	// Parse/compare token with secret
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}; return []byte(os.Getenv("JWT_SECRET")), nil })
	if err != nil {unauthorized()}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// check for expiry
		if float64(time.Now().Unix()) > claims["exp"].(float64) {unauthorized()}
		// Get user
		var user models.User
		initializers.DB.First(&user, claims["sub"])
		if user.ID == 0 {unauthorized()}
		con.Set("user", user)
	} else {unauthorized()}
	// authorized
	con.Next()
};