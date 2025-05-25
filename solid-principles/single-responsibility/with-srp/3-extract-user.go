package main

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	fmt.Println("Extract user: With Single Responsibility Principle")
}

func extractUser(header http.Header) string {
	headerRawInfo := extractRawInfo(header)
	claims := extractClaims(headerRawInfo)
	if claims == nil {
		return ""
	}
	return claims["username"].(string)
}

func extractRawInfo(header http.Header) string {
	return header.Get("Authorization")
}

func extractClaims(headerRawInfo string) jwt.MapClaims {

	parser := &jwt.Parser{}
	token, _, err := parser.ParseUnverified(headerRawInfo, jwt.MapClaims{})
	if err != nil {
		return nil
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil
	}

	return claims
}
