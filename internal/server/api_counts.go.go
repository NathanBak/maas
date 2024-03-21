package server

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (s *Server) getCount(w http.ResponseWriter, r *http.Request, userId uuid.UUID) {
	ctx := r.Context()

	count, err := s.counter.Lookup(userId.String())
	if err != nil {
		s.RespondWithError(r.Context(), w, r, http.StatusInternalServerError,
			MAAS1000007, "unable to count", nil)
		return
	}

	s.RespondWithMessage(ctx, w, http.StatusOK, fmt.Sprintf("You have made %d api calls", count))
}
