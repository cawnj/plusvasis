package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type APIKey struct {
	Secret string `json:"secret"`
}

func GenerateJWT(key APIKey) (string, error) {
	if key.Secret == string(os.Getenv("SAMPLE_KEY")) {
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
	} else {
		fmt.Println("Invalid API Key")
		return "Invalid API Key", nil
	}
}
