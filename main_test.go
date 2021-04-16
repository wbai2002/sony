package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestGetGeoMap(t *testing.T) {
	file, err := ioutil.ReadFile("data.json")
	if err != nil {
		panic(err)
	}
	var p getCountryCodeOutput
	err = p.getGeoMap(&getGeoMapInput{
		rawFile: file,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(p.geoMap)
}
