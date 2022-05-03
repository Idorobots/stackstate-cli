package conf

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

/*
This config is used throughout the CLI.

Note: when updating this struct, please also update the:
 1. Conf struct
 2. Constants (e.g. MinimumRequiredEnvVars)
 2. Bindings (i.e. flags, environment variables and YAML)
 3. Validations
 4. and last but not least, the tests!
*/
type Conf struct {
	URL      string
	ApiToken string
	ApiPath  string
}

const (
	XDGConfigSubPath       = "stackstate-cli"
	ViperConfigName        = "config"
	ViperConfigType        = "yaml"
	MinimumRequiredEnvVars = "STS_CLI_URL, STS_CLI_API_TOKEN"
	MinimumRequiredFlags   = "url, api-token"
)

//nolint:golint,errcheck
func bind(cmd *cobra.Command, vp *viper.Viper) Conf {
	// defaults
	vp.SetDefault("api-path", "/api")

	// bind environment variables
	vp.BindEnv("url", "STS_CLI_URL")
	vp.BindEnv("api-token", "STS_CLI_API_TOKEN")
	vp.BindEnv("api-path", "STS_CLI_API_PATH")

	// bind flags
	vp.BindPFlag("url", cmd.Flags().Lookup("url"))
	vp.BindPFlag("api-token", cmd.Flags().Lookup("api-token"))
	vp.BindPFlag("api-path", cmd.Flags().Lookup("api-path"))

	// bind YAML
	return Conf{
		URL:      vp.GetString("url"),
		ApiToken: vp.GetString("api-token"),
		ApiPath:  vp.GetString("api-path"),
	}
}

func validate(conf Conf, errors *[]error) {
	if conf.URL == "" {
		*errors = append(*errors, MissingFieldError{FieldName: "url"})
	}

	if conf.ApiToken == "" {
		*errors = append(*errors, MissingFieldError{FieldName: "api-token"})
	}
}

func convertConfToYaml(conf Conf) string {
	return fmt.Sprintf(""+
		`# URL of your StackState instance
url: %s
# Your API token: visit https://<your stackstate url>#/cli to get your token
api-token: %s
# API path (defaults to /api)
api-path: %s
`, conf.URL, conf.ApiToken, conf.ApiPath)
}
