package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// TODO add test for this function
func ReadObjectFromJSONInput(toUnmarshallObj interface{}, writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(request.Body, 1048576))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	if err := request.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &toUnmarshallObj); err != nil {
		writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
		writer.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(writer).Encode(err); err != nil {
			panic(err)
		}
	}
}
