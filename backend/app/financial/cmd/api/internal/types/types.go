// Code generated by goctl. DO NOT EDIT.
package types
import (
	"crypto/md5"
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"time"
)
type Response struct {
	Code uint32
	Msg  string
}

type CalReq struct {
	Id       string  `json:"id"`
	Jwt      string  `json:"jwt"`
	WorkHour float32 `json:"workHour"`
}

type CalResp struct {
	Response
	WageBeforeTax float32 `json:"wageBeforeTax"`
	WageAfterTax  float32 `json:"wageAfterTax"`
	ActualWage    float32 `json:"actualWage"`
}

type CodeError struct {
	Code uint32
	Msg  string
}

// 返回给前端的错误码
func (e *CodeError) GetErrCode() uint32 {
	return e.Code
}

// 返回给前端显示端错误信息
func (e *CodeError) GetErrMsg() string {
	return e.Msg
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.Code, e.Msg)
}
func NewResponse(code uint32, msg string) Response {
	return Response{code, msg}
}
func NewErrCodeMsg(errCode uint32, errMsg string) *CodeError {
	return &CodeError{Code: errCode, Msg: errMsg}
}

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
