package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
	"github.com/gorilla/mux"
	"errors"
	"io/ioutil"
	"sync"
	"os/exec"
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

func api(_ http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	target := vars["target"] + ".ini"
	version := vars["version"]
	fmt.Println("target: ", target)
	fmt.Println("version: ", version)

	ansibleCommands := buildCommands(target, version)
	wg := new(sync.WaitGroup)
	wg.Add(3)
	for i:=0; i<len(ansibleCommands); i++ {
		execCommand("echo " + ansibleCommands[i],wg)
	}


}

func buildCommands(target string, manifest string) []string {
	config, _ := unmarshallFromFile("manifest.yml")

	fmt.Println("resultat: \n", config.ReactPlatform.Version)

	var ansibleCommands = []string {}
	ansibleCommands = append(ansibleCommands, "ansible-playbook -i inventories/" + target + " plateforme_reactive.yml")

	for i:=0; i<len(config.Applications); i++ {
		if config.Applications[i].Spark.Version != ""  {
			ansibleCommands = append(ansibleCommands, "ansible-playbook -i inventories/" + target + " deploy_spark_app.yml")
		}
		if config.Applications[i].Api.Version != "" {
			ansibleCommands = append(ansibleCommands, "ansible-playbook -i inventories/" + target + " deploy_apiserver.yml")
		}
	}

	return ansibleCommands
}

func unmarshall(yamlText []byte) (*Manifest, error) {
	var config Manifest
	var err = yaml.Unmarshal(yamlText, &config)
	if err != nil {
		err_msg := fmt.Sprintf("Error when reading YAML file. Can't create Manifest Object. Yaml Error: %v\n", err)
		return nil, errors.New(err_msg)
	}

	fmt.Println("line", config.ReactPlatform.Version)
	return &config, nil
}

func unmarshallFromFile(filePath string) (*Manifest, error) {
	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	config, err := unmarshall([]byte(data))

	if err != nil {
		return nil, err
	}

	return config, nil
}

func execCommand(cmd string, wg *sync.WaitGroup) (int, error) {
	fmt.Println("Prepare to execute command : ", cmd)
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		err_msg := fmt.Sprintf("Error when running %s command. Error details: %v\n", cmd, err)
		return -1, errors.New(err_msg)
	}
	fmt.Printf("\tCommand result : \n\t%s", out)
	wg.Done()
	return 0, nil
}
