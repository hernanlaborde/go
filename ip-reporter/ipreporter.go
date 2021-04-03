package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	geodata := geoData()
	fmt.Println(geodata.Country)
	fmt.Println(geodata.City)
	fmt.Println(geodata.Latitude)
	fmt.Println(geodata.Longitude)
	fmt.Println(geodata.Ipaddress)

}

type datosConsulta struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zipcode     string  `json:"zip"`
	Latitude    float64 `json:"lat"`
	Longitude   float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	IspOrg      string  `json:"org"`
	IspAs       string  `json:"as"`
	Ipaddress   string  `json:"query"`
}

func geoData() datosConsulta {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		//return err.Error()
		fmt.Println(err.Error())
	}

	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		//return err.Error()
		fmt.Println(err.Error())
	}

	var geodatos datosConsulta

	//json.Unmarshal(body, &ip)
	err = json.Unmarshal(body, &geodatos)
	if err != nil {
		//return err.Error()
		fmt.Println(err.Error())
	}
	//return ip.Query
	return geodatos
}
