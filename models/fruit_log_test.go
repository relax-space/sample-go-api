package models_test

import (
	"fmt"
	"sample-go-api/models"
	"testing"
)

func TestFruitLogCreate(t *testing.T) {
	f := &models.FruitLog{
		Code: "123",
	}
	affectedRow, err := f.Create(db)
	fmt.Println(affectedRow, err, f)
}

func TestFruitLogBatch(t *testing.T) {
	f := []*models.FruitLog{
		&models.FruitLog{
			Code: "456",
		},
	}
	affectedRow, err := (models.FruitLog{}).CreateBatch(db, f)
	fmt.Println(affectedRow, err, f)
}
