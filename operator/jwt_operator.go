package operator

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GenerateToken(user_id uint) (string, error) {

	token_lifespan := 1

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("44CE8E4543D8203B2105A798AB63E778D1217FF992B279D28BA49EFA91168E55"))

}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("44CE8E4543D8203B2105A798AB63E778D1217FF992B279D28BA49EFA91168E55"), nil
	})
	if err != nil {
		log.Println("TokenValid",err)
		return err
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		log.Println(token)
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		log.Println(strings.Split(bearerToken, " ")[1])
		return strings.Split(bearerToken, " ")[1]
	}
	log.Println("token is empty")
	return ""
}

func ExtractTokenID(c *gin.Context) (uint, error) {

	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("44CE8E4543D8203B2105A798AB63E778D1217FF992B279D28BA49EFA91168E55"), nil
	})
	if err != nil {
		log.Println("extractedTokenId : ",err)
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			log.Println("extractedTokenId : ",err)
			return 0, err
		}
		log.Println("extractedTokenId : ",uid)
		return uint(uid), nil
	}
	log.Println("extractedTokenId : return 0")
	return 0, nil
}
