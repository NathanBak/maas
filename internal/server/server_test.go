package server

import (
	"os"
	"testing"

	"github.com/NathanBak/cfgbuild"
)

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
