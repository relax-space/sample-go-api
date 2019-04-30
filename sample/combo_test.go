package sample_test

import (
	"bytes"
	"fmt"
	"log"
	"sample-go-api/sample"
	"testing"
)

func TestFather(t *testing.T) {
	son := sample.InitSon()
	fmt.Printf("son name:%v \n", son.Name)
	fmt.Printf("father name:%v\n", son.Father.Name)
	fmt.Printf("father sex:%v\n", son.Sex)
}

// TestLog go test ./combo_test.go -v
func TestLog(t *testing.T) {
	sample.ErrLog.Print("show in console")

	const expected = "prefix: test\n"
	buffer := bytes.NewBuffer(make([]byte, 0, 64))
	logger := log.New(buffer, "prefix: ", 0)

	// print
	sample.SetLogger(logger)
	sample.ErrLog.Print("test") //not  show in console

	// check result
	if actual := buffer.String(); actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestLog2(t *testing.T) {

	// print
	sample.SetLogger(sample.MyLogger{})
	sample.ErrLog.Print("test") //not  show in console
}
