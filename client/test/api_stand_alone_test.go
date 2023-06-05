/*
Taikun - WebApi

Testing StandAloneApiService

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

func Test_taikuncore_StandAloneApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test StandAloneApiService StandAloneCommit", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		httpRes, err := apiClient.StandAloneApi.StandAloneCommit(context.Background(), v).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StandAloneApiService StandAloneCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		resp, httpRes, err := apiClient.StandAloneApi.StandAloneCreate(context.Background(), v).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StandAloneApiService StandAloneDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		httpRes, err := apiClient.StandAloneApi.StandAloneDelete(context.Background(), v).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StandAloneApiService StandAloneDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32
		var v string

		resp, httpRes, err := apiClient.StandAloneApi.StandAloneDetails(context.Background(), projectId, v).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StandAloneApiService StandAloneIpManagement", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		httpRes, err := apiClient.StandAloneApi.StandAloneIpManagement(context.Background(), v).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StandAloneApiService StandAloneList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		resp, httpRes, err := apiClient.StandAloneApi.StandAloneList(context.Background(), v).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StandAloneApiService StandAloneListForPoller", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		resp, httpRes, err := apiClient.StandAloneApi.StandAloneListForPoller(context.Background(), v).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StandAloneApiService StandAloneProjectDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32
		var v string

		resp, httpRes, err := apiClient.StandAloneApi.StandAloneProjectDetails(context.Background(), projectId, v).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StandAloneApiService StandAlonePurge", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		httpRes, err := apiClient.StandAloneApi.StandAlonePurge(context.Background(), v).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StandAloneApiService StandAloneRepair", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		httpRes, err := apiClient.StandAloneApi.StandAloneRepair(context.Background(), v).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StandAloneApiService StandAloneReset", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		httpRes, err := apiClient.StandAloneApi.StandAloneReset(context.Background(), v).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StandAloneApiService StandAloneUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		httpRes, err := apiClient.StandAloneApi.StandAloneUpdate(context.Background(), v).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StandAloneApiService StandAloneUpdateFlavor", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		httpRes, err := apiClient.StandAloneApi.StandAloneUpdateFlavor(context.Background(), v).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}