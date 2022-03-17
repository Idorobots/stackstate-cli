package script

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	sts "gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

const (
	ScriptFlag          = "script"
	FileFlag            = "file"
	TimeoutFlag         = "timeout"
	ArgumentsScriptFlag = "arguments-script"
)

func ScriptExecuteCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "execute {-s SCRIPT | -f FILE}",
		Short: "execute an STSL script",
		Example: "# execute a script from file\n" +
			"sts execute --file \"path/to/my.script\"\n" +
			"\n" +
			"# execute a script with variables provided by an arguments-script\n" +
			"sts execute --script \"x+y\" --arguments-script \"[x: 1, y: 2]\"",
		RunE: di.CmdRunEWithDeps(cli, RunScriptExecuteCommand),
	}

	cmd.Flags().StringP(ScriptFlag, "s", "", "the actual script to execute")
	cmd.Flags().StringP(ArgumentsScriptFlag, "a", "", "an extra script that generates arguments to be used as variables when the main script is executed, return format: java.util.Map")
	cmd.Flags().IntP(TimeoutFlag, "t", 0, "timeout in milli-seconds for script execution")
	cmd.Flags().StringP(FileFlag, "f", "", "path to a file that contains the script to execute")

	return cmd
}

func RunScriptExecuteCommand(cli *di.Deps, cmd *cobra.Command, args []string) common.CLIError {
	var script string

	script, err := cmd.Flags().GetString(ScriptFlag)
	if err != nil {
		return common.NewCLIError(err)
	}
	file, err := cmd.Flags().GetString(FileFlag)
	if err != nil {
		return common.NewCLIError(err)
	}

	if file != "" && script != "" {
		return common.NewCLIArgParseError(fmt.Errorf("can not load script both from the \"%s\" and the \"%s\" flags. Pick one or the other", ScriptFlag, FileFlag))
	}

	if file != "" {
		b, err := os.ReadFile(file)
		if err != nil {
			return common.NewCLIError(err)
		}
		script = string(b)
	}

	if script == "" {
		return common.NewCLIArgParseError(fmt.Errorf("required flag \"%s\" or \"%s\" not set", ScriptFlag, FileFlag))
	}

	var argumentsScript *string
	a, _ := cmd.Flags().GetString(ArgumentsScriptFlag)
	if a != "" {
		argumentsScript = &a
	}
	var timeoutMs *int32
	t, _ := cmd.Flags().GetInt(TimeoutFlag)
	if t != 0 {
		t32 := int32(t)
		timeoutMs = &t32
	}

	// execute script
	scriptRequest := sts.ExecuteScriptRequest{
		TimeoutMs:       timeoutMs,
		Script:          script,
		ArgumentsScript: argumentsScript,
	}

	scriptResponse, resp, err := cli.Client.ScriptingApi.
		ScriptExecute(cli.Context).
		ExecuteScriptRequest(scriptRequest).
		Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	// print response
	cli.Printer.PrintStruct(scriptResponse.Result["value"])
	return nil
}
