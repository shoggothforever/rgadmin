// Code generated by goctl. DO NOT EDIT.
package model

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var prefixWageCacheKey = "cache:wage:"

type wageModel interface {
	Insert(ctx context.Context, data *Wage) error
	FindOne(ctx context.Context, id string) (*Wage, error)
	Update(ctx context.Context, data *Wage) (*mongo.UpdateResult, error)
	Delete(ctx context.Context, id string) (int64, error)
}

type defaultWageModel struct {
	conn *monc.Model
}

func newDefaultWageModel(conn *monc.Model) *defaultWageModel {
	return &defaultWageModel{conn: conn}
}

func (m *defaultWageModel) Insert(ctx context.Context, data *Wage) error {
	if data.ID.IsZero() {
		data.ID = primitive.NewObjectID()
		data.CreateAt = time.Now()
		data.UpdateAt = time.Now()
	}

	key := prefixWageCacheKey + data.ID.Hex()
	_, err := m.conn.InsertOne(ctx, key, data)
	return err
}

func (m *defaultWageModel) FindOne(ctx context.Context, id string) (*Wage, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidObjectId
	}

	var data Wage
	key := prefixWageCacheKey + id
	err = m.conn.FindOne(ctx, key, &data, bson.M{"_id": oid})
	switch err {
	case nil:
		return &data, nil
	case monc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultWageModel) Update(ctx context.Context, data *Wage) (*mongo.UpdateResult, error) {
	data.UpdateAt = time.Now()
	key := prefixWageCacheKey + data.ID.Hex()
	res, err := m.conn.UpdateOne(ctx, key, bson.M{"_id": data.ID}, bson.M{"$set": data})
	return res, err
}

func (m *defaultWageModel) Delete(ctx context.Context, id string) (int64, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, ErrInvalidObjectId
	}
	key := prefixWageCacheKey + id
	res, err := m.conn.DeleteOne(ctx, key, bson.M{"_id": oid})
	return res, err
}