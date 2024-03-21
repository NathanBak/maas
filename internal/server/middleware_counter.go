package server

import (
	"net/http"

	"github.com/google/uuid"
)

type handleWithUserId func(w http.ResponseWriter, r *http.Request, userId uuid.UUID)

func (s *Server) countApiCall(next handleWithUserId) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// get auth header
		id := r.Header.Get("Authorization")

		// TODO: Verify auth header and return 401 if invalid

		var uid uuid.UUID

		if id != "" {
			err := s.counter.Increment(id, 1)
			if err != nil {
				s.RespondWithError(r.Context(), w, r, http.StatusInternalServerError,
					MAAS1000005, "unable to access counter", nil)
				return
			}

			uid, err = uuid.Parse(id)
			if err != nil {
				s.RespondWithError(r.Context(), w, r, http.StatusInternalServerError,
					MAAS1000006, "invalid auth header", nil)
				return
			}
		}

		next(w, r, uid)
	})
}
