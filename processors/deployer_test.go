package processors

import (
	"testing"
	"github.com/ZacharyGroff/GitHooks/config"
	"os"
	"strings"
)

func TestDeployScript(t *testing.T) {
	dir, _ := os.Getwd()
	testPath := dir + "/test.sh"
	config := config.Config{ScriptPath: testPath}
	deployer := Deployer{&config}
	actual := deployer.Deploy()

	expected := "test\n"
	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}
}
