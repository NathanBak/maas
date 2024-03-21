package server

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLivez(t *testing.T) {
	s := initTestServer(t)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/livez", nil)

	s.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Result().StatusCode)

	body, err := io.ReadAll(w.Body)
	assert.NoError(t, err)

	var m responseMessage
	err = json.Unmarshal(body, &m)
	assert.NoError(t, err)

	assert.Equal(t, "It's alive!", m.Message)
}

func TestReadyz(t *testing.T) {
	s := initTestServer(t)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/readyz", nil)

	s.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Result().StatusCode)

	body, err := io.ReadAll(w.Body)
	assert.NoError(t, err)

	var m responseMessage
	err = json.Unmarshal(body, &m)
	assert.NoError(t, err)

	assert.Equal(t, "Ready to rock!", m.Message)
}
