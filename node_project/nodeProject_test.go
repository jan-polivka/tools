package nodeProject

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddDevDependencies(t *testing.T) {

	t.Skip("Takes too long")

	setupTestDirectory()
	addDevDependencies()

	dat, _ := os.ReadFile("package.json")

	assert.NotNil(t, dat)
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
