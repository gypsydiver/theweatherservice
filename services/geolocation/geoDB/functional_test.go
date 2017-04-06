package geoDB

import (
	"net"
	"os"

	"github.com/gypsydiver/theweatherservice/services/geolocation/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GeoDBFunctionalSuite struct {
	suite.Suite
	A  *assert.Assertions
	DB QueryableDB
}

func (s *GeoDBFunctionalSuite) SetupSuite() {
	s.A = assert.New(s.T())
}

func (s *GeoDBFunctionalSuite) TearDownSuite() {
	osErr := os.Remove("GeoLite2-City.mmdb")
	s.A.Nil(osErr)
	s.DB.Close()
}

func (s *GeoDBFunctionalSuite) SetupTest() {
	s.DB = &geoIPDB{}
}

func (s *GeoDBFunctionalSuite) TestAUpdate() {
	err := s.DB.UpdateDB()
	s.A.Nil(err)

	util.Config.GeoliteDBDownloadURL = "///"
	err = s.DB.UpdateDB()
	s.A.NotNil(err)
}

func (s *GeoDBFunctionalSuite) TestBQueryIP() {
	s.DB.openDB()
	_, err := s.DB.QueryIP(net.ParseIP("127.0.0.1"))
	s.A.Nil(err)
}
