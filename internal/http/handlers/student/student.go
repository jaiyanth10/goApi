package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jaiyanth10/goApi/internal/types"
	"github.com/jaiyanth10/goApi/internal/utils/response"
)

// The below function will return http.HandlerFunc type “
// The below function will return mutiple http methods associated functions written in this file.
func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		slog.Info("creating new student info")
		var student types.Student

		//checking errors

		err := json.NewDecoder(r.Body).Decode(&student) //here we are serializing the reciving json data in r.body and storing in student struct which automatically coverts to struct format
		if errors.Is(err, io.EOF) {                     //checking if the request body is empty
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body"))) //fmt.ErrorF will convert nrml string to Error
			return                                                                                        // to stop execution
		}
		if err != nil { //checking for other errors
			response.WriteJson(w, http.StatusBadGateway, response.GeneralError(err))
			return
		}

		//Validating request using validator(external package)
        er:= validator.New().Struct(student) // we are passing the struct(which is associated with data in request body) to validate it according to struct tags(see type.go for tags)
		if er!=nil{
			validation_err:=er.(validator.ValidationErrors)//type casting err's type to validatioErrors type as validationerroe function expects this type
            response.WriteJson(w, http.StatusBadGateway, response.Validationerror(validation_err))
			return
		}

		slog.Info("Creating Student!")
		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "ok!"}) //data is sent in map format ; map syntax - map[key data type] value data type{data}
	}
}
