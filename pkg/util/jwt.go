package util

import (
	"github.com/HashCell/golang/cloudgo/pkg/setting"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

// 自定义的密钥
var jwtSecret = []byte(setting.JwtSecret)

// 继承standardClaims，另外自定义2个字段
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// 创建token
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	// 过期时间为３小时
	expireTIme := nowTime.Add(3 * time.Hour)

	// payload　jwt数据负载
	claims := Claims{
		Username:username,
		Password:password,
		StandardClaims:jwt.StandardClaims{
			ExpiresAt:expireTIme.Unix(),
			Issuer:"hash-cell",
		},
	}

	// 添加header,指定签名算法SigningMethodHS256，返回一个 {header, claim｝JSON
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用密钥对header.payload进行签名
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	//　从token字符串解出　Token对象，包含Claims对象
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		// 类型断言，从中获取到Claims对象
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}