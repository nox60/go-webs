package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	fmt.Println("test")
	JwtSign()
}

func JwtSign() {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["sub"] = "kkkkkkkkkkkkkkkkkkkkkkkkkksadfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasfdasfasdfasdfasdfasdfasfasdfasdfasdf"
	token.Claims = claims

	fmt.Println(claims)
	fmt.Println(token.Claims)

	tokenString, _ := token.SignedString([]byte("testkey"))

	fmt.Println(tokenString)

	//if err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	fmt.Fprintln(w, "Error extracting the key")
	//	fatal(err)
	//}

	//tokenString, err := token.SignedString([]byte(SecretKey))
	//if err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	fmt.Fprintln(w, "Error while signing the token")
	//	fatal(err)
	//}
}
