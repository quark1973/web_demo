package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)
func HashPassword(password string)(string,error) {
	hash,err :=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil {
		return "",err
	}
	return string(hash), nil
}

func GenerateJWT(username string)(string,error){
	token :=jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	signedString,err:=token.SignedString([]byte("secret"))
	if err != nil {
		return "",err
	}
	return "Bearer "+signedString, nil
}

func CheckPassword(password string,hash string)(bool){
	return bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))==nil
}

func ParseJWT(tokenStr string)(string,error){
	if len(tokenStr) > 7 && tokenStr[:7] == "Bearer "{
		tokenStr = tokenStr[7:]
	}
	token,err:=jwt.Parse(tokenStr,func(token *jwt.Token) (interface{}, error) {
		if _,ok:=token.Method.(*jwt.SigningMethodHMAC);!ok{
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return "",err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid{
		username, ok := claims["username"].(string)
		if !ok {
			return "", jwt.ErrTokenExpired
		}
		return username, nil
	}
	return "", jwt.ErrTokenExpired
}