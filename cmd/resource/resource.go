package resource

import (
	"github.com/spf13/cobra"

	"github.com/odpf/optimus/config"
)

type resourceCommand struct {
	configFilePath string
	clientConfig   *config.ClientConfig

	rootCommand *cobra.Command
}

// NewResourceCommand initializes command for resource
func NewResourceCommand(rootCmd *cobra.Command) *cobra.Command {
	resource := &resourceCommand{
		clientConfig: &config.ClientConfig{},
		rootCommand:  rootCmd,
	}

	cmd := &cobra.Command{
		Use:   "resource",
		Short: "Interact with data resource",
		Annotations: map[string]string{
			"group:core": "true",
		},
		PersistentPreRunE: resource.PersistentPreRunE,
	}
	cmd.PersistentFlags().StringVarP(&resource.configFilePath, "config", "c", resource.configFilePath, "File path for client configuration")

	cmd.AddCommand(NewCreateCommand(resource.clientConfig))
	return cmd
}

func (r *resourceCommand) PersistentPreRunE(cmd *cobra.Command, args []string) error {
	r.rootCommand.PersistentPreRun(cmd, args)

	// TODO: find a way to load the config in one place
	c, err := config.LoadClientConfig(r.configFilePath, cmd.Flags())
	if err != nil {
		return err
	}
	*r.clientConfig = *c
	return nil
}
