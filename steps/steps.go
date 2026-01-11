package steps

import (
	"os/exec"
	"time"

	"github.com/buildium-org/buildium_harness/testcli"
	"github.com/buildium-org/buildium_harness/testserver"
)

func ExampleServerStep(config *testserver.ServerTestConfig) error {
	config.Logger.LogTitle("Example Server Step")
	config.Logger.LogInfo("This is an example server step")
	config.Server.Start()
	defer config.Server.Stop()
	time.Sleep(1000 * time.Millisecond)
	config.Logger.LogSuccess("Server started successfully")
	return nil
}

func ExampleCliStep(config *testcli.CliTestConfig) error {
	config.Logger.LogTitle("Example Cli Step")
	config.Logger.LogInfo("This is an example cli step")
	cmd := exec.Command(config.Executable)
	cmd.Stdout = config.Logger.Writer()
	cmd.Stderr = config.Logger.Writer()
	err := cmd.Run()
	if err != nil {
		config.Logger.LogError("Failed to run command")
		return err
	}
	return nil
}
