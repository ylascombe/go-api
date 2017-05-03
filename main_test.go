package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

const YAML = `format_version: 0.1

reactive_platform:
  version: 3.1.2
  extra_vars:
    var1: value1
    var2: value2

applications:
  - name: colis360
    spark:
      version: 1.4.2
      extra_vars: []
    api:
     version: 1.4.3
     extra_vars: toto`

func TestParseManifest(t *testing.T) {
	manifest := unmarshall([]byte(YAML))
	assert.Equal(t, "0.1", manifest.FormatVersion)
	assert.Equal(t, 1, len(manifest.Applications))
	//assert.Equal(t, "1.4.2", len(manifest.Applications[0].Spark.Version))
}

func TestParseManifestReactivePlatform(t *testing.T) {
	manifest := unmarshall([]byte(YAML))
	assert.Equal(t, "3.1.2", manifest.ReactPlatform.Version)
	assert.Equal(t, 2, len(manifest.ReactPlatform.ExtraVars))
	assert.Equal(t, "value1", manifest.ReactPlatform.ExtraVars["var1"])
	assert.Equal(t, "value2", manifest.ReactPlatform.ExtraVars["var2"])
}

func TestParseManifestApplications(t *testing.T) {
	manifest := unmarshall([]byte(YAML))
	assert.Equal(t, 1, len(manifest.Applications))
	assert.Equal(t, "colis360", manifest.Applications[0].Name)
	assert.Equal(t, Spark{}, manifest.Applications[0].Spark)
}