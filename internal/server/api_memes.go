package server

import (
	"net/http"

	"github.com/NathanBak/maas/pkg/meme"
)

func (s *Server) getMeme(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	meme, err := meme.NewTextMeme("All your base are belong to us")
	if err != nil {
		s.RespondWithError(r.Context(), w, r, http.StatusInternalServerError,
			MAAS1000001, "unable to retrieve meme", nil)
		return
	}

	s.RespondWithJSON(ctx, w, http.StatusOK, &meme)
}
