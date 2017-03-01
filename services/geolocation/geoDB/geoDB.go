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

var (
	db    *geoip2.Reader
	mutex sync.Mutex
)

// UpdateDB just refreshens the DB
func UpdateDB() error {
	if err := downloadDB(); err != nil {
		return err
	}
	return openDB()
}

func downloadDB() error {
	//TODO: validate hash
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

func openDB() (err error) {
	mutex.Lock()
	db, err = geoip2.Open(util.Config.GeoliteDBName)
	mutex.Unlock()
	return err
}

func writeGZContent(gzReader *gzip.Reader) error {
	gzFile, err := os.OpenFile(util.Config.GeoliteDBName, os.O_CREATE|os.O_WRONLY, 0660)
	if err == nil {
		_, err = io.Copy(gzFile, gzReader)
	}
	return err
}

// QueryIP returns information (if any) regarding the IP provided
func QueryIP(ip net.IP) (*geoip2.City, error) {
	mutex.Lock()
	defer mutex.Unlock()
	return db.City(ip)
}
