package job

import (
	"fmt"

	"github.com/odpf/salt/log"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"

	"github.com/odpf/optimus/cmd/logger"
	"github.com/odpf/optimus/cmd/survey"
	"github.com/odpf/optimus/config"
	"github.com/odpf/optimus/models"
	"github.com/odpf/optimus/store/local"
)

type addHookCommand struct {
	logger           log.Logger
	clientConfig     *config.ClientConfig
	jobSurvey        *survey.JobSurvey
	jobAddHookSurvey *survey.JobAddHookSurvey
	namespaceSurvey  *survey.NamespaceSurvey
}

// NewAddHookCommand initializes command for adding hook
func NewAddHookCommand(cliengConfig *config.ClientConfig) *cobra.Command {
	addHook := &addHookCommand{
		clientConfig: cliengConfig,
	}
	cmd := &cobra.Command{
		Use:     "addhook",
		Aliases: []string{"add_hook", "add-hook", "addHook", "attach_hook", "attach-hook", "attachHook"},
		Short:   "Attach a new Hook to existing job",
		Long:    "Add a runnable instance that will be triggered before or after the base transformation.",
		Example: "optimus addhook",
		RunE:    addHook.RunE,
		PreRunE: addHook.PreRunE,
	}
	return cmd
}

func (a *addHookCommand) PreRunE(cmd *cobra.Command, args []string) error {
	a.logger = logger.NewClientLogger(a.clientConfig.Log)
	a.jobSurvey = survey.NewJobSurvey()
	a.jobAddHookSurvey = survey.NewJobAddHookSurvey()
	a.namespaceSurvey = survey.NewNamespaceSurvey(a.logger)
	return nil
}

func (a *addHookCommand) RunE(cmd *cobra.Command, args []string) error {
	namespace, err := a.namespaceSurvey.AskToSelectNamespace(a.clientConfig)
	if err != nil {
		return err
	}

	pluginRepo := models.PluginRegistry
	jobSpecFs := afero.NewBasePathFs(afero.NewOsFs(), namespace.Job.Path)
	jobSpecRepo := local.NewJobSpecRepository(
		jobSpecFs,
		local.NewJobSpecAdapter(pluginRepo),
	)

	selectedJobName, err := a.jobSurvey.AskToSelectJobName(jobSpecRepo)
	if err != nil {
		return err
	}
	jobSpec, err := jobSpecRepo.GetByName(selectedJobName)
	if err != nil {
		return err
	}
	jobSpec, err = a.jobAddHookSurvey.AskToAddHook(jobSpec, pluginRepo)
	if err != nil {
		return err
	}
	if err := jobSpecRepo.Save(jobSpec); err != nil {
		return err
	}
	a.logger.Info(fmt.Sprintf("Hook successfully added to %s", selectedJobName))
	return nil
}
