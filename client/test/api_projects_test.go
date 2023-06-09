/*
Taikun - WebApi

Testing ProjectsApiService

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

func Test_taikuncore_ProjectsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProjectsApiService ProjectsAiAnalyzer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.ProjectsApi.ProjectsAiAnalyzer(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService ProjectsChatCompletions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProjectsApi.ProjectsChatCompletions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService ProjectsCommit", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		httpRes, err := apiClient.ProjectsApi.ProjectsCommit(context.Background(), projectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService ProjectsCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProjectsApi.ProjectsCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService ProjectsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ProjectsApi.ProjectsDelete(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService ProjectsDeleteWholeProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ProjectsApi.ProjectsDeleteWholeProject(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService ProjectsDescribe", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.ProjectsApi.ProjectsDescribe(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService ProjectsDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.ProjectsApi.ProjectsDetails(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService ProjectsDropdown", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProjectsApi.ProjectsDropdown(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService ProjectsEdit", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		httpRes, err := apiClient.ProjectsApi.ProjectsEdit(context.Background(), projectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService ProjectsEditHealth", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ProjectsApi.ProjectsEditHealth(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService ProjectsEditStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ProjectsApi.ProjectsEditStatus(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService ProjectsExtendLifetime", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ProjectsApi.ProjectsExtendLifetime(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService ProjectsForAlerting", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProjectsApi.ProjectsForAlerting(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService ProjectsForBilling", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProjectsApi.ProjectsForBilling(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService ProjectsForPoller", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProjectsApi.ProjectsForPoller(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService ProjectsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProjectsApi.ProjectsList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService ProjectsLockManager", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ProjectsApi.ProjectsLockManager(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService ProjectsLokiLogs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ProjectsApi.ProjectsLokiLogs(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService ProjectsMonitoring", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ProjectsApi.ProjectsMonitoring(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService ProjectsMonitoringAlerts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProjectsApi.ProjectsMonitoringAlerts(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService ProjectsPrometheusMetrics", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProjectsApi.ProjectsPrometheusMetrics(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService ProjectsPurge", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ProjectsApi.ProjectsPurge(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService ProjectsPurgeWholeProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ProjectsApi.ProjectsPurgeWholeProject(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService ProjectsRepair", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		httpRes, err := apiClient.ProjectsApi.ProjectsRepair(context.Background(), projectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService ProjectsToggleFullSpot", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ProjectsApi.ProjectsToggleFullSpot(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService ProjectsToggleSpotVms", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ProjectsApi.ProjectsToggleSpotVms(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService ProjectsToggleSpotWorkers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ProjectsApi.ProjectsToggleSpotWorkers(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService ProjectsUpgrade", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		httpRes, err := apiClient.ProjectsApi.ProjectsUpgrade(context.Background(), projectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsApiService ProjectsVisibility", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.ProjectsApi.ProjectsVisibility(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
