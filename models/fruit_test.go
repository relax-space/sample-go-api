package models_test

import (
	"fmt"
	"sample-go-api/models"
	"testing"

	"github.com/relax-space/go-kit/test"
)

func TestFruitCreate(t *testing.T) {
	f := &models.Fruit{
		Code: "123",
	}
	affectedRow, err := f.Create(ctx)
	fmt.Println(affectedRow, err, f)
}
func TestStoreCreate(t *testing.T) {
	f := &models.Store{
		Id:   3,
		Code: "4",
	}
	affectedRow, err := f.Create(ctx)
	fmt.Println(affectedRow, err, f)
}

func TestFruitUpdate(t *testing.T) {
	f := &models.Fruit{
		Code: "123",
	}
	affectedRow, err := f.Update(ctx, 71)
	fmt.Println(affectedRow, err)
}

func TestFruitDelete(t *testing.T) {
	affectedRow, err := models.Fruit{}.Delete(ctx, 2)
	fmt.Println(affectedRow, err)
}

func TestFruitGetAll(t *testing.T) {
	total, items, err := models.Fruit{}.GetAll(ctx, nil, nil, 0, 2, nil)
	fmt.Println(total, items, err)
}
func TestFruitGetById(t *testing.T) {
	has, v, err := models.Fruit{}.GetById(ctx, 1)
	fmt.Println(has, v, err)
}
func TestChanSession_norm(t *testing.T) {
	for index := 0; index < 100; index++ {
		err := models.Fruit{}.ChanSession(ctx, &models.Fruit{
			Id:   71,
			Code: "127",
		}, &models.Store{
			Code: "6",
		})
		test.Ok(t, err)
		fmt.Println(err)
	}

}
func TestChanSession_transaction(t *testing.T) {
	err := models.Fruit{}.ChanSession(ctx, &models.Fruit{
		Id:   71,
		Code: "128",
	}, &models.Store{
		Id:   8,
		Code: "7",
	})
	fmt.Println(err)
	test.Ok(t, err)
}
