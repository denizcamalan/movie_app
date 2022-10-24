package operator

import (
	"fmt"

	"strconv"
	"strings"
	"time"

	log "github.com/siruspen/logrus"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JWTManager interface{
	TokenValid(c *gin.Context) error 
	ExtractToken(c *gin.Context) string
	ExtractTokenID(c *gin.Context) (uint, error)
}

type JWTModel struct{
}

func NewJWTModel() (*JWTModel){
	var models JWTModel
	return &models
}

func (j *JWTModel) TokenValid(c *gin.Context) error {
	tokenString := j.ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("44CE8E4543D8203B2105A798AB63E778D1217FF992B279D28BA49EFA91168E55"), nil
	})
	if err != nil {
		log.Error("TokenValid",err)
		return err
	}
	return nil
}

func (*JWTModel) ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		log.Error(token)
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		log.Info(strings.Split(bearerToken, " ")[1])
		return strings.Split(bearerToken, " ")[1]
	}
	log.Info("token is empty")
	return ""
}

func (j *JWTModel) ExtractTokenID(c *gin.Context) (uint, error) {

	tokenString := j.ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("44CE8E4543D8203B2105A798AB63E778D1217FF992B279D28BA49EFA91168E55"), nil
	})
	if err != nil {
		log.Error("extractedTokenId : ",err)
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			log.Error("extractedTokenId : ",err)
			return 0, err
		}
		log.Infof("extractedTokenId : %d",uid)
		return uint(uid), nil
	}
	log.Info("extractedTokenId : return 0")
	return 0, nil
}

func generateToken(user_id uint) (string, error) {

	token_lifespan := 1

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("44CE8E4543D8203B2105A798AB63E778D1217FF992B279D28BA49EFA91168E55"))

}
