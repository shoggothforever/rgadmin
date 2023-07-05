package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Leavereq struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Status    *bool              `bson:"status" json:"status"`
	StaffCode string             `bson:"staffcode" json:"stuffCode"`
	Subject   string             `bson:"subject" json:"subject"`
	Reason    string             `bson:"reason" json:"reason"`
	Name      string             `bson:"name" json:"name"`
	Date      string             `bson:"date" json:"date"`
}
