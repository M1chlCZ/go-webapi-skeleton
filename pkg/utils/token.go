package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
	"time"
)

type JWT struct {
	privateKey []byte
	publicKey  []byte
}

var JWTTOKEN JWT

func NewJWT() {
	pubKey, err := os.ReadFile(".cert/id_rsa.pub")
	if err != nil {
		log.Fatalln(err)
	}
	prvKey, err := os.ReadFile(".cert/id_rsa")
	if err != nil {
		log.Fatalln(err)
	}
	JWTTOKEN = JWT{
		privateKey: prvKey,
		publicKey:  pubKey,
	}
}

func CreateKeyToken(userid uint64) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(JWTTOKEN.privateKey)
	if err != nil {
		return "", fmt.Errorf("create: parse key: %w", err)
	}

	claims := make(jwt.MapClaims)
	claims["magic"] = GenerateSecureToken(32)
	claims["authorized"] = true
	claims["idUser"] = userid
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	return token, nil
}

func CreateMasternodeKey(url string) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(JWTTOKEN.privateKey)
	if err != nil {
		return "", fmt.Errorf("create: parse key: %w", err)
	}

	claims := make(jwt.MapClaims)
	claims["magic"] = GenerateSecureToken(32)
	claims["authorized"] = true
	claims["idUser"] = url
	claims["exp"] = time.Now().Add(time.Hour * 24 * 365).Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	return token, nil
}

func ValidateKeyToken(token string) (int, interface{}, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(JWTTOKEN.publicKey)
	if err != nil {
		return 0, "", fmt.Errorf("validate: parse key: %w", err)
	}

	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return key, nil
	})
	if err != nil {
		return 0, "", fmt.Errorf("validate: %w", err)
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return 0, "", fmt.Errorf("validate: invalid")
	}

	return int(claims["idUser"].(float64)), claims["magic"], nil
}
