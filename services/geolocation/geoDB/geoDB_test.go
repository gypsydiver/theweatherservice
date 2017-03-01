package geoDB

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gypsydiver/theweatherservice/services/geolocation/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GeoDBTestSuite struct {
	suite.Suite
	augmentationName string
	A                *assert.Assertions
	mockDBServer     *httptest.Server
}

func (s *GeoDBTestSuite) SetupSuite() {
	s.A = assert.New(s.T())
}

func (s *GeoDBTestSuite) SetupTest() {
	db = nil
	s.mockDBServer = httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "./test-db.gz")
		}))
}

func (s *GeoDBTestSuite) TearDownSuite() {
	s.mockDBServer.Close()
}

func (s *GeoDBTestSuite) TestDownloadMalformedGZBody() {
	s.mockDBServer = httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "whatever")
		}))
	util.Config.GeoliteDBName = "wrong-db.mmdb"
	util.Config.GeoliteDBDownloadURL = s.mockDBServer.URL

	err := downloadDB()
	s.A.NotNil(err)
}

func (s *GeoDBTestSuite) TestQueryIP() {
	util.Config.GeoliteDBName = "test-db.mmdb"
	openDB()
	ip := net.ParseIP("123.125.71.29")

	_, err := QueryIP(ip)
	s.A.Nil(err)
}

func (s *GeoDBTestSuite) TestUpdateDB() {
	util.Config.GeoliteDBName = "wrong-db.mmdb"
	util.Config.GeoliteDBDownloadURL = s.mockDBServer.URL

	err := UpdateDB()

	s.A.Nil(err)
	_, err = os.Stat(util.Config.GeoliteDBName)
	s.A.Nil(err)
	os.Remove(util.Config.GeoliteDBName)

	util.Config.GeoliteDBDownloadURL = "///"
	err = UpdateDB()
	s.A.NotNil(err)
}

func TestGeoDB(t *testing.T) {
	suite.Run(t, new(GeoDBTestSuite))
}
