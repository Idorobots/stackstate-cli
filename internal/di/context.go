package di

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
	sts "gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

// Depedency Injection context for the CLI
type Context struct {
	Config *config.Config
	Client *sts.APIClient
}

func NewContext() Context {
	return Context{
		Config: nil,
		Client: nil,
	}
}

func CmdRunEWithDI(cli *Context, runFn func(*Context, *context.Context, *cobra.Command, []string) error) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {

		log.Ctx(cmd.Context()).Info().Msg(fmt.Sprintf("Loaded config %+v", cli.Config))

		configuration := sts.NewConfiguration()
		configuration.Servers[0] = sts.ServerConfiguration{
			URL:         cli.Config.URL,
			Description: "",
			Variables:   nil,
		}
		client := sts.NewAPIClient(configuration)
		cli.Client = client

		auth := make(map[string]sts.APIKey)
		auth["ApiToken"] = sts.APIKey{
			Key:    cli.Config.ApiToken,
			Prefix: "",
		}
		ctx := context.WithValue(
			cmd.Context(),
			sts.ContextAPIKeys,
			auth,
		)

		return runFn(cli, &ctx, cmd, args)
	}
}
