package auth

import (
	"github.com/dgrijalva/jwt-go"
)

// 自定义载荷部分（对外提供的对象）
type CustomerClaim struct {
	// 登录用户id
	UserId int32 `json:"user_id"`
	// 用户名称
	UserName string `json:"user_name"`
	// 商户id,商户编号(平台端为0)
	MerchantId int32 `json:"merchant_id"`
	// 对应的企业微信的企业id(如果为空,标识小程序|外部登录)
	Corpid string `json:"corpid" `
	// token类型, 枚举 token_type
	TokenType int32 `json:"token_type"`
	// 来源,枚举source_type
	Source int32 `json:"source"`
	// 是否超级管理员, 1表示超级管理员
	SuperAdmin int32 `json:"super_admin"`
}

// access-token载荷;
type TokenClaim struct {
	CustomerClaim
	jwt.StandardClaims
}

// access-token载荷;
type RefreshTokenClaim struct {
	// 登录用户id
	UserId int32 `json:"user_id"`
	// 商户id
	MerchantId int32 `json:"merchant_id"`
	jwt.StandardClaims
}

// 内部的授权token
type InnerAccessToken struct {
	CustomerClaim
	// 绝对的过期时间(使用时间戳)
	Expire int64 `json:"expire"`
}
