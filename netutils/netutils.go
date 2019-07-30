package netutils

import (
	"net/http"
	"io/ioutil"
	"strings"
	"log"
	"encoding/base64"
	"encoding/json"
)

// Login : Artifactory login
var Login string


// Password : Artifactory password
var Password string

// URL : Url of the artifactory server (https://FQDN)
var URL string

func basicAuth(username, password string) string {
	auth := username + ":" + password
	res := base64.StdEncoding.EncodeToString([]byte(auth))
	log.Println( res)

	return res
}

// SendAQL : Send an aql query
func SendAQL(aql string) []interface{} {
	client := http.Client {}
	request, err := http.NewRequest("POST", URL + "/artifactory/api/search/aql",  strings.NewReader(aql))
	request.Header.Add("Content-Type", "text/plain")
	request.Header.Add("Authorization","Basic " + basicAuth(Login,Password))
	
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

	return data["results"].([]interface{})
}