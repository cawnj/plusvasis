package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func GenerateJWT() (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mySigningKey := []byte(os.Getenv("SAMPLE_KEY"))
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "svelte"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Println("Something Went Wrong")
		return "", err
	}

	return tokenString, nil
}

func main() {
	tokenString, err := GenerateJWT()
	if err != nil {
		fmt.Println("error")
	}

	fmt.Println(tokenString)
}
