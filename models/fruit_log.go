package models

import (
	"time"

	"github.com/go-xorm/xorm"
)

type FruitLog struct {
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
	Success   bool       `json:"success"`
}

func (d *FruitLog) Create(db *xorm.Engine) (affectedRow int64, err error) {
	affectedRow, err = db.MustCols("success").Insert(d)
	return
}

func (FruitLog) CreateBatch(db *xorm.Engine, fruitLogs []*FruitLog) (affectedRow int64, err error) {
	affectedRow, err = db.MustCols("success").Insert(&fruitLogs)
	return
}
