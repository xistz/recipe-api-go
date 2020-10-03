package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingHandler(t *testing.T) {
	db, _ := initDBMock()

	store := NewMySQLStore(db)

	t.Run("returns ok when db is connected", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/ping", nil)
		rec := httptest.NewRecorder()

		PingHandler(store)(rec, req, nil)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, "application/json", res.Header.Get("Content-Type"))

		var got pingResponse
		json.NewDecoder(res.Body).Decode(&got)
		assert.Equal(t, "pong", got.Message)
	})

	t.Run("returns service unavailable when db is not connected", func(t *testing.T) {
		db.Close()

		req := httptest.NewRequest(http.MethodGet, "/ping", nil)
		rec := httptest.NewRecorder()

		PingHandler(store)(rec, req, nil)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusServiceUnavailable, res.StatusCode)
		assert.Equal(t, "application/json", res.Header.Get("Content-Type"))
	})

}
