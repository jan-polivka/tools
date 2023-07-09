package nodeProject

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddDevDependencies(t *testing.T) {

	t.Skip("Takes too long")

	setupTestDirectory()
	addDevDependencies()

	dat, _ := os.ReadFile("package.json")

	assert.True(t, strings.Contains(string(dat), "jest"))
	destroyTestDirectory()
}

func TestAddScripts(t *testing.T) {

	setupTestDirectory()

	dummyMap := map[string]string{"dummy": "dummy"}

	marshaledDummy, _ := json.Marshal(dummyMap)

	os.WriteFile("package.json", marshaledDummy, 0644)

	addScripts()

	dat, _ := os.ReadFile("package.json")

	var packageJson map[string]interface{}

	json.Unmarshal(dat, &packageJson)

	assert.NotNil(t, packageJson["scripts"])
	scriptsJson := (packageJson["scripts"]).(map[string]interface{})
	assert.NotNil(t, scriptsJson["build"].(string))
	assert.True(t, strings.Contains(string(dat), "prod"))

	destroyTestDirectory()
}

func setupTestDirectory() {
	testDirectory := "testDirectory"
	os.Mkdir(testDirectory, 0755)
	os.Chdir(testDirectory)
}

func destroyTestDirectory() {
	testDirectory := "testDirectory"
	os.Chdir("..")
	os.RemoveAll(testDirectory)
}
