package response

import (
	"encoding/json"
	"net/http"
)

type Response struct{
	Status string 
	Error string
}

const (
	StatusOk = "OK"
	StatusError = "Error"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data) //encdode convets the struct data to json format 
}

//in student.go there is a fuction that chk that the header body is empty or not 
//instead of returning a EOF , we are tryign to return the json format 
func GeneralError(err error) Response{
	return Response{
		Status: StatusError,
		Error:err.Error(),
	}
}