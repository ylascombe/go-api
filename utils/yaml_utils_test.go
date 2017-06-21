package utils

import (
	"github.com/stretchr/testify/assert"
	"arc-api/models"
	"testing"
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
      artifact_name: spark-colis-360
      extra_vars:
        var3: value3
        var4: value4
    api:
      version: 1.4.3
      artifact_name: api-colis-360
      extra_vars:
        var5: value5
        var6: value6
  `

const FILE_CONTENT = `---
reactive_platform:
  version: "3.1.2"
  extra_vars:
    var1: value1
    var2: value2
  features_status:
    spark1: present
    spark2: absent
...`

func TestParseManifest(t *testing.T) {
	manifest, err := unmarshall([]byte(YAML))
	assert.Nil(t, err)
	assert.Equal(t, "0.1", manifest.FormatVersion)
	assert.Equal(t, 1, len(manifest.Applications))
	//assert.Equal(t, "1.4.2", len(manifest.Applications[0].Spark.Version))
}

func TestParseManifestReactivePlatform(t *testing.T) {
	manifest, err := unmarshall([]byte(YAML))
	assert.Nil(t, err)
	assert.Equal(t, "3.1.2", manifest.ReactPlatform.Version)
	assert.Equal(t, 2, len(manifest.ReactPlatform.ExtraVars))
	assert.Equal(t, "value1", manifest.ReactPlatform.ExtraVars["var1"])
	assert.Equal(t, "value2", manifest.ReactPlatform.ExtraVars["var2"])

	assert.Equal(t, 2, len(manifest.ReactPlatform.FeaturesStatus))
	assert.Equal(t, "present", manifest.ReactPlatform.FeaturesStatus["spark1"])
	assert.Equal(t, "absent", manifest.ReactPlatform.FeaturesStatus["spark2"])
}

func TestParseManifestApplications(t *testing.T) {
	manifest, err := unmarshall([]byte(YAML))
	assert.Nil(t, err)
	assert.Equal(t, 1, len(manifest.Applications))
	assert.Equal(t, "colis360", manifest.Applications[0].Name)
	assert.Equal(t, "1.4.2", manifest.Applications[0].Spark.Version)
	assert.Equal(t, "spark-colis-360", manifest.Applications[0].Spark.ArtifactName)

	assert.Equal(t, "1.4.3", manifest.Applications[0].Api.Version)
	assert.Equal(t, "api-colis-360", manifest.Applications[0].Api.ArtifactName)

	// TODO
	//assert.Equal(t, Spark{Version:"1.4.2",ExtraVars:{"var4":"value4", "var3":"value3"}}, manifest.Applications[0].Spark)

	assert.Equal(t, 2, len(manifest.Applications[0].Spark.ExtraVars))
	assert.Equal(t, "value3", manifest.Applications[0].Spark.ExtraVars["var3"])
	assert.Equal(t, "value4", manifest.Applications[0].Spark.ExtraVars["var4"])

	assert.Equal(t, 2, len(manifest.Applications[0].Api.ExtraVars))
	assert.Equal(t, "value5", manifest.Applications[0].Api.ExtraVars["var5"])
	assert.Equal(t, "value6", manifest.Applications[0].Api.ExtraVars["var6"])
}

func TestUnmarshallFromFile(t *testing.T) {
	manifest, err := UnmarshallFromFile("test/basic_yaml_file.yml")
	assert.Nil(t, err)

	assert.Equal(t, "3.1.2", manifest.ReactPlatform.Version)
	assert.Equal(t, 2, len(manifest.ReactPlatform.ExtraVars))
	assert.Equal(t, "value1", manifest.ReactPlatform.ExtraVars["var1"])
	assert.Equal(t, "value2", manifest.ReactPlatform.ExtraVars["var2"])
	assert.Equal(t, 2, len(manifest.ReactPlatform.FeaturesStatus))
	assert.Equal(t, "present", manifest.ReactPlatform.FeaturesStatus["spark1"])
	assert.Equal(t, "absent", manifest.ReactPlatform.FeaturesStatus["spark2"])
}

func TestMarshall(t *testing.T) {

	expectYaml := `version: "1.0"
artifact_name: artifactTest
extra_vars: {}
`
	api := models.Api{Version: "1.0", ArtifactName: "artifactTest"}

	result := Marshall(api)
	assert.NotNil(t, result)
	assert.Equal(t, expectYaml, result)
}

func TestMarshallList(t *testing.T) {

	expectYaml := `apis:
- version: "1.0"
  artifact_name: artifactTest1
  extra_vars: {}
- version: "2.0"
  artifact_name: artifactTest2
  extra_vars: {}
`

	api1 := models.Api{Version: "1.0", ArtifactName: "artifactTest1"}
	api2 := models.Api{Version: "2.0", ArtifactName: "artifactTest2"}
	apiArray := []models.Api{api1, api2}

	apis := models.ApiList{List: apiArray}
	result := Marshall(apis)
	assert.NotNil(t, result)
	assert.Equal(t, expectYaml, result)
}
