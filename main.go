package main

// https://www.callicoder.com/golang-packages/
import (	
	"artifactorytools/netutils"
	"fmt"
	"log"
)

func main(){
	fmt.Println("Ready !!")
	netutils.Login = ""
	netutils.Password = ""
	netutils.URL = ""
	getNbOlderThan(600)
}

//https://medium.com/@masnun/making-http-requests-in-golang-dd123379efe7

func getNbOlderThan(day int) (int, []interface{}) {
	aql := fmt.Sprintf(`items.find( 
			{
				"repo": "platformdxc-docker",
				"modified" : {"$before" : "%dd"}
			})`,day)

	res := netutils.SendAQL ( aql)
	s := fmt.Sprintf("%d", len(res))

	log.Println ("getNbOlderThan: " + s)
	
	return  len(res), res
}



