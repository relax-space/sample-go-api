package models

import (
	"context"
	"time"

	"sample-go-api/factory"
)

type FruitConsumer struct {
	Id        int64      `json:"id" xorm:"int64 notnull autoincr pk 'id'"`
	Code      string     `json:"code"`
	Name      string     `json:"name"`
	Color     string     `json:"color"`
	Price     int64      `json:"price"`
	StoreCode string     `json:"storeCode"`
	CreatedAt *time.Time `json:"createdAt" xorm:"created"`
	UpdatedAt *time.Time `json:"updatedAt" xorm:"updated"`
	DeletedAt *time.Time `json:"deletedAt" xorm:"deleted"`
	UniqueId  uint64     `json:"uniqueId"`
}

func (d *FruitConsumer) Create(ctx context.Context) (affectedRow int64, err error) {
	affectedRow, err = factory.DB(ctx).Insert(d)
	return
}
