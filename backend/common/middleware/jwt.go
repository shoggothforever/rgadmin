package common

import (
	"crypto/md5"
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"time"
)

// three month
const (
	EXPIRE        = 6 * time.Hour
	issuer string = "hduer"
)

type Claims struct {
	ID   string
	Name string
	jwtgo.StandardClaims
}

func NewClaim() Claims {
	return Claims{}
}
func Messagedigest5(s, salt string) string {
	if (s + salt) == "" {
		s = "123456"
	}
	data := md5.Sum([]byte(s + salt))
	return fmt.Sprintf("%x", data)
}
func (c *Claims) GeneraterJwt(id, name, salt string, duration time.Duration) (string, error) {
	now := time.Now().In(time.Local)
	c.ID = id
	c.Name = name
	c.Issuer = issuer + name
	c.ExpiresAt = now.Add(duration).Unix()
	tokenclaim := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, c)
	jwt, err := tokenclaim.SignedString([]byte(salt))
	return jwt, err
}
