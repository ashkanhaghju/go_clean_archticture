package auth

import (
	"crypto/subtle"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type TokenType string

const (
	AccessToken  TokenType = "access_token"
	RefreshToken TokenType = "refresh_token"
)

type JWTAuth struct {
	SecretKey                      []byte
	AccessTokenExpirationDuration  time.Duration
	RefreshTokenExpirationDuration time.Duration
	Audience                       string
	Issuer                         string
}

func NewJWTAuth(secretKey, audience, issuer string, accessTokenDuration, refreshTokenDuration time.Duration) JWTAuth {
	return JWTAuth{
		SecretKey:                      []byte(secretKey),
		AccessTokenExpirationDuration:  accessTokenDuration,
		RefreshTokenExpirationDuration: refreshTokenDuration,
		Audience:                       audience,
		Issuer:                         issuer,
	}
}

type JWTAccessTokenPayload struct {
	Id   uint64   `json:"id,omitempty"`
	Name string   `json:"name,omitempty"`
	Role []string `json:"role,omitempty"`
}

type JWTRefreshTokenPayload struct {
	Id uint64 `json:"id,omitempty"`
}

type JwtClaim struct {
	TokenType TokenType   `json:"token_type,omitempty"`
	Payload   interface{} `json:"payload,omitempty"`
	Audience  string      `json:"aud,omitempty"`
	ExpiresAt int64       `json:"exp,omitempty"`
	IssuedAt  int64       `json:"iat,omitempty"`
	Issuer    string      `json:"iss,omitempty"`
}

type JwtResponse struct {
	AccessToken  string
	RefreshToken string
}

func (a JWTAuth) GenerateJWTToken(accessData JWTAccessTokenPayload) (*JwtResponse, error) {
	accessClaims := JwtClaim{
		TokenType: AccessToken,
		Payload:   accessData,
		Audience:  a.Audience,
		ExpiresAt: time.Now().Add(a.AccessTokenExpirationDuration).Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    a.Issuer,
	}

	refreshClaims := JwtClaim{
		TokenType: RefreshToken,
		Payload:   JWTRefreshTokenPayload{Id: accessData.Id},
		Audience:  a.Audience,
		ExpiresAt: time.Now().Add(a.RefreshTokenExpirationDuration).Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    a.Issuer,
	}

	aToken := jwt.NewWithClaims(jwt.SigningMethodHS512, accessClaims)
	rToken := jwt.NewWithClaims(jwt.SigningMethodHS512, refreshClaims)

	// Generate encoded token and send it as response.
	accessToken, err := aToken.SignedString(a.SecretKey)
	if err != nil {
		return nil, err
	}

	refreshToken, err := rToken.SignedString(a.SecretKey)
	if err != nil {
		return nil, err
	}

	return &JwtResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a JWTAuth) ParseToken(authToken string) (*JwtClaim, error) {
	keyFunc := func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != "HS512" {
			return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
		}
		return a.SecretKey, nil
	}
	claims := &JwtClaim{}
	token, err := jwt.ParseWithClaims(authToken, claims, keyFunc)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return claims, nil
}

func (c JwtClaim) Valid() error {
	vErr := new(jwt.ValidationError)
	now := jwt.TimeFunc().Unix()

	// The claims below are optional, by default, so if they are set to the
	// default value in Go, let's not fail the verification for them.
	if !c.VerifyExpiresAt(now, false) {
		delta := time.Unix(now, 0).Sub(time.Unix(c.ExpiresAt, 0))
		vErr.Inner = fmt.Errorf("token is expired by %v", delta)
		vErr.Errors |= jwt.ValidationErrorExpired
	}

	if !c.VerifyIssuedAt(now, false) {
		vErr.Inner = fmt.Errorf("Token used before issued")
		vErr.Errors |= jwt.ValidationErrorIssuedAt
	}

	if vErr.Errors == 0 {
		return nil
	}

	return vErr
}

// Compares the aud claim against cmp.
// If required is false, this method will return true if the value matches or is unset
func (c *JwtClaim) VerifyAudience(cmp string, req bool) bool {
	return verifyAud([]string{c.Audience}, cmp, req)
}

// Compares the exp claim against cmp.
// If required is false, this method will return true if the value matches or is unset
func (c *JwtClaim) VerifyExpiresAt(cmp int64, req bool) bool {
	return verifyExp(c.ExpiresAt, cmp, req)
}

// Compares the iat claim against cmp.
// If required is false, this method will return true if the value matches or is unset
func (c *JwtClaim) VerifyIssuedAt(cmp int64, req bool) bool {
	return verifyIat(c.IssuedAt, cmp, req)
}

// Compares the iss claim against cmp.
// If required is false, this method will return true if the value matches or is unset
func (c *JwtClaim) VerifyIssuer(cmp string, req bool) bool {
	return verifyIss(c.Issuer, cmp, req)
}

func verifyAud(aud []string, cmp string, required bool) bool {
	if len(aud) == 0 {
		return !required
	}
	// use a var here to keep constant time compare when looping over a number of claims
	result := false

	var stringClaims string
	for _, a := range aud {
		if subtle.ConstantTimeCompare([]byte(a), []byte(cmp)) != 0 {
			result = true
		}
		stringClaims = stringClaims + a
	}

	// case where "" is sent in one or many aud claims
	if len(stringClaims) == 0 {
		return !required
	}

	return result
}

func verifyExp(exp int64, now int64, required bool) bool {
	if exp == 0 {
		return !required
	}
	return now <= exp
}

func verifyIat(iat int64, now int64, required bool) bool {
	if iat == 0 {
		return !required
	}
	return now >= iat
}

func verifyIss(iss string, cmp string, required bool) bool {
	if iss == "" {
		return !required
	}
	if subtle.ConstantTimeCompare([]byte(iss), []byte(cmp)) != 0 {
		return true
	} else {
		return false
	}
}
