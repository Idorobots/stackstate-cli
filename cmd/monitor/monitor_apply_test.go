package monitor

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

//nolint:deadcode,unused
func setMonitorApplyCmd() (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps()
	cmd := MonitorApplyCommand(&cli.Deps)
	return &cli, cmd
}

// PLEASE WRITE TESTS!
