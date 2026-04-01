package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/ghost/restAPI/internal/types"
	"github.com/ghost/restAPI/internal/utils/response"
)

//crud operations
func New() http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){

		var student types.Student
		err:=json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err,io.EOF){
				//checks thet the body is empty
				response.WriteJson(w,http.StatusBadRequest,response.GeneralError(err))
				//400 bad request ,if the body is empty
				return 
		}
		slog.Info("Creating a student")
		w.Write([]byte("Welcome to student API"))
		response.WriteJson(w,http.StatusCreated,map[string]string{"success":"OK"})
	}
}