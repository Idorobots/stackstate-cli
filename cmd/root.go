package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/hierynomus/taipan"

	home "github.com/mitchellh/go-homedir"

	color "github.com/logrusorgru/aurora/v3"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
)

const (
	EnvPrefix = "STS"
)

func RootCommand() *cobra.Command {
	var verbosity int

	cmd := &cobra.Command{
		Use:   "sts-toolbox",
		Short: "Development CLI for StackState",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			switch verbosity {
			case 0: //nolint:gomnd
				// Nothing to do
			case 1: //nolint:gomnd
				zerolog.SetGlobalLevel(zerolog.InfoLevel)
			case 2: //nolint:gomnd
				zerolog.SetGlobalLevel(zerolog.DebugLevel)
			default:
				zerolog.SetGlobalLevel(zerolog.TraceLevel)
			}

			return nil
		},
	}

	cmd.PersistentFlags().CountVarP(&verbosity, "verbose", "v", "Print more verbose logging")

	return cmd
}

func AllCommands(cfg *config.Config) *cobra.Command {
	cmd := RootCommand()
	cmd.AddCommand(VersionCommand())
	cmd.AddCommand(ValuesCommand(cfg))
	cmd.AddCommand(SandboxCommand(cfg))
	cmd.AddCommand(ClusterCommand())
	cmd.AddCommand(ScaleCommand(cfg))
	cmd.AddCommand(TrafficMirrorCommand(cfg))
	cmd.AddCommand(DevCommand(cfg))

	return cmd
}

func Execute(ctx context.Context) {
	cfg := &config.Config{}
	cmd := AllCommands(cfg)

	homeFolder, err := home.Expand("~/.sts")
	if err != nil {
		fmt.Printf("%s", color.Red(err))
		os.Exit(1)
	}

	zerolog.SetGlobalLevel(zerolog.ErrorLevel)

	taipanConfig := &taipan.Config{
		DefaultConfigName:  "sts-toolbox",
		ConfigurationPaths: []string{".", homeFolder},
		EnvironmentPrefix:  EnvPrefix,
		AddConfigFlag:      true,
		ConfigObject:       cfg,
		PrefixCommands:     true,
	}

	t := taipan.New(taipanConfig)
	t.Inject(cmd)

	if err := cmd.ExecuteContext(ctx); err != nil {
		fmt.Printf("🎃 %s\n", color.Red(err))
		os.Exit(1)
	}
}
