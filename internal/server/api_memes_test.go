package server

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/NathanBak/maas/pkg/meme"
	"github.com/stretchr/testify/assert"
)

func TestMemes01(t *testing.T) {
	s := initTestServer(t)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/api/v1/memes", nil)

	s.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Result().StatusCode)

	body, err := io.ReadAll(w.Body)
	assert.NoError(t, err)

	var m meme.Meme
	err = json.Unmarshal(body, &m)
	assert.NoError(t, err)

	assert.Equal(t, "All your base are belong to us", m.Text)
	assert.Empty(t, m.URL)
	assert.NotEmpty(t, m.ID)
}
