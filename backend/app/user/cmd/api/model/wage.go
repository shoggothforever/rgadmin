package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var basewagemap map[string]float64

const (
	ROLE1   string  = "admissions staff"
	ROLE2   string  = "assistant"
	ROLE3   string  = "cleaner"
	ROLE4   string  = "computer administrator"
	ROLE5   string  = "driver"
	ROLE6   string  = "filing clerk"
	ROLE7   string  = "hr"
	ROLE8   string  = "librarian"
	ROLE9   string  = "restaurant staff"
	ROLE10  string  = "security bureau"
	ROLE11  string  = "teacher"
	ROLE12  string  = "vice-principal"
	ElseFee float64 = 1000
)

type Wage struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UpdateAt      time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt      time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
	StaffCode     string             `bson:"staffcode" json:"stuffCode"`
	Name          string             `bson:"name" json:"name"`
	Year          int                `bson:"year" json:"year"`
	Month         int                `bson:"month" json:"month"`
	WorkTime      float64            `bson:"worktime" json:"workTime"`
	WageBeforeTax float64            `bson:"wagebeforetax" json:"wageBeforeTax"`
	Tax           float64            `bson:"tax" json:"tax" `
	ActualWage    float64            `bson:"actualwage" json:"actualWage"`
}

func init() {
	basewagemap = make(map[string]float64)
	basewagemap[ROLE1] = 10000
	basewagemap[ROLE2] = 9000
	basewagemap[ROLE3] = 6000
	basewagemap[ROLE4] = 7000
	basewagemap[ROLE5] = 6000
	basewagemap[ROLE6] = 10000
	basewagemap[ROLE7] = 9000
	basewagemap[ROLE8] = 8000
	basewagemap[ROLE9] = 6000
	basewagemap[ROLE10] = 6000
	basewagemap[ROLE11] = 9000
	basewagemap[ROLE12] = 15000
}
func GetBasewage(role string) float64 {
	return basewagemap[role]
}
func CalTax(base float64) float64 {
	var ans float64
	switch {
	case base <= 5000:
		ans = 0
		break
	case base <= 8000:
		ans = (base - 5000) * 0.03
	case base <= 17000:
		ans = 90 + (base-8000)*0.10
		break
	case base <= 30000:
		ans = 990 + (base-17000)*0.20
		break
	case base <= 40000:
		ans = 3590 + (base-30000)*0.25
		break
	case base <= 70000:
		ans = 6090 + (base-40000)*0.30
		break
	case base <= 85000:
		ans = 15090 + (base-70000)*0.35
	default:
		ans = 20340 + (base-85000)*0.45
	}
	return ans
}
