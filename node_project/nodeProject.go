package nodeProject

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Hello, World")
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
