package server

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"

	"github.com/NathanBak/cfgbuild"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCallCounts(t *testing.T) {
	s := initTestServer(t)

	id := uuid.New().String()

	wg := sync.WaitGroup{}

	before, err := s.counter.Lookup(id)
	assert.NoError(t, err)
	assert.Equal(t, 0, before)

	for i := 1; i <= 123; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			simpleCallWithId(t, s, id)
		}()
	}

	wg.Wait()

	after, err := s.counter.Lookup(id)
	assert.NoError(t, err)
	assert.Equal(t, 123, after)

}

func simpleCallWithId(t *testing.T, s *Server, id string) bool {
	w := httptest.NewRecorder()

	path := fmt.Sprintf("/api/v1/memes?lat=%f", rand.Float64())
	// path := "/api/v1/memes"
	r := httptest.NewRequest(http.MethodGet, path, nil)
	r.Header.Add("Authorization", id)

	s.ServeHTTP(w, r)

	return assert.Equal(t, http.StatusOK, w.Result().StatusCode)
}

func initTestServer(t *testing.T) *Server {

	os.Setenv("PORT", "8123")

	builder := cfgbuild.Builder[*Config]{}
	cfg, err := builder.Build()
	if err != nil {
		t.Error(err)
	}

	s, err := New(*cfg)
	if err != nil {
		t.Error(err)
	}

	return s
}
