package status

import (
	"net/http"
	"testing"

	"github.com/appleboy/gofight/v2"
	"github.com/buger/jsonparser"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func httpRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	r.GET("/api/status", StatusHandler)

	return r
}

func TestStatusHandler(t *testing.T) {

	r := gofight.New()

	r.GET("/api/status").
		Run(httpRouter(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			data := []byte(r.Body.String())

			value, _ := jsonparser.GetString(data, "go_version")

			assert.NotEmpty(t, value)
			assert.Equal(t, http.StatusOK, r.Code)
		})
}
