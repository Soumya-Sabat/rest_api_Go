package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/ghost/restAPI/internal/types"
	"github.com/ghost/restAPI/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

//crud operations
func New() http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
		slog.Info("Creating a student")

		var student types.Student
		err:=json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err,io.EOF){
				//checks thet the body is empty
				response.WriteJson(w,http.StatusBadRequest,response.GeneralError(err))
				//400 bad request ,if the body is empty
				return 
		}

		if err != nil{
			response.WriteJson(w, http.StatusBadRequest,response.GeneralError(err))
		}

		//validation of requests 
		//using the github url - go get github.com/go-playground/validator/v10 

		if err := validator.New().Struct(student); err!=nil{
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w,http.StatusBadRequest,response.ValidatorError(validateErrs))
			return //to stop frrom further validation 
		}


		

		response.WriteJson(w,http.StatusCreated,map[string]string{"success":"OK"})
	}
}