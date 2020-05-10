// Create basic API REST https://www.soberkoder.com/consume-rest-api-go - https://www.youtube.com/watch?v=W5b64DXeP0o
// https://ipinfo.io/ instead of whois "github.com/likexian/whois-go"
// get meta https://home.urlmeta.org/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

//request to ssllabs
type Request struct {
	ServersChanged   string
	SslGrade         string `json:"ssl_grade"`
	PreviousSslGrade string `json:"previous_ssl_grade"`
	Logo             string `json:"logo"`
	Title            string `json:"title"`
	Status           string
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
	IP       string
	Hostname string
	Country  string
	Org      string
}

// Meta to https://home.urlmeta.org/
type MetaInfo struct {
	Meta MetaObj
}

type MetaObj struct {
	Title string
	Image string
}

func getSsllabs(ctx *fasthttp.RequestCtx) Request {
	domain := ctx.UserValue("domain").(string)
	url := "https://api.ssllabs.com/api/v3/analyze?host="
	// get request
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
	if err != nil {
		log.Println(err)
	}
	return server
}

func getIpInfo(ctx *fasthttp.RequestCtx, ip string) IPInfo {
	url := "http://ipinfo.io/" + ip + "/json"
	//get request
	response, err := http.Get(url)
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
	return ipinfo
}

func getMetaInfo(ctx *fasthttp.RequestCtx, domain string) MetaInfo {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.urlmeta.org/?url=https://"+domain, nil)
	req.Header.Add("Authorization", "Basic "+os.Getenv("key"))
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var meta MetaInfo
	if err := json.Unmarshal(contents, &meta); err != nil {
		log.Println(err)
	}
	return meta
}

func makeRequest(ctx *fasthttp.RequestCtx) {
	domain := ctx.UserValue("domain").(string)
	var server Request
	//request to the serverinfo
	server = getSsllabs(ctx)
	if server.Status == "READY" {
		log.Println("true")
	}
	grade := "@" //ASCII.
	for i := range server.Endpoints {
		ip := server.Endpoints[i].IpAddress
		if server.Endpoints[i].Grade > grade {
			grade = server.Endpoints[i].Grade
		}
		//request to the ipinfo
		var ipinfo IPInfo
		ipinfo = getIpInfo(ctx, ip)
		server.Endpoints[i].Country = ipinfo.Country
		server.Endpoints[i].Owner = ipinfo.Org
	}
	if grade != "@" {
		server.SslGrade = grade
	}
	var meta MetaInfo
	//request to the metadata
	meta = getMetaInfo(ctx, domain)
	server.Title = meta.Meta.Title
	server.Logo = meta.Meta.Image
	json.NewEncoder(ctx).Encode(server)
}

func homePage(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome to my first Go project!. Please go to /domain/yourdomain \n")
}

func test() {
	router := fasthttprouter.New()
	router.GET("/", homePage)
	router.GET("/domain/:domain", makeRequest)
	log.Fatal(fasthttp.ListenAndServe(":8081", router.Handler))
}
