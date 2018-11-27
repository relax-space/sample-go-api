package sample_test

import (
	"sample-go-api/sample"
	"testing"
)

func Test_RightChan(t *testing.T) {
	sample.ChanRight()
}

func Test_WrongChan(t *testing.T) {
	sample.ChanWrong()
}
