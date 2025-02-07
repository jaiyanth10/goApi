package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct{
	Status string `json:"status"`
	Error string   `json:"error"`
}

const (StatusOk = "OK"  
       StatusError = "Error" 
	  )

//remember interface{} is parnet class of interface which can holde any type of value
func WriteJson(w http.ResponseWriter, status int , data interface{}) error{

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	//here the the encode method will emcode go types to json(as we mentioned "application/json") and New Encoder will write it w(respose writer).
	return json.NewEncoder(w).Encode(data)

}

func GeneralError(err error) Response {
      return Response{
		Status: StatusError,
		Error: err.Error(),//.Error() is a method that implements the error interface. It is used to convert an error type into a readable string.
	  }
}

func Validationerror(validation_errs validator.ValidationErrors) Response{
       var v_err[] string // creating a slice

	   for _, err_value:= range validation_errs{ //range will always return index and value, we dont want index so we will use _
         switch err_value.ActualTag(){
	     case "required":
			v_err = append(v_err, fmt.Sprintf("%v is required field" , err_value.Field())) // bro, In slice when u append new value u will get another copy of the slice,so store it again same slice variable 
		 default:
			v_err = append(v_err, fmt.Sprint("The sent JSON data is failing validation checks"))
	   }}

	   return Response{
		Status:  StatusError,
		Error: strings.Join(v_err, ","),
	   }
}