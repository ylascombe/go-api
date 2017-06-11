package utils

import (
	"errors"
	"io/ioutil"
	"log"
	"github.com/ylascombe/go-api/models"
	"gopkg.in/yaml.v2"
	"fmt"
)

func unmarshall(yamlText []byte) (*models.Manifest, error) {
	var config models.Manifest
	var err = yaml.Unmarshal(yamlText, &config)
	if err != nil {
		err_msg := fmt.Sprintf("Error when reading YAML file. Can't create Manifest Object. Yaml Error: %v\n", err)
		return nil, errors.New(err_msg)
	}

	return &config, nil
}

func UnmarshallFromFile(filePath string) (*models.Manifest, error) {
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

func Marshall(in interface{}) string {
	d, err := yaml.Marshal(in)
	result := string(d)
	if err != nil {
		err_msg := fmt.Sprintf("Error when marshalling object ", in, err)
		fmt.Println(err_msg)
		//return nil, errors.New(err_msg)
	}
	return string(result)
}
