package main

import (
   "fmt"
   "html"
   "log"
   "net/http"
   "os"

   "github.com/gorilla/mux"
   //"github.com/ghodss/yaml"
	"encoding/json"
)

type Manifest struct {
   formatVersion    string `json:"format_version" yaml:"format_version"`
   //reactivePlatform ReactivePlatform `json:"reactive_platform" yaml:"reactive_platform"`
   //applications []Application `json:"applications" yaml:"applications"`
}

type ReactivePlatform struct {
   Version string `json:"version" yaml:"version"` // Affects YAML field names too.
   ExtraVars  int    `json:"extra_vars" yaml:"extra_vars"`
}

type Application struct {
   name string `json:"name" yaml:"name"`
   spark Spark `json:"spark" yaml:"spark"`
   api Api `json:"api" yaml:"api"`

}

type Spark struct {

}

type Api struct {

}

func main() {

   router := mux.NewRouter().StrictSlash(true)
   router.HandleFunc("/", Index)
   router.HandleFunc("/reac", api)
   log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func api(w http.ResponseWriter, r *http.Request) {

   file, err := os.Open("manifest.json") // For read access.
   if err != nil {
       log.Fatal(err)
   }

   data := make([]byte, 1000)
   count, err := file.Read(data)
   if err != nil {
       log.Fatal(err)
   }
   fmt.Printf("read %d bytes: %q\n", count, data[:count])

	var config Manifest
//	err = yaml.Unmarshal(data[:count], &config)
//	if err != nil {
//		fmt.Printf("err: %v\n", err)
//		return
//	}
//	fmt.Println(config)
//	fmt.Println(config.formatVersion)

	data = []byte(`{"format_version": "0.1"}`)

	count = len(data)
	if err := json.Unmarshal(data[:count], &config); err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(config)
}
