package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"time"
)

// token有效期
const TokenExpireDuration = time.Hour * 24 * 30

type Claims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

// 生成token
func GenToken(ID int, secret string)  (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		ID:             ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer: "vvvstore", //签发人
		},
	})
	return token.SignedString([]byte(secret))
}

// 解析token
func ParseToken(tokenString string, secret string) (*Claims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

