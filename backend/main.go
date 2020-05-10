// https://www.soberkoder.com/consume-rest-api-go/
// https://www.youtube.com/watch?v=W5b64DXeP0o
// problems dealing with "github.com/likexian/whois-go"
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

//request to ssllabscccc
type Request struct {
	ServersChanged   string `json:""`
	SslGrade         string `json:"ssl_grade"`
	PreviousSslGrade string `json:"previous_ssl_grade"`
	Logo             string `json:"logo"`
	Title            string `json:"title"`
	IsDown           bool   `json:"is_down "`
	Endpoints        []Endpoint
}

//endpoint ssllabs
type Endpoint struct {
	IpAddress string
	Grade     string
	Country   string
	Owner     string
}

// IPInfo describes a particular IP address.
type IPInfo struct {
	// IP holds the described IP address.
	IP string
	// Hostname holds a DNS name associated with the IP address.
	Hostname string
	// City holds the city of the ISP location.
	City string
	// Country holds the two-letter country code.
	Country string
	// Loc holds the latitude and longitude of the
	// ISP location as a comma-separated northing, easting
	// pair of floating point numbers.
	Loc string
	// Org describes the organization that is
	// responsible for the IP address.
	Org string
	// Postal holds the post code or zip code region of the ISP location.
	Postal string
}

func makeRequest(ctx *fasthttp.RequestCtx) {
	domain := ctx.UserValue("domain").(string)
	url := "https://api.ssllabs.com/api/v3/analyze?host="
	//creating request
	resp, err := http.Get(url + domain)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var server Request
	err = json.Unmarshal(body, &server)
	log.Println(server)

	if err != nil {
		log.Println(err)
	}

	for i, item := range server.Endpoints {
		ip := item.IpAddress
		response, err := http.Get("http://ipinfo.io/" + ip + "/json")
		if err != nil {
			log.Println(err)
		}
		defer response.Body.Close()

		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Println(err)
		}
		var ipinfo IPInfo
		if err := json.Unmarshal(contents, &ipinfo); err != nil {
			log.Println(err)
		}
		log.Println(i, ipinfo.Country, ipinfo.Org)
	}

	json.NewEncoder(ctx).Encode(server)
}

func homePage(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome to my first Go project!. Please go to /domain/yourdomain \n")
}

func main() {
	router := fasthttprouter.New()
	router.GET("/", homePage)
	//handle endpoint
	router.GET("/domain/:domain", makeRequest)
	log.Fatal(fasthttp.ListenAndServe(":8081", router.Handler))
}
