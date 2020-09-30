package ping

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/ping", nil)
	rec := httptest.NewRecorder()

	Handler(rec, req, nil)

	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode, "should return ok")
	assert.Equal(t, "application/json", res.Header.Get("Content-Type"), "should return json content")

	var got pingResponse
	json.NewDecoder(res.Body).Decode(&got)
	assert.Equal(t, "pong", got.Message, "should return pong in message")
}
