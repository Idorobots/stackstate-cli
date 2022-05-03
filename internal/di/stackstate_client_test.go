package di

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	stackstate_client "gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
)

func setupClient() (StackStateClient, *stackstate_client.ServerApiMock) {
	serverApi := stackstate_client.NewServerApiMock()
	api := stackstate_client.APIClient{
		ServerApi: &serverApi,
	}
	client := NewStackStateClient(&api, context.Background())
	return client, &serverApi
}

func TestConnectError(t *testing.T) {
	client, serverApi := setupClient()
	err := fmt.Errorf("authentication error")
	resp := http.Response{StatusCode: 401}
	serverApi.ServerInfoResponse.Response = &resp
	serverApi.ServerInfoResponse.Error = err

	_, _, actualErr := client.Connect()

	assert.Equal(t, common.NewConnectError(err, &resp), actualErr)
}

func TestConnectSuccess(t *testing.T) {
	client, serverApi := setupClient()
	serverApi.ServerInfoResponse.Result = stackstate_client.ServerInfo{DeploymentMode: "test"}

	_, serverInfo, _ := client.Connect()

	assert.Equal(t, "test", serverInfo.DeploymentMode)
}
