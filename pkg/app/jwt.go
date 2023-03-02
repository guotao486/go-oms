package app

import (
	"oms/global"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// jwt 结构体
//
type Claims struct {
	AppKey    string `json:"app_key`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims
}

// GetJWTsecret
//
/**
 * @description: 获取配置中的Secret
 * @return {*}
 */
func GetJwtSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

// GetJwtExpire
//
/**
 * @description: 获取配置中的失效时间
 * @return {*}
 */
func GetJwtExpire() time.Duration {
	return global.JWTSetting.Expire
}

func GetJwtIssuer() string {
	return global.JWTSetting.Issuer
}

// GenerateToken
//
// jwt.NewWithClaims：根据 Claims 结构体创建 Token 实例，它一共包含两个形参，第一个参数是 SigningMethod，其包含 SigningMethodHS256、SigningMethodHS384、SigningMethodHS512 三种 crypto.Hash 加密算法的方案。第二个参数是 Claims，主要是用于传递用户所预定义的一些权利要求，便于后续的加密、校验等行为。
//
// tokenClaims.SignedString：生成签名字符串，根据所传入 Secret 不同，进行签名并返回标准的 Token。
//
/**
 * @description: 生成 jwt token
 * @param {*} appKey
 * @param {string} appSecret
 * @return {*}
 */
func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(GetJwtExpire())
	Claims := Claims{
		AppKey:    appKey,
		AppSecret: appSecret,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 过期时间
			Issuer:    GetJwtIssuer(),    // 签发者
		},
	}

	// 根据 Claims 结构体创建 Token 实例, 便于后续的加密和校验
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	// 根据 Secret 生成签名字符串
	token, err := tokenClaims.SignedString(GetJwtSecret())
	return token, err
}

// ParseToken
//
// ParseWithClaims：用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回 *Token。
//
// Valid：验证基于时间的声明，例如：过期时间（ExpiresAt）、签发者（Issuer）、生效时间（Not Before），需要注意的是，如果没有任何声明在令牌中，仍然会被认为是有效的。
/**
 * @description: 解析和校验token
 * @param {string} token
 * @return {*}
 */
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return GetJwtSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
