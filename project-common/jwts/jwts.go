package jwts

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type AuthClaim struct {
	UID int64 `json:"uid"`
	jwt.StandardClaims
}

var Secret = "私钥"
var hmacSampleSecret = []byte(Secret)

const TokenExpireDuration = 24 * time.Hour //过期时间

func CreateToken(uid int64) (tokenStr string) {
	var authClaim AuthClaim
	authClaim.UID = uid
	authClaim.StandardClaims.ExpiresAt = time.Now().Add(TokenExpireDuration).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, authClaim)
	tokenString, _ := token.SignedString(hmacSampleSecret) //私钥加密
	return tokenString
}

func Parse(tokenString string) (auth AuthClaim, Valid bool) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})
	Valid = token.Valid //token是否有效 true有效  false无效
	if claims, ok := token.Claims.(jwt.MapClaims); ok && Valid {
		auth.UID = int64(claims["uid"].(float64))       //自定义的UID
		auth.ExpiresAt = int64(claims["exp"].(float64)) //过期时间
	}
	return
}
