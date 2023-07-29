package main

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitWorkingDirectory(t *testing.T) {

	testDir := "testDirectory"

	initWorkingDirectory(testDir)
	cwd, _ := os.Getwd()
	assert.Contains(t, cwd, testDir)
	destroyTestDirectory()
}

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

func TestSetupTsCOnfig(t *testing.T) {
	setupTestDirectory()

	setupTsConfig()

	dat, _ := os.ReadFile("tsconfig.json")

	assert.True(t, strings.Contains(string(dat), "build/src"))

	destroyTestDirectory()
}

func TestSetupGitgnore(t *testing.T) {
	setupTestDirectory()

	setupGitIgnore()

	dat, _ := os.ReadFile(".gitignore")

	assert.True(t, strings.Contains(string(dat), "node_modules"))

	destroyTestDirectory()
}

func TestSetupJestConfig(t *testing.T) {
	setupTestDirectory()

	setupJestConfig()

	dat, _ := os.ReadFile("jest.config.js")

	assert.True(t, strings.Contains(string(dat), "preset"))

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
