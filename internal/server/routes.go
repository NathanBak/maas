package server

import "net/http"

// A Route describes a rest endpoint.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// routes returns an array of all Routes.
func (s *Server) routes() []Route {
	return []Route{

		{
			Name:        "Livez",
			Method:      http.MethodGet,
			Pattern:     "/livez",
			HandlerFunc: s.livez,
		},

		{
			Name:        "Readyz",
			Method:      http.MethodGet,
			Pattern:     "/readyz",
			HandlerFunc: s.readyz,
		},

		{
			Name:        "GetMeme",
			Method:      http.MethodGet,
			Pattern:     "/api/v1/memes",
			HandlerFunc: s.countApiCall(s.getMeme),
		},

		{
			Name:        "GetCount",
			Method:      http.MethodGet,
			Pattern:     "/api/v1/counts",
			HandlerFunc: s.countApiCall(s.getCount),
		},
	}
}
