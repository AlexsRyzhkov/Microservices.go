package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

var jwtSecret = []byte("oP80AQ")

func generateToken(id uint, d time.Duration) (string, error) {

	claims := jwt.MapClaims{
		"user_id": id,
		"exp":     time.Now().Add(d).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenAccessToken(id uint) (string, error) {
	return generateToken(id, 30*time.Minute)
}

func GenRefreshToken(id uint) (string, error) {
	return generateToken(id, 30*24*time.Hour)
}

func GenTokens(id uint) (string, string, error) {
	accToken, err := GenAccessToken(id)
	if err != nil {
		return "", "", status.Errorf(codes.Aborted, "Ошибка создания пользователя")
	}

	refreshToken, err := GenRefreshToken(id)
	if err != nil {
		return "", "", status.Errorf(codes.Aborted, "Ошибка создания пользователя")
	}

	return accToken, refreshToken, nil
}

func Parse(tokenString string) (uint64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			expiration := time.Unix(int64(exp), 0)
			fmt.Printf("Token expires at: %v\n", expiration)
		}

		fmt.Println(claims)

		return uint64(claims["user_id"].(float64)), nil
	}

	return 0, fmt.Errorf("Token is invalid")

}
