package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alishokri1661s/SMS-Gateway/API/internals/server"
)

func Test_api(t *testing.T) {
	engine := server.Router
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/sms/log", nil)
	if err != nil {
		t.Fatalf("building request: %v", err)
	}
	engine.ServeHTTP(recorder, request)
	if recorder.Code != 200 {
		t.Fatalf("bad status code: %d", recorder.Code)
	}
}
