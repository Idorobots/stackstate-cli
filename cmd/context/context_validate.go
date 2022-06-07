package context

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

type ValidateArgs struct {
	Name string
}

func ValidateCommand(cli *di.Deps) *cobra.Command {
	args := &ValidateArgs{}
	cmd := &cobra.Command{
		Use:   "validate",
		Short: "validate context",
		Long:  "Validate a context.",
		RunE:  cli.CmdRunE(RunValidateCommand(args)),
	}

	common.AddNameFlagVar(cmd, &args.Name, "name of the context")

	return cmd
}

func RunValidateCommand(args *ValidateArgs) func(cli *di.Deps, cmd *cobra.Command) common.CLIError {
	return func(cli *di.Deps, cmd *cobra.Command) common.CLIError {
		ctxName := cli.StsConfig.CurrentContext
		ctx := cli.CurrentContext
		if args.Name != "" {
			ctx = cli.StsConfig.GetContext(args.Name).Context
			ctxName = args.Name
		}

		if err := ValidateContext(cli, cmd, ctx); err != nil {
			return err
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"context": ctxName,
				"valid":   true,
			})
		} else {
			cli.Printer.Successf("Context %s is valid.\n", ctxName)
		}

		return nil
	}
}
