package main

// https://www.callicoder.com/golang-packages/
import (	
	"artifactorytools/netutils"
	"fmt"
	"log"
	"strconv"
	"strings"

)

func main(){
	fmt.Println("Ready !!")
	netutils.Login = ""
	netutils.Password = ""
	netutils.URL = ""
	i := getNbArtifact()
	log.Println(i)
	//getNbOlderThan(600)
}

//https://medium.com/@masnun/making-http-requests-in-golang-dd123379efe7

func execAQL(aql string) (int, []interface{}) {
	res := netutils.SendAQL ( aql)
	s := fmt.Sprintf("%d", len(res))
	log.Println ("Number of results: " + s)
	
	return  len(res), res
}

func getNbOlderThan(day int) (int, []interface{}) {
	aql := fmt.Sprintf(`items.find( 
			{
				"repo": "platformdxc-docker",
				"modified" : {"$before" : "%dd"}
			})`,day)

	return execAQL(aql)
}

func getNbArtifact() int {
	data := netutils.GetRestAPI("/artifactory/api/storageinfo")
	log.Println(data)
	c := data["binariesSummary"].(map[string]interface{})
	var tmp = c["artifactsCount"].(string)
	tmp = strings.Replace(tmp, ",", "", -1)
	res,err := strconv.Atoi(tmp)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println ("getNbArtifact :" + tmp)
	return res
}


