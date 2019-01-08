package sample_test

import (
	"sample-go-api/sample"
	"fmt"
	"testing"
)

func Test_ChangeLive(t *testing.T) {
	value := sample.ChangeLive()
	fmt.Printf("result:%v\n", value)
}
