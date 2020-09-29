package controllers

import (
	"encoding/json"
	"github.com/nickdipreta/golang-microservices/mvc/services"
	"github.com/nickdipreta/golang-microservices/mvc/utils"
	"net/http"
	"strconv"
)
// controller validates that we have all of the parameters that we need
// if so, we send this to the service
func GetUser(resp http.ResponseWriter, req *http.Request) {
	// takes userId, if there is an error create Application Error
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		//404 requests/bad requests handled here
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		// try to marshal as a json and write the error code as a header and write the json value
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}
	user, apiErr := services.GetUser(userId)
	// if there is an error, put the status code as the header, and write the apiErr
	if apiErr != nil {
		//handle the err and return to the client
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(jsonValue))
		return
	}
	//return user to client
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)

}
