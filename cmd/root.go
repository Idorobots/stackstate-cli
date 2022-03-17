package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/cmd/persistent_flags"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func RootCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sts",
		Short: "StackState CLI: " + Version,
	}
	cmd.AddCommand(VersionCommand(cli))
	cmd.AddCommand(ScriptCommand(cli))
	cmd.AddCommand(CliCommand(cli))
	cmd.AddCommand(MonitorCommand(cli))
	cmd.AddCommand(SettingsCommand(cli))
	persistent_flags.AddPersistentFlags(cmd)

	return cmd
}
