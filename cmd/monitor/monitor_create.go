package monitor

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

const (
	FileFlag = "file"
)

func CreateMonitorCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a monitor.",
		RunE:  di.CmdRunEWithDeps(cli, RunCreateMonitorCommand),
	}
	cmd.Flags().StringP(FileFlag, "f", "", "The file with the monitor in it. Can either be a YAML or JSON file.")
	cmd.MarkFlagRequired(FileFlag)

	return cmd
}

func RunCreateMonitorCommand(cli *di.Deps, cmd *cobra.Command, args []string) common.CLIError {
	file, err := cmd.Flags().GetString(FileFlag)
	if err != nil {
		return common.NewCLIError(err)
	}

	var monitor stackstate_client.CreateMonitor
	err = util.ReadInputStructFromFile(file, &monitor)
	if err != nil {
		return common.NewCLIError(err)
	}

	_, resp, err := cli.Client.MonitorApi.CreateMonitor(cli.Context).CreateMonitor(monitor).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	cli.Printer.Success("Monitor created.")
	return nil
}
