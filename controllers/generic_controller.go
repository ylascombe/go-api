package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/ylascombe/go-api/utils"
)

func genericReadResponse(writer http.ResponseWriter, request *http.Request, result interface{}, err error) {
	if err == nil {
		// YAML Marshalling : text := utils.Marshall(users)
		text, err := json.Marshal(result)

		if err == nil {
			writer.WriteHeader(200)
			fmt.Fprintf(writer, string(text))
			return
		}
	}

	// if this code is executed, so there is an error
	writer.WriteHeader(500)
	resp := ApiResponse{ErrorMessage: string(err.Error())}
	text := utils.Marshall(resp)
	fmt.Fprintf(writer, text)
}
