package sample

import (
	"fmt"
	"testing"
)

func Test_ChangeLive(t *testing.T) {
	value := ChangeLive()
	fmt.Printf("result:%v\n", value)
}
