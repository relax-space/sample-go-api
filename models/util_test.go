package models_test

import (
	"fmt"
	"sample-go-api/models"
	"testing"
)

func TestRandom(t *testing.T) {
	for index := 0; index < 100; index++ {
		f := models.UuIdInt64()
		fmt.Println(f)
	}

}
