package main

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	fmt.Println("Extract user: Violating Single Responsibility Principle")
}

func extractUser(header http.Header) string {
	headerRawInfo := header.Get("Authorization")
	parser := &jwt.Parser{}
	token, _, err := parser.ParseUnverified(headerRawInfo, jwt.MapClaims{})
	if err != nil {
		return ""
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return ""
	}

	return claims["username"].(string)
}
