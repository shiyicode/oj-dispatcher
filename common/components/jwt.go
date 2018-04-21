package components

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/open-fightcoder/oj-dispatcher/common/g"
)

var mySigningKey = "let's to change the world!"

func CreateToken(userId int64) (string, error) {
	cfg := g.Conf()
	claims := jwt.MapClaims{
		"uid": userId,
		"exp": time.Now().Add(time.Second * 24 * 60 * 60 * time.Duration(cfg.Jwt.MaxEffectiveTime)).Unix(),
	}

	var signingMethod *jwt.SigningMethodHMAC
	switch cfg.Jwt.EncodeMethod {
	case "HS256":
		signingMethod = jwt.SigningMethodHS256
	case "HS384":
		signingMethod = jwt.SigningMethodHS384
	case "HS512":
		signingMethod = jwt.SigningMethodHS512
	default:
		signingMethod = jwt.SigningMethodHS256
	}

	token := jwt.NewWithClaims(signingMethod, claims)

	tokenStr, err := token.SignedString([]byte(mySigningKey))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func RequireTokenAuthentication(tokenStr string) (bool, string) {
	token, _ := jwt.ParseWithClaims(tokenStr, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error")
		}
		return mySigningKey, nil
	})
	userId, _ := token.Claims.(jwt.MapClaims)["uid"].(string)
	return token.Valid, userId
}
