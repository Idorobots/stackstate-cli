package script

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

func ScriptExecuteCommand(cli *di.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "execute <script>",
		Args:  cobra.ExactArgs(1),
		Short: "Execute a STSL script.",
		RunE:  di.CmdRunEWithCLIContext(cli, RunScriptExecuteCommand),
	}

	cmd.Flags().StringP("arguments-script", "a", "", "A script that returns a java.util.Map with arguments that can be used as variables within the actual script.")
	cmd.Flags().IntP("timeout", "t", 0, "Timeout in milli-seconds.")

	return cmd
}

func RunScriptExecuteCommand(cli *di.Context, cmd *cobra.Command, args []string) error {
	var argumentsScript *string
	a, _ := cmd.Flags().GetString("arguments-script")
	if a != "" {
		argumentsScript = &a
	}
	var timeoutMs *int32
	t, _ := cmd.Flags().GetInt("timeout")
	if t != 0 {
		t32 := int32(t)
		timeoutMs = &t32
	}
	verbose, _ := cmd.Flags().GetCount("verbose")

	auth := make(map[string]stackstate_client.APIKey)
	auth["ApiToken"] = stackstate_client.APIKey{
		Key:    cli.Config.ApiToken,
		Prefix: "",
	}
	ctx := context.WithValue(
		cmd.Context(),
		stackstate_client.ContextAPIKeys,
		auth,
	)

	scriptExecute := cli.Client.ScriptingApi.ScriptExecute(ctx)
	scriptRequest := stackstate_client.ExecuteScriptRequest{
		TimeoutMs:       timeoutMs,
		Script:          args[0],
		ArgumentsScript: argumentsScript,
	}

	if verbose > 0 {
		j, _ := json.Marshal(scriptRequest)
		log.
			Ctx(cmd.Context()).
			Info().
			RawJSON("ExecuteScriptRequest", j).
			Msg("Executing script request")
	}

	scriptResponse, resp, err := scriptExecute.ExecuteScriptRequest(scriptRequest).Execute()

	if err != nil {
		var status string
		if resp != nil {
			status = resp.Status + ". "
		}

		switch v := err.(type) {
		case stackstate_client.GenericOpenAPIError:
			return fmt.Errorf("%vError response: %+v", status, string(v.Body()))
		default:
			return fmt.Errorf("%v%+v", status, v)
		}
	}

	cmd.Println(scriptResponse["result"])
	return nil
}
