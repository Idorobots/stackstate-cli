package settings

import (
	"io/ioutil"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

func setupCommandExport() (di.MockDeps, *cobra.Command) {
	mockCli := di.NewMockDeps()
	cmd := SettingsExportCommand(&mockCli.Deps)

	return mockCli, cmd
}

func TestSettingsExportPrintsToTable(t *testing.T) {
	expectedStr := `{"nodes": [{ "description": "description-1", "id": -214, "name": "name-1", "ownedBy": "urn:stackpack:common", 
"parameters": [{ "name": "name-param", "type": "LONG"}], "script": { "scriptBody": "script-bdy-1"}}]}`
	cli, cmd := setupCommandExport()
	cli.MockClient.ApiMocks.ExportApi.ExportSettingsResponse.Result = expectedStr

	_, err := util.ExecuteCommandWithContext(cli.Context, cmd, "--ids", "-214")
	assert.Nil(t, err)
	assert.Equal(t, []string{expectedStr}, *cli.MockPrinter.PrintLnCalls)
}

func TestSettingsExportIdsPrintsToTable(t *testing.T) {
	expectedStr := `{"nodes": [{ "description": "description-1", "id": -214, "name": "name-1", "ownedBy": "urn:stackpack:common", 
"parameters": [{ "name": "name-param", "type": "LONG"}], "script": { "scriptBody": "script-bdy-1"}},
{ "description": "description-1", "id": 314, "name": "name-1", "ownedBy": "urn:stackpack:common", "parameters": 
[{ "name": "name-param", "type": "LONG"}], "script": { "scriptBody": "script-bdy-1"}}]}`
	cli, cmd := setupCommandExport()
	cli.MockClient.ApiMocks.ExportApi.ExportSettingsResponse.Result = expectedStr

	_, err := util.ExecuteCommandWithContext(cli.Context, cmd, "--ids", "-214", "--ids", "314")
	assert.Nil(t, err)
	assert.Equal(t, []string{expectedStr}, *cli.MockPrinter.PrintLnCalls)
}

func TestSettingsExportMutuallyExclusiveFlags(t *testing.T) {
	cli, cmd := setupCommandExport()
	_, err := util.ExecuteCommandWithContext(cli.Context, cmd)

	assert.Equal(t, common.NewMutuallyExclusiveFlagsRequiredError([]string{Ids, Namespace, TypeName}), err)

	_, err = util.ExecuteCommandWithContext(cli.Context, cmd, "--ids", "-214", "--namespace", "default")
	assert.Equal(t, common.NewMutuallyExclusiveFlagsMultipleError([]string{Ids, Namespace, TypeName}, []string{Ids, Namespace}), err)
}

func TestRunSettingsExportWithReferencePrintToTable(t *testing.T) {
	expectedStr := `{"nodes": [{ "description": "description-1", "id": -214, 
"identifier": "urn:stackpack:common:baseline-function:median-absolute-deviation", "name": "name-1", 
"ownedBy": "urn:stackpack:common", "parameters": [{ "name": "name-param", "type": "LONG"}],
"script": { "scriptBody": "script-bdy-1"}}]}`
	cli, cmd := setupCommandExport()
	cli.MockClient.ApiMocks.ExportApi.ExportSettingsResponse.Result = expectedStr

	_, err := util.ExecuteCommandWithContext(cli.Context, cmd, "--namespace", "default",
		"--allowed-namespace-refs", "urn:stackpack:common:baseline-function:median-absolute-deviation", "--allowed-namespace-refs", "urn:stackpack")
	assert.Nil(t, err)
	assert.Equal(t, []string{expectedStr}, *cli.MockPrinter.PrintLnCalls)
}

func TestRunSettingsExportToFile(t *testing.T) {
	filePath := "./result.sjt"
	expectedStr := `{"nodes": [{ "description": "description-1", "id": -214, "description": "description-1", "name": "name-1",
"ownedBy": "urn:stackpack:common", "parameters": [{ "name": "name-param", "type": "LONG"}], 
"script": { "scriptBody": "script-bdy-1"}}]}`
	cli, cmd := setupCommandExport()
	cli.MockClient.ApiMocks.ExportApi.ExportSettingsResponse.Result = expectedStr
	_, err := util.ExecuteCommandWithContext(cli.Context, cmd, "--ids", "-214", "--file", filePath)
	assert.Nil(t, err)
	body, err := ioutil.ReadFile(filePath)
	assert.Nil(t, err)
	assert.Equal(t, expectedStr, string(body))
}

func TestRunSettingsExportTypesPrintsToTable(t *testing.T) {
	expectedStr := `{"nodes": [{ "description": "Base line function", "id": -214, "name": "name-1", 
"ownedBy": "urn:stackpack:common", "parameters": [{ "name": "name-param", "type": "BaselineFunction"}], 
"script": { "scriptBody": "script-bdy-1"}},{ "description": "Check function", "id": 314, "name": "name 314", 
"ownedBy": "urn:stackpack:common", "parameters": [{ "name": "name-param", "type": "CheckFunction"}], 
"script": { "scriptBody": "script-bdy-1"}}]}`
	cli, cmd := setupCommandExport()
	cli.MockClient.ApiMocks.ExportApi.ExportSettingsResponse.Result = expectedStr

	_, err := util.ExecuteCommandWithContext(cli.Context, cmd, "--type", "BaselineFunction", "--type", "CheckFunction")
	assert.Nil(t, err)
	assert.Equal(t, []string{expectedStr}, *cli.MockPrinter.PrintLnCalls)
}
