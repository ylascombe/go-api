package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
	"github.com/gorilla/mux"
)

type Manifest struct {
	FormatVersion string           `json:"format_version" yaml:"format_version"`
	ReactPlatform ReactivePlatform `json:"reactive_platform" yaml:"reactive_platform"`
	Applications  []Application    `json:"applications" yaml:"applications"`
}

type ReactivePlatform struct {
	Version   string `json:"version" yaml:"version"`
	ExtraVars map[string]string `json:"extra_vars" yaml:"extra_vars"`
	FeaturesStatus map[string]string `json:"features_status" yaml:"features_status"`
}

type Application struct {
	Name  string `json:"name" yaml:"name"`
	Spark Spark  `json:"spark" yaml:"spark"`
	Api   Api    `json:"api" yaml:"api"`
}

type Spark struct {
	Version   string `json:"version" yaml:"version"`
	ExtraVars map[string]string `json:"extra_vars" yaml:"extra_vars"`
}

type Api struct {
	Version   string `json:"version" yaml:"version"`
	ExtraVars map[string]string `json:"extra_vars" yaml:"extra_vars"`
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/reactive-platform/target/{target}/manifestversion/{version}", api)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func api(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	target := vars["target"]
	version := vars["version"]
	fmt.Println("target: ", target)
	fmt.Println("version: ", version)
	data := readFile("manifest.yml")

	fmt.Println("data: %v", string(data))
	config := unmarshall(data)

	fmt.Println("resultat: \n", config.ReactPlatform.Version)
}

func unmarshall(yamlText []byte) *Manifest {
	var config Manifest
	var err = yaml.Unmarshal(yamlText, &config)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil
	}

	return &config
}

func readFile(filePath string) []byte {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 1000)
	_, err = file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	count := len(data)
	//fmt.Printf("read %d bytes: %q\n", count, data[:count])
	return data[:count]
}