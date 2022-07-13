package jwt_handle

import (
	"project1/domain"

	"github.com/golang-jwt/jwt/v4"
)

var (
	jwt_method = jwt.SigningMethodES256
)

type customClaim struct {
	data interface{} `json:"data"`
	jwt.StandardClaims
}

func GenerateToken(data interface{}) string {
	claims := customClaim{
		data,
		jwt.StandardClaims{
			ExpiresAt: 60 * 2, // 2 minutes
		},
	}

	token := jwt.NewWithClaims(jwt_method, claims)

	tokenString, err := token.SignedString(token)
	if err != nil {
		panic(err)
	}

	return tokenString
}

func ParseToData(token string) (interface{}, error) {
	jwt_token, err := jwt.ParseWithClaims(token, &customClaim{}, func(token *jwt.Token) (interface{}, error) {
		return nil, nil
	})

	claims, ok := jwt_token.Claims.(*customClaim)

	if ok && jwt_token.Valid {
		return claims.data, err
	} else {
		return nil, domain.ErrTokenInvalid
	}
}
