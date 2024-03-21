package server

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/NathanBak/maas/pkg/meme"
)

func (s *Server) getMeme(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	latVal := r.URL.Query().Get("lat")
	lonVal := r.URL.Query().Get("lon")
	query := r.URL.Query().Get("query")

	var lat float64
	if latVal != "" {
		parsed, err := strconv.ParseFloat(latVal, 64)
		if err != nil {
			s.RespondWithError(r.Context(), w, r, http.StatusBadRequest,
				MAAS1000001, "invalid lattitude", nil)
			return
		}
		// TODO:  Verify latitude is valid
		lat = parsed
	}

	var lon float64
	if lonVal != "" {
		parsed, err := strconv.ParseFloat(lonVal, 64)
		if err != nil {
			s.RespondWithError(r.Context(), w, r, http.StatusBadRequest,
				MAAS1000002, "invalid longitude", nil)
			return
		}
		// TODO:  Verify longitude is valid
		lon = parsed
	}

	if query != "" {
		// TODO: Add checking for injection attacks
		if strings.TrimSpace(query) == "" {
			s.RespondWithError(r.Context(), w, r, http.StatusBadRequest,
				MAAS1000003, "invalid query", nil)
			return
		}
	}

	s.log.Debug(fmt.Sprintf("Lon: %f, Lat: %f, Query: %s", lon, lat, query))

	meme, err := meme.NewTextMeme("All your base are belong to us")
	if err != nil {
		s.RespondWithError(r.Context(), w, r, http.StatusInternalServerError,
			MAAS1000004, "unable to retrieve meme", nil)
		return
	}

	s.RespondWithJSON(ctx, w, http.StatusOK, &meme)
}
