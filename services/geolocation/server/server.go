package server

import (
	"context"

	geoGen "github.com/gypsydiver/theweatherservice/services/geolocation/generated"
)

func init() {
	// Init stuff
}

// GeolocationServer implementation, it will serve all ip location requests
type GeolocationServer struct {
	// vars
}

// Locate receives a context and an Array of IPs to transform into LatLong
// values, which is just a lookup in the Geolite2 database
func (s *GeolocationServer) Locate(ctx *context.Context, req *geoGen.LocateRequest) (*geoGen.LocateResponse, error) {
	return &geoGen.LocateResponse{}, nil
}
