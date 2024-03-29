package crypto

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"io"
	"myserver/pkg/config"
	"time"
)

// GenerateRandomBytes returns securely generated random bytes. It will return
// an error if the system's secure random number generator fails to function
// correctly, in which case the caller should not continue.
func GenerateRandomBytes(n int) []byte {
	b := make([]byte, n)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		panic(err)
	}
	return b
}

// Timestamp returns the current timestamp, in seconds.
func Timestamp() int64 {
	return time.Now().Unix()
}

// Base64Encode encodes a value using base64.
func Base64Encode(value []byte) []byte {
	enc := make([]byte, base64.RawURLEncoding.EncodedLen(len(value)))
	base64.RawURLEncoding.Encode(enc, value)
	return enc
}

// Base64Decode decodes a value using base64.
func Base64Decode(value []byte) ([]byte, error) {
	dec := make([]byte, base64.RawURLEncoding.DecodedLen(len(value)))
	b, err := base64.RawURLEncoding.Decode(dec, value)
	if err != nil {
		return nil, err
	}
	return dec[:b], nil
}

type jwtCustomClaims struct {
	UserID 	  string   `json:"userID"`
	Roles    []string `json:"roles"`
	jwt.StandardClaims
}
const jwtExpiry = 12 * 60 * 60

func CreateToken(userID string, roles []string) (string, error) {
	claims := jwtCustomClaims{
		UserID: userID,
		Roles:    roles,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + jwtExpiry,
			IssuedAt:  time.Now().Unix(),
			Subject:   "Session Token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	b := []byte(config.GetConfig().JWTkey)
	tokenstring, err := token.SignedString(b)
	if err != nil {
		return "", err
	}
	return tokenstring, nil
}
