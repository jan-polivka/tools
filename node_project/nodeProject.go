package nodeProject

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("Hello, World")
}

func addDevDependencies() {
	fmt.Println("add dev dependencies")
	npmCommand := exec.Command("npm", "i", "-D", "jest")
	npmCommand.Output()
}
