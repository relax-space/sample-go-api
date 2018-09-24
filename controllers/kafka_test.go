package controllers_test

import (
	"fmt"
	"net/http/httptest"
	"sample-go-api/controllers"
	"testing"

	"github.com/labstack/echo"
	"github.com/relax-space/go-kit/test"
)

//go test .\controllers\init_test.go  .\controllers\kafka_test.go  -v
func Test_kafak(t *testing.T) {
	req := httptest.NewRequest(echo.GET, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	test.Ok(t, handleWithFilter(controllers.KafkaApiController{}.Producer, echoApp.NewContext(req, rec)))
	fmt.Println(string(rec.Body.Bytes()))
	fmt.Printf("http status:%v", rec.Result().StatusCode)
}
