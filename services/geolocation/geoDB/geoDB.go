package geoDB

import (
	"compress/gzip"
	"io"
	"net"
	"net/http"
	"os"
	"sync"

	"github.com/gypsydiver/theweatherservice/services/geolocation/util"
	geoip2 "github.com/oschwald/geoip2-golang"
)

// DB is our main Geoip2 entrypoint
var DB QueryableDB

// QueryableDB functions as an abstraction of geoip2.Reader
// for testing purposes
type QueryableDB interface {
	UpdateDB() error
	QueryIP(net.IP) (*geoip2.City, error)
	openDB() error
	Close() error
}

type geoIPDB struct {
	mutex sync.Mutex
	db    *geoip2.Reader
}

// UpdateDB just refreshens the DB
func (s *geoIPDB) UpdateDB() error {
	if err := downloadDB(); err != nil {
		return err
	}
	return s.openDB()
}

func downloadDB() error {
	//TODO: validate hash/sig
	response, err := http.Get(util.Config.GeoliteDBDownloadURL)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	gzReader, err := gzip.NewReader(response.Body)
	if err != nil {
		return err
	}
	defer gzReader.Close()
	return writeGZContent(gzReader)
}

func writeGZContent(gzReader *gzip.Reader) error {
	gzFile, err := os.OpenFile(util.Config.GeoliteDBName,
		os.O_CREATE|os.O_WRONLY, 0600)
	if err == nil {
		_, err = io.Copy(gzFile, gzReader)
	}
	return err
}

func (s *geoIPDB) openDB() (err error) {
	s.mutex.Lock()
	s.db, err = geoip2.Open(util.Config.GeoliteDBName)
	s.mutex.Unlock()
	return err
}

// QueryIP returns information (if any) regarding the IP provided
func (s *geoIPDB) QueryIP(ip net.IP) (*geoip2.City, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.db.City(ip)
}

func (s *geoIPDB) Close() error {
	return s.db.Close()
}
