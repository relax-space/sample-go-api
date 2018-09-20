package sample

import (
	"fmt"
	"testing"
)

func Test_Scholarship(t *testing.T) {
	money, err := Scholarship(140, 120)
	fmt.Println(money, err)
}
