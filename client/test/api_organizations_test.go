/*
Taikun - WebApi

Testing OrganizationsApiService

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

func Test_taikuncore_OrganizationsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrganizationsApiService OrganizationsAcceptOffer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.OrganizationsApi.OrganizationsAcceptOffer(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrganizationsApiService OrganizationsCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OrganizationsApi.OrganizationsCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrganizationsApiService OrganizationsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		httpRes, err := apiClient.OrganizationsApi.OrganizationsDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrganizationsApiService OrganizationsDetawils", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OrganizationsApi.OrganizationsDetawils(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrganizationsApiService OrganizationsExportCsv", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.OrganizationsApi.OrganizationsExportCsv(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrganizationsApiService OrganizationsLeave", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OrganizationsApi.OrganizationsLeave(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrganizationsApiService OrganizationsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OrganizationsApi.OrganizationsList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrganizationsApiService OrganizationsOrganizationList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OrganizationsApi.OrganizationsOrganizationList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrganizationsApiService OrganizationsToggle", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.OrganizationsApi.OrganizationsToggle(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrganizationsApiService OrganizationsUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.OrganizationsApi.OrganizationsUpdate(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrganizationsApiService OrganizationsUpdatePayment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.OrganizationsApi.OrganizationsUpdatePayment(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
