package nodeProject

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("Hello, World")
}

func initWorkingDirectory(dirName string) {
	os.Mkdir(dirName, 0755)
	os.Chdir(dirName)
}

func addDevDependencies() {
	dependencies := []string{"@jest/globals", "jest", "nodemon", "ts-jest", "typescript"}
	for _, element := range dependencies {
		npmCommand := exec.Command("npm", "i", "-D", element)
		npmCommand.Output()
	}
}

func addScripts() {
	packageFile, _ := os.ReadFile("package.json")
	var packageJson map[string]interface{}
	json.Unmarshal(packageFile, &packageJson)
	scripts := map[string]string{
		"test":        "jest unit",
		"integration": "jest int",
		"dev":         "nodemon",
		"prod":        "node build/src/main.js",
		"build":       "tsc",
	}
	packageJson["scripts"] = scripts
	marshaledPackageJson, _ := json.Marshal(packageJson)
	os.WriteFile("package.json", marshaledPackageJson, 0644)
}

func setupTsConfig() {

	compileOptions := map[string]string{
		"outDir": "build/src",
	}

	exclude := []string{"**/*.test.ts"}

	tsConfigJson := make(map[string]interface{})

	tsConfigJson["compileOptions"] = compileOptions

	tsConfigJson["exclude"] = exclude

	marshaledTsConfigJson, _ := json.Marshal(tsConfigJson)

	os.WriteFile("tsconfig.json", marshaledTsConfigJson, 0644)

}

func setupGitIgnore() {
	ignored := []string{"node_modules", "build"}

	var sb strings.Builder

	for _, element := range ignored {
		sb.WriteString(element)
		sb.WriteString("\n")
	}

	os.WriteFile(".gitignore", []byte(sb.String()), 0644)
}

func setupJestConfig() {

	jestConfig := `/** @type {import('ts-jest').JestConfigWithTsJest} */
module.exports = {
    preset: 'ts-jest',
    testEnvironment: 'node',
};`
	os.WriteFile("jest.config.js", []byte(jestConfig), 0644)
}
