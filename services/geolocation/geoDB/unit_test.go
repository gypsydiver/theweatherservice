package geoDB

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/gypsydiver/theweatherservice/services/geolocation/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GeoDBUnitSuite struct {
	suite.Suite
	A            *assert.Assertions
	mockDBServer *httptest.Server
	DB           QueryableDB
}

func (s *GeoDBUnitSuite) SetupSuite() {
	s.A = assert.New(s.T())
}

func (s *GeoDBUnitSuite) SetupTest() {
	s.DB = &geoIPDB{}
	s.mockDBServer = httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "./test-db.gz")
		}))
}

func (s *GeoDBUnitSuite) TearDownSuite() {
	s.mockDBServer.Close()
	s.DB.Close()
}

func (s *GeoDBUnitSuite) TestDownloadMalformedGZBody() {
	s.mockDBServer = httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "whatever")
		}))
	util.Config.GeoliteDBName = "wrong-db.mmdb"
	util.Config.GeoliteDBDownloadURL = s.mockDBServer.URL

	err := downloadDB()
	s.A.NotNil(err)
}

func (s *GeoDBUnitSuite) TestQueryIP() {
	util.Config.GeoliteDBName = "test-db.mmdb"
	s.DB.openDB()
	ip := net.ParseIP("123.125.71.29")

	_, err := s.DB.QueryIP(ip)
	s.A.Nil(err)
}

func (s *GeoDBUnitSuite) TestUpdateDB() {
	util.Config.GeoliteDBName = "wrong-db.mmdb"
	util.Config.GeoliteDBDownloadURL = s.mockDBServer.URL

	err := s.DB.UpdateDB()

	s.A.Nil(err)
	_, err = os.Stat(util.Config.GeoliteDBName)
	s.A.Nil(err)
	os.Remove(util.Config.GeoliteDBName)

	util.Config.GeoliteDBDownloadURL = "///"
	err = s.DB.UpdateDB()
	s.A.NotNil(err)
}
