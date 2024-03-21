package server

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"sync"
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

func TestManyMemes(t *testing.T) {
	s := initTestServer(t)

	wg := sync.WaitGroup{}

	for i := 1; i <= 10000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			simpleCall(t, s)
		}()
	}

	wg.Wait()
}

func simpleCall(t *testing.T, s *Server) bool {
	w := httptest.NewRecorder()

	path := fmt.Sprintf("/api/v1/memes?lat=%f", rand.Float64())
	// path := "/api/v1/memes"
	r := httptest.NewRequest(http.MethodGet, path, nil)

	s.ServeHTTP(w, r)

	return assert.Equal(t, http.StatusOK, w.Result().StatusCode)
}
