package monitor

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/cmd/settings"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	stscobra "github.com/stackvista/stackstate-cli/internal/cobra"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
)

const LongDescription = `Edit a monitor.

The edit command allows you to directly edit any StackState Monitor. It will open
the editor defined by your VISUAL, or EDITOR environment variables, or fall back to 'vi' for Linux or 'notepad' for
Windows.
When '--unlock' is specified, the CLI will always unlock the Monitor when editing it.
This might introduce changes that prevent the originating StackPack from upgrading correctly. Any changes you make are not the responsibility of the StackPack developer.
`

const MonitorNodeType = "Monitor"

type EditArgs struct {
	ID         int64
	Identifier string
	Unlock     bool
}

func MonitorEditCommand(cli *di.Deps) *cobra.Command {
	args := &EditArgs{}
	cmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit a monitor",
		Long:  LongDescription,
		RunE:  cli.CmdRunEWithApi(RunMonitorEditCommand(args)),
	}

	common.AddIDFlagVar(cmd, &args.ID, IDFlagUsage)
	common.AddIdentifierFlagVar(cmd, &args.Identifier, IdentifierFlagUsage)
	stscobra.MarkMutexFlags(cmd, []string{common.IDFlag, common.IdentifierFlag}, "identifier", true)
	cmd.Flags().BoolVar(&args.Unlock, Unlock, false, UnlockFlagUsage)

	return cmd
}

func RunMonitorEditCommand(args *EditArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		id := args.ID

		if len(args.Identifier) > 0 {
			monitor, resp, err := api.MonitorApi.GetMonitor(cli.Context, args.Identifier).Execute()
			if err != nil {
				return common.NewResponseError(err, resp)
			}
			id = monitor.Id
		}

		export := stackstate_api.NewExport()
		export.NodesWithIds = []int64{id}

		orig, resp, err := api.ExportApi.ExportSettings(cli.Context).Export(*export).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		c, err := cli.Editor.Edit("monitor-", ".json", resp.Body)
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if strings.Compare(orig, string(c)) == 0 {
			if cli.IsJson() {
				cli.Printer.PrintJson(map[string]interface{}{"nodes": []string{}, "message": "No changes made"})
			} else {
				cli.Printer.PrintWarn("No changes made")
			}

			return nil
		}

		if args.Unlock {
			err := settings.UnlockNodes(cli, api, []int64{id}, MonitorNodeType)
			if err != nil {
				return err
			}
		}

		nodes, resp, err := api.ImportApi.ImportSettings(cli.Context).Body(string(c)).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"nodes": nodes,
			})
		} else {
			tableData := make([][]interface{}, 0)
			for _, node := range nodes {
				tableData = append(tableData, []interface{}{node["_type"], node["id"], node["status"], node["identifier"], node["name"]})
			}

			cli.Printer.Success(fmt.Sprintf("Updated <bold>%d</> monitor(s).", len(nodes)))
			if len(nodes) > 0 {
				cli.Printer.Table(printer.TableData{
					Header: []string{"Type", "Id", "Status", "Identifier", "Name"},
					Data:   tableData,
				})
			}
		}

		return nil
	}
}
