package kubernetes

import (
	"io"

	"github.com/SAP/jenkins-library/pkg/command"
	"github.com/SAP/jenkins-library/pkg/piperutils"

	"github.com/SAP/jenkins-library/pkg/log"
)

// DeployUtils interface
type DeployUtils interface {
	SetEnv(env []string)
	Stdout(out io.Writer)
	Stderr(err io.Writer)
	RunExecutable(e string, p ...string) error

	piperutils.FileUtils
}

// deployUtilsBundle struct  for utils
type deployUtilsBundle struct {
	*command.Command
	*piperutils.Files
}

// NewDeployUtilsBundle initialize using deployUtilsBundle struct
func NewDeployUtilsBundle() DeployUtils {
	utils := deployUtilsBundle{
		Command: &command.Command{
			ErrorCategoryMapping: map[string][]string{
				log.ErrorConfiguration.String(): {
					"Error: Get * no such host",
					"Error: path * not found",
					"Error: rendered manifests contain a resource that already exists.",
					"Error: unknown flag",
					"Error: UPGRADE FAILED: * failed to replace object: * is invalid",
					"Error: UPGRADE FAILED: * failed to create resource: * is invalid",
					"Error: UPGRADE FAILED: an error occurred * not found",
					"Error: UPGRADE FAILED: query: failed to query with labels:",
					"Invalid value: \"\": field is immutable",
				},
				log.ErrorCustom.String(): {
					"Error: release * failed, * timed out waiting for the condition",
				},
			},
		},
		Files: &piperutils.Files{},
	}
	// reroute stderr output to logging framework, stdout will be used for command interactions
	utils.Stderr(log.Writer())
	return &utils
}
