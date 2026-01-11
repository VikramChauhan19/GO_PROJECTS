package util //util package contain helper func use in controllers

import (
	"encoding/json" //json to struct or struct to json
	"net/http"// provide http.Request
	"io" // use to read from request body
)
func ParseBody(r *http.Request, x interface{}){ //ParseBody extracts JSON from the request body and converts it into a Go struct.
	body , err := io.ReadAll(r.Body)// ReadAll reads from r.Body until an error or EOF and returns the data it read.
	if err == nil{
		err :=json.Unmarshal([]byte(body),x)//it takes a byte slice of JSON data and a pointer to a Go variable where the decoded data will be stored.
		if err != nil{
			return
		}
	}
}