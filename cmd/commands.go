package cmd

import (
	"os"

	"github.com/MakeNowJust/heredoc"
	"github.com/odpf/salt/cmdx"
	"github.com/odpf/salt/term"
	cli "github.com/spf13/cobra"

	"github.com/odpf/optimus/cmd/admin"
	"github.com/odpf/optimus/cmd/backup"
	"github.com/odpf/optimus/cmd/deploy"
	"github.com/odpf/optimus/cmd/extension"
	"github.com/odpf/optimus/cmd/initialize"
	"github.com/odpf/optimus/cmd/job"
	"github.com/odpf/optimus/cmd/logger"
	"github.com/odpf/optimus/cmd/namespace"
	"github.com/odpf/optimus/cmd/project"
	"github.com/odpf/optimus/cmd/replay"
	"github.com/odpf/optimus/cmd/resource"
	"github.com/odpf/optimus/cmd/secret"
	"github.com/odpf/optimus/cmd/serve"
	"github.com/odpf/optimus/cmd/version"
	"github.com/odpf/optimus/utils"
)

var disableColoredOut = false

// New constructs the 'root' command. It houses all other sub commands
// default output of logging should go to stdout
// interactive output like progress bars should go to stderr
// unless the stdout/err is a tty, colors/progressbar should be disabled
func New() *cli.Command {
	disableColoredOut = !utils.IsTerminal(os.Stdout)

	cmd := &cli.Command{
		Use: "optimus <command> <subcommand> [flags]",
		Long: heredoc.Doc(`
			Optimus is an easy-to-use, reliable, and performant workflow orchestrator for 
			data transformation, data modeling, pipelines, and data quality management.
		
			For passing authentication header, set one of the following environment
			variables:
			1. OPTIMUS_AUTH_BASIC_TOKEN
			2. OPTIMUS_AUTH_BEARER_TOKEN`),
		SilenceUsage: true,
		Example: heredoc.Doc(`
				$ optimus job create
				$ optimus backup create
				$ optimus backup list
				$ optimus replay create
			`),
		Annotations: map[string]string{
			"group:core": "true",
			"help:learn": heredoc.Doc(`
				Use 'optimus <command> <subcommand> --help' for more information about a command.
				Read the manual at https://odpf.github.io/optimus/
			`),
			"help:feedback": heredoc.Doc(`
				Open an issue here https://github.com/odpf/optimus/issues
			`),
		},
		PersistentPreRun: func(cmd *cli.Command, args []string) {
			// initialise color if not requested to be disabled
			cs := term.NewColorScheme()
			if !disableColoredOut {
				logger.ColoredNotice = func(s string, a ...interface{}) string {
					return cs.Yellowf(s, a...)
				}
				logger.ColoredError = func(s string, a ...interface{}) string {
					return cs.Redf(s, a...)
				}
				logger.ColoredSuccess = func(s string, a ...interface{}) string {
					return cs.Greenf(s, a...)
				}
			}
		},
	}

	cmdx.SetHelp(cmd)
	cmd.PersistentFlags().BoolVar(&disableColoredOut, "no-color", disableColoredOut, "Disable colored output")

	cmd.AddCommand(admin.NewAdminCommand(cmd))
	cmd.AddCommand(backup.NewBackupCommand(cmd))
	cmd.AddCommand(deploy.NewDeployCommand())
	cmd.AddCommand(initialize.NewInitializeCommand())
	cmd.AddCommand(job.NewJobCommand(cmd))
	cmd.AddCommand(namespace.NewNamespaceCommand())
	cmd.AddCommand(project.NewProjectCommand())
	cmd.AddCommand(replay.NewReplayCommand(cmd))
	cmd.AddCommand(resource.NewResourceCommand(cmd))
	cmd.AddCommand(secret.NewSecretCommand(cmd))
	cmd.AddCommand(version.NewVersionCommand())

	cmd.AddCommand(serve.NewServeCommand())

	extension.UpdateWithExtension(cmd)
	return cmd
}
