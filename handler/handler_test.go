package handler_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dellykaos/goworld/handler"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	assert.NotNil(handler.New())
}

func TestHello(t *testing.T) {
	assert := assert.New(t)
	h := handler.New()

	testCases := []struct {
		name       string
		url        string
		params     httprouter.Params
		wantResp   string
		wantStatus int
	}{
		{
			name:       "without parameters",
			url:        "/hello",
			params:     nil,
			wantResp:   "Hello, world!",
			wantStatus: 200,
		},
		{
			name:       "with parameters name",
			url:        "/hello/budi",
			params:     httprouter.Params{httprouter.Param{Key: "name", Value: "budi"}},
			wantResp:   "Hello, budi!",
			wantStatus: 200,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp := httptest.NewRecorder()
			req := httptest.NewRequest("GET", tc.url, nil)
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				h.Hello(w, r, tc.params)
			})
			handler.ServeHTTP(resp, req)

			body, _ := ioutil.ReadAll(resp.Body)
			assert.Equal(tc.wantStatus, resp.Code)
			assert.Equal(tc.wantResp, string(body))
		})
	}
}
