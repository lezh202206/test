package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"test/t_viper"
	"time"
)

var (
	errParseToken        = errors.New("无法解析token数据")
	errTokenType         = errors.New("未指定tokenType参数")
	errAccessTokenExpire = errors.New("access-token已过期")
	errRefreshToken      = errors.New("refresh-token无效")
)

var (
// 本地会话缓存，防止重复解析jwt信息;
// sessionCache = cache.New()
)

// 返回jwt-token
func GenJwtToken(claim CustomerClaim) (string, error) {
	if claim.TokenType == 0 {
		return "", errTokenType
	}
	var token TokenClaim
	token.StandardClaims = jwt.StandardClaims{
		Audience:  "",                // 受众群体
		Id:        "",                // 编号
		IssuedAt:  time.Now().Unix(), // 签发时间
		Issuer:    "zby",             // 签发人
		NotBefore: time.Now().Unix(), // 生效时间
		Subject:   "zby",             // 主题
	}
	token.ExpiresAt = time.Now().Add(t_viper.JWTViper().ExpireMin).Unix()

	//自定义的字段
	token.CustomerClaim = claim
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, token)
	return tokenClaims.SignedString(t_viper.JWTViper().SecretBytes)
}

// 解码token，获取载荷字典(claims)
func ParseJwtToken(token string) (jwt.MapClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(t_viper.JWTViper().SecretBytes), nil
	})
	if err != nil {
		if err.Error() == "Token is expired" {
			return nil, errAccessTokenExpire
		}
		return nil, err
	}
	if jwtToken != nil && jwtToken.Valid {
		if claim, ok := jwtToken.Claims.(jwt.MapClaims); ok {
			return claim, nil
		}
	}
	return nil, errParseToken
}

// 获取自定义的载荷信息;
func getCustomerClaim(token string) (*CustomerClaim, error) {
	//if session, ok := sessionCache.Get(token); ok {
	//	return session.(*CustomerClaim), nil
	//}
	claims, err := ParseJwtToken(token)
	if err != nil {
		return nil, err
	}
	customer := &CustomerClaim{
		UserId:     int32(claims["user_id"].(float64)),
		MerchantId: int32(claims["merchant_id"].(float64)),
		TokenType:  int32(claims["token_type"].(float64)),
		Source:     int32(claims["source"].(float64)),
	}
	if v, ok := claims["corpid"]; ok {
		customer.Corpid = v.(string)
	}
	if v, ok := claims["user_name"]; ok {
		customer.UserName = v.(string)
	}

	// 设置缓存
	//sessionCache.Set(token, customer, t_viper.JWTViper().ExpireMin)
	return customer, nil
}

// 检测token是否有效
func IsRefreshTokenValidate(token string) error {
	jwtToken, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return t_viper.JWTViper().SecretBytes, nil
	})
	if err != nil {
		return err
	}
	if jwtToken == nil || !jwtToken.Valid {
		return errRefreshToken
	}
	return nil
}
