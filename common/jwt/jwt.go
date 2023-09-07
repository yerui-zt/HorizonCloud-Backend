package jwt

import (
	jwtpkg "github.com/golang-jwt/jwt/v5"
	"github.com/spf13/cast"
	"time"
)

type JWT struct {
	Key []byte
}

func NewJWT(key string) *JWT {
	return &JWT{
		Key: []byte(key),
	}
}

type GeneralUserToken struct {
	UID string `json:"uid"`

	jwtpkg.RegisteredClaims
}

func (jwt *JWT) IssueGeneralUserToken(uid int64, expirePeriod int64, issuer string) (string, error) {
	now := time.Now()

	claims := GeneralUserToken{
		UID: cast.ToString(uid),

		RegisteredClaims: jwtpkg.RegisteredClaims{
			NotBefore: jwtpkg.NewNumericDate(now),
			IssuedAt:  jwtpkg.NewNumericDate(now),
			ExpiresAt: jwtpkg.NewNumericDate(now.Add(time.Duration(expirePeriod) * time.Second)),
			Issuer:    issuer,
		},
	}

	token := jwtpkg.NewWithClaims(jwtpkg.SigningMethodHS256, claims)
	return token.SignedString(jwt.Key)
}
