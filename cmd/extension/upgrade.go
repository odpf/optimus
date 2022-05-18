package extension

import (
	"fmt"

	"github.com/odpf/salt/log"
	"github.com/spf13/cobra"

	"github.com/odpf/optimus/extension"
)

type upgradeCommand struct {
	logger log.Logger

	project              *extension.RepositoryProject
	reservedCommandNames []string
}

func newUpgradeCommand(logger log.Logger, project *extension.RepositoryProject, reservedCommandNames []string) *cobra.Command {
	upgrade := &upgradeCommand{
		logger:               logger,
		project:              project,
		reservedCommandNames: reservedCommandNames,
	}

	cmd := &cobra.Command{
		Use:   "upgrade",
		Short: "upgrade is a sub command to allow user to upgrade an extension command",
		RunE:  upgrade.RunE,
	}
	return cmd
}

func (u *upgradeCommand) RunE(cmd *cobra.Command, _ []string) error {
	verbose, _ := cmd.Flags().GetBool("verbose")

	manager, err := getExtensionManager(verbose, u.reservedCommandNames...)
	if err != nil {
		return err
	}

	u.logger.Info(fmt.Sprintf("Upgrading [%s] ...", u.project.CommandName))
	if err := manager.Upgrade(u.project.CommandName); err != nil {
		u.logger.Error("... finished with error")
		return err
	}
	u.logger.Info("... finished successfully")
	return nil
}
