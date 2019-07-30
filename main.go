package main

import (
	//"bytes"
	"io/ioutil"
	"strings"
	"fmt"
	"net/http"
	"log"
	//"time"
	"encoding/base64"
	"encoding/json"
	"reflect"
	//"crypto/tls"

	//"jfrog-client-go"
)

func main(){

	fmt.Println("Ready !!")
	sendAQL("<artifactory url>/artifactory/api/search/aql", getAQL())
}

//https://artifactory.csc.com/artifactory/api/search/aql
//https://medium.com/@masnun/making-http-requests-in-golang-dd123379efe7

func basicAuth(username, password string) string {
	auth := username + ":" + password
	res := base64.StdEncoding.EncodeToString([]byte(auth))
	log.Println( res)

	return res
}

func  getAQL() string {
	return "items.find()"
}

func sendAQL(url string, aql string){
	/*tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	

	timeout := time.Duration(3000)*/
	client := http.Client {}

	//request, err := http.NewRequest("POST", url, strings.NewReader(json))
	
	request, err := http.NewRequest("POST", url,  strings.NewReader(aql))
	request.Header.Add("Content-Type", "text/plain")
	request.Header.Add("Authorization","Basic " + basicAuth("login","password"))
	
	
	if err != nil {
		log.Fatalln(err)
	}

	resp,err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var data map[string]interface{}

	err = json.Unmarshal(body, &data)
	
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(len(data["results"].([]interface{})))
}

