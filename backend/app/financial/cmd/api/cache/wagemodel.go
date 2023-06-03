package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
)

var _ WageModel = (*customWageModel)(nil)

type (
	// WageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customWageModel.
	WageModel interface {
		wageModel
	}

	customWageModel struct {
		*defaultWageModel
	}
)

// NewWageModel returns a model for the mongo.
func NewWageModel(url, db, collection string, c cache.CacheConf) WageModel {
	conn := monc.MustNewModel(url, db, collection, c)
	return &customWageModel{
		defaultWageModel: newDefaultWageModel(conn),
	}
}
