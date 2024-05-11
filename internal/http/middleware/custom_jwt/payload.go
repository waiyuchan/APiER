package custom_jwt

import "github.com/dgrijalva/jwt-go"

type Payload struct {
	UserId   int64  `json: "user_id"`
	UserName string `json: "username"`
	jwt.StandardClaims
}
