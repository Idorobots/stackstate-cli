package monitor

import (
	"testing"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

//nolint:deadcode,unused
func setMonitorListCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := MonitorListCommand(&cli.Deps)
	return &cli, cmd
}

// PLEASE WRITE TESTS!
