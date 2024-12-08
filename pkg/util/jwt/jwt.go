package jwt

import (
	"errors"

	"time"

	"github.com/LucienLSA/go-gin-mall/consts"
	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("LucienLSA")

type Claims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// 签发用户的token
func GenerateToken(id uint, username string) (accesToken, refreshToken string, err error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(consts.AccessTokenExpireDuration)    // 30天过期
	rtExpireTime := nowTime.Add(consts.RefreshTokenExpireDuration) // 30天过期

	claims := Claims{
		ID:       id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-mall",
		},
	}
	// 加密并获得完整的编码后的字符串token
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: rtExpireTime.Unix(),
		Issuer:    "gin-mall",
	}).SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

// ParseToken 解析token 验证用户token
func ParseToken(tokenString string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// RefreshToken 刷新token 验证用户yoken
func ParseRefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	accessClaim, err := ParseToken(aToken)
	if err != nil {
		return
	}
	refreshClaim, err := ParseToken(rToken)
	if err != nil {
		return
	}
	if accessClaim.ExpiresAt > time.Now().Unix() {
		// 如果 access token 未过期，每一次请求都刷新 refresh token 和 access token
		return GenerateToken(accessClaim.ID, accessClaim.Username)
	}
	if refreshClaim.ExpiresAt > time.Now().Unix() {
		// 如果 access token 已过期，但 refresh token 未过期，都刷新 refresh token 和 access token
		return GenerateToken(accessClaim.ID, accessClaim.Username)
	}
	// 如果都过期
	return "", "", errors.New("身份过期，重新登录")
}

// EmailClaims
type EmailClaims struct {
	UserID        uint   `json:"user_id"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	OperationType uint   `json:"operation_type"`
	jwt.StandardClaims
}

// 生成邮箱验证token
func GenerateEmailToken(userID, Operation uint, email, password string) (string, error) {
	nowTime := time.Now()
	// 十五分钟有效
	expireTime := nowTime.Add(15 * time.Minute)
	claims := EmailClaims{
		UserID:        userID,
		Email:         email,
		Password:      password,
		OperationType: Operation,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-mall",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// 验证邮箱验证token
func ParseEmailToken(token string) (*EmailClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &EmailClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*EmailClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
