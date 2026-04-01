package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct{
	Status string `json:"status"` //custom case for empty body 
	Error string `json:"error"`
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


func ValidatorError(errs validator.ValidationErrors) Response{
	var errMsgs []string
	for _,err:=range errs{
		switch err.ActualTag(){
		case "requried":
			errMsgs=append(errMsgs,fmt.Sprintf("feild %s is required : ",err.Field()) )
		default:
			errMsgs=append(errMsgs,fmt.Sprintf("feild %s is invalid : ",err.Field()) )

		}
	}


	return Response{
		Status: StatusError,
		Error: strings.Join(errMsgs,","), //joins the slices recieved from the errMsgs separated by commas
	}
}