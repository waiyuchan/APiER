package custom_jwt

import (
	"apier/internal/global/custom_errors"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 使用工厂创建一个 JWT 结构体
func CreateJWT(signKey string) *JwtSign {
	if len(signKey) <= 0 {
		signKey = "apierSign"
	}
	return &JwtSign{[]byte(signKey)}
}

// 定义一个 JWT验签 结构体
type JwtSign struct {
	SigningKey []byte
}

// CreateToken 生成一个token
func (js *JwtSign) CreateToken(payload Payload) (string, error) {
	// 生成jwt格式的 header、claims 部分
	tokenPartA := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	// 继续添加秘钥值，生成最后一部分
	return tokenPartA.SignedString(js.SigningKey)
}

// 解析Token
func (js *JwtSign) ParseToken(tokenString string) (*Payload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Payload{}, func(token *jwt.Token) (interface{}, error) {
		return js.SigningKey, nil
	})
	if token == nil {
		return nil, errors.New(custom_errors.ErrorsTokenInvalid)
	}
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New(custom_errors.ErrorsTokenMalFormed)
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New(custom_errors.ErrorsTokenNotActiveYet)
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// 如果 TokenExpired ,只是过期（格式都正确），我们认为他是有效的，接下可以允许刷新操作
				token.Valid = true
				goto labelHere
			} else {
				return nil, errors.New(custom_errors.ErrorsTokenInvalid)
			}
		}
	}
labelHere:
	if claims, ok := token.Claims.(*Payload); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New(custom_errors.ErrorsTokenInvalid)
	}
}

// 更新token
func (j *JwtSign) RefreshToken(tokenString string, extraAddSeconds int64) (string, error) {

	if CustomClaims, err := j.ParseToken(tokenString); err == nil {
		CustomClaims.ExpiresAt = time.Now().Unix() + extraAddSeconds
		return j.CreateToken(*CustomClaims)
	} else {
		return "", err
	}
}
