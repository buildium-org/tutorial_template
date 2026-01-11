package tutorialtemplate

import (
	"fmt"

	"tutorial_template/steps"

	"github.com/buildium-org/buildium_harness/meta"
	"github.com/buildium-org/buildium_harness/testcli"
	"github.com/buildium-org/buildium_harness/testserver"
)

var cliSteps = []func(config *testcli.CliTestConfig) error{
	steps.ExampleCliStep,
}

var serverSteps = []func(config *testserver.ServerTestConfig) error{
	steps.ExampleServerStep,
}

func main() {
	meta := meta.NewMeta()
	fmt.Println("Running up to stage: ", meta.Stage)

	testcli.RunCliTest(cliSteps)
	testserver.RunServerTest(serverSteps)
}
