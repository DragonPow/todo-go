package jwt_handle

import (
	"project1/domain"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	jwt_method = jwt.SigningMethodHS256
	key        = []byte("vungocthach")
)

type any interface{}

type customClaim struct {
	data int32 `json:"data"`
	jwt.StandardClaims
}

func GenerateToken(user_id int32) string {
	claims := customClaim{
		data: user_id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 60, // 2 minutes
		},
	}

	token := jwt.NewWithClaims(jwt_method, claims)

	tokenString, err := token.SignedString(key)
	if err != nil {
		panic(err)
	}

	return tokenString
}

func ParseToData(tokenString string) (data any, err error) {
	at(time.Unix(0, 0), func() {
		jwt_token, jwt_err := jwt.ParseWithClaims(tokenString, &customClaim{}, func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})

		claims, ok := jwt_token.Claims.(*customClaim)

		if ok && jwt_token.Valid {
			data = claims.data
			err = jwt_err
			return
		} else {
			err = domain.ErrTokenInvalid
			data = nil
		}
	})

	return data, err
}

// Override time value for tests.  Restore default value after.
func at(t time.Time, f func()) {
	jwt.TimeFunc = func() time.Time {
		return t
	}
	f()
	jwt.TimeFunc = time.Now
}
