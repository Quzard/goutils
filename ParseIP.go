package goutils

import (
	"github.com/oschwald/geoip2-golang"
	"net"
)

var DB *geoip2.Reader

//  path GeoLite2-City.mmdb 地址
func initDB(path string) error {
	var err error
	DB, err = geoip2.Open(path)
	return err
}

// ParseIP 通过ip地址获取IP的城市和国家
func ParseIP(blockIP string) (city, country string, err error) {
	ip := net.ParseIP(blockIP)
	record, err := DB.City(ip)
	if err != nil {
		return "", "", err
	}
	city = record.City.Names["zh-CN"]
	country = record.Country.Names["zh-CN"]
	return city, country, nil

	//fmt.Printf("Portuguese (BR) city name: %v\n", record.City.Names["zh-CN"])
	//fmt.Printf("English subdivision name: %v\n", record.Subdivisions[0].Names["en"])
	//fmt.Printf("Russian country name: %v\n", record.Country.Names["en"])
	//fmt.Printf("ISO country code: %v\n", record.Country.IsoCode)
	//fmt.Printf("Time zone: %v\n", record.Location.TimeZone)
	//fmt.Printf("Coordinates: %v, %v\n", record.Location.Latitude, record.Location.Longitude)
	// Output:
	// Portuguese (BR) city name: Londres
	// English subdivision name: England
	// Russian country name: Великобритания
	// ISO country code: GB
	// Time zone: Europe/London
	// Coordinates: 51.5142, -0.0931
}
