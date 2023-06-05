/*
Taikun - WebApi

Testing ProjectInfracostsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package taikuncore

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/chnyda/taikungoclient"
)

func Test_taikuncore_ProjectInfracostsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProjectInfracostsApiService ProjectInfracostsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		httpRes, err := apiClient.ProjectInfracostsApi.ProjectInfracostsDelete(context.Background(), v).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectInfracostsApiService ProjectInfracostsDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32
		var v string

		resp, httpRes, err := apiClient.ProjectInfracostsApi.ProjectInfracostsDetails(context.Background(), projectId, v).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectInfracostsApiService ProjectInfracostsEdit", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32
		var v string

		httpRes, err := apiClient.ProjectInfracostsApi.ProjectInfracostsEdit(context.Background(), projectId, v).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
