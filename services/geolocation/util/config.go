package util

//Configuration holds all values relevant to the connector and its functions
type Configuration struct {
	GeoliteDBDownloadURL     string `yaml:"GeoliteDBDownloadURL"`
	GeoliteDBName            string `yaml:"GeoliteDBName"`
	IntervalUpdateDBInMonths int    `yaml:"IntervalUpdateDBInMonths"`
}

// Config is a global obj to be used throughout the connector
var Config = Configuration{
	GeoliteDBDownloadURL:     "http://geolite.maxmind.com/download/geoip/database/GeoLite2-City.mmdb.gz",
	GeoliteDBName:            "GeoLite2-City.mmdb",
	IntervalUpdateDBInMonths: 1,
}
