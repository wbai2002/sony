package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	output, err := getCountryCode(&getCountryCodeInput{
		osArgs: os.Args,
	})
	if err != nil {
		panic(err)
	}

	// print help message
	if output.helpMsg != "" {
		fmt.Println(output.helpMsg)
		return
	}
	file, err := ioutil.ReadFile("data.json")
	if err != nil {
		panic(err)
	}
	err = output.getGeoMap(&getGeoMapInput{
		rawFile: file,
	})
	if err != nil {
		panic(err)
	}
	if output.server {
		con := http.NewServeMux()
		con.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			output.handler(w, r)
		})
		panic(Server(&ServerConfig{
			BindPort: "8080",
			ServeMux: con,
		}))
		return
	}
	for _, countryCode := range output.countryCodes {
		fmt.Println(output.geoMap[strings.ToUpper(countryCode)])
	}
}

type getGeoMapInput struct {
	rawFile []byte
}

func (p *getCountryCodeOutput) getGeoMap(i *getGeoMapInput) error {
	var currentCountryCode string
	toReturn := make(map[string]string)
	for _, line := range bytes.Split(i.rawFile, []byte("\n")) {
		if bytes.Contains(line, []byte("\"iso_alpha2\"")) {
			splitOne := bytes.Split(line, []byte("\"iso_alpha2\": \""))
			countryCode := bytes.Replace(splitOne[1], []byte("\","), []byte(""), -1)
			currentCountryCode = string(countryCode)
		}
		if bytes.Contains(line, []byte("\"name\"")) {
			splitOne := bytes.Split(line, []byte("\"name\": \""))
			country := bytes.Replace(splitOne[1], []byte("\","), []byte(""), -1)
			toReturn[currentCountryCode] = string(country)
		}
	}
	p.geoMap = toReturn
	return nil
}
func (p *getCountryCodeOutput) handler(w http.ResponseWriter, r *http.Request) {
	/* Logic on GET method, then return 405 for invalid method */
	/* Logic on r.RequestURI on /convert */
	/* logic on URI /metrics make sure of github.com/prometheus/client_golang/prometheus */
	return
}

type getCountryCodeInput struct {
	osArgs []string
}
type getCountryCodeOutput struct {
	helpMsg      string
	server       bool
	countryCodes []string
	geoMap       map[string]string
}

func getCountryCode(i *getCountryCodeInput) (getCountryCodeOutput, error) {
	if len(i.osArgs) < 2 {
		return getCountryCodeOutput{}, fmt.Errorf("Not enough arguments e.g lookup --countryCode=AU")
	}

	// Specify help msg
	var output getCountryCodeOutput
	usageMsg := "Usage:\n./lookup --countryCode=AU\nTo start a server do ./lookup -s"
	for _, args := range i.osArgs {
		if args == "-h" || args == "--help" {
			return getCountryCodeOutput{
				helpMsg: usageMsg,
			}, nil
		}
		if args == "-s" {
			return getCountryCodeOutput{
				server: true}, nil
		}
		if strings.Contains(args, "=") {
			output.countryCodes = append(output.countryCodes, strings.Split(args, "=")[1])
		}
	}
	//
	if len(output.countryCodes) == 0 {
		return getCountryCodeOutput{
			helpMsg: usageMsg,
		}, nil
	}
	return output, nil
}
