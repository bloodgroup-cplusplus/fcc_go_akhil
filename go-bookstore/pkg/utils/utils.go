package utils 

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// create function 


func ParseBody(r *http.Request, x interface{}) {
	if body,err := ioutil.ReadAll(r.Body); err == nil {
		if err :=json.Unmarshal([] byte(body),x); err!=nil{
			return 
		}
	}
}
	// while creating the book
