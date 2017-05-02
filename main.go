package main

import (
   "fmt"
   "html"
   "log"
   "net/http"
   "os"

   "github.com/gorilla/mux"
   //"github.com/ghodss/yaml"
)

type Manifest struct {
   formatVersion    string `json:"format_version" yaml:"format_version"`
   reactivePlatform ReactivePlatform `json:"reactive_platform" yaml:"reactive_platform"`
   applications []Application `json:"applications" yaml:"applications"`
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

   file, err := os.Open("manifest.yml") // For read access.
   if err != nil {
       log.Fatal(err)
   }

   data := make([]byte, 100)
   count, err := file.Read(data)
   if err != nil {
       log.Fatal(err)
   }
   fmt.Printf("read %d bytes: %q\n", count, data[:count])
}
