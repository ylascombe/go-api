package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"strings"
)

const YAML = `format_version: 0.1

reactive_platform:
  version: 3.1.2
  extra_vars:
    var1: value1
    var2: value2
  features_status:
    spark1: present
    spark2: absent

applications:
  - name: colis360
    spark:
      version: 1.4.2
      extra_vars:
        var3: value3
        var4: value4
    api:
      version: 1.4.3
      extra_vars:
        var5: value5
        var6: value6
  `

const FILE_CONTENT = `---
reactive_platform:
  version: 3.1.2
  extra_vars:
    var1: value1
    var2: value2
  features_status:
    spark1: present
    spark2: absent
...`

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

	assert.Equal(t, 2, len(manifest.ReactPlatform.FeaturesStatus))
	assert.Equal(t, "present", manifest.ReactPlatform.FeaturesStatus["spark1"])
	assert.Equal(t, "absent", manifest.ReactPlatform.FeaturesStatus["spark2"])
}

func TestParseManifestApplications(t *testing.T) {
	manifest := unmarshall([]byte(YAML))
	assert.Equal(t, 1, len(manifest.Applications))
	assert.Equal(t, "colis360", manifest.Applications[0].Name)

	// TODO
	//assert.Equal(t, Spark{Version:"1.4.2",ExtraVars:{"var4":"value4", "var3":"value3"}}, manifest.Applications[0].Spark)

	assert.Equal(t, 2, len(manifest.Applications[0].Spark.ExtraVars))
	assert.Equal(t, "value3", manifest.Applications[0].Spark.ExtraVars["var3"])
	assert.Equal(t, "value4", manifest.Applications[0].Spark.ExtraVars["var4"])

	assert.Equal(t, 2, len(manifest.Applications[0].Api.ExtraVars))
	assert.Equal(t, "value5", manifest.Applications[0].Api.ExtraVars["var5"])
	assert.Equal(t, "value6", manifest.Applications[0].Api.ExtraVars["var6"])
}

func TestReadFile(t *testing.T) {
	fileContent := readFile("test/basic_yaml_file.yml")

	count := len(fileContent)
	str := strings.Trim(string(fileContent[:count]),"\x00")
	assert.Equal(t, FILE_CONTENT, str)
}

