/*
Taikun - WebApi

Testing CatalogApiService

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

func Test_taikuncore_CatalogApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CatalogApiService CatalogAppLockManager", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		httpRes, err := apiClient.CatalogApi.CatalogAppLockManager(context.Background(), v).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogApiService CatalogAvailablePackageDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var repoName string
		var packageName string
		var v string

		resp, httpRes, err := apiClient.CatalogApi.CatalogAvailablePackageDetails(context.Background(), repoName, packageName, v).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogApiService CatalogAvailableVersions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		resp, httpRes, err := apiClient.CatalogApi.CatalogAvailableVersions(context.Background(), v).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogApiService CatalogBindProjects", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		httpRes, err := apiClient.CatalogApi.CatalogBindProjects(context.Background(), v).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogApiService CatalogCatalogAppDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32
		var v string

		resp, httpRes, err := apiClient.CatalogApi.CatalogCatalogAppDetails(context.Background(), id, v).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogApiService CatalogCatalogAppParamsDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32
		var v string

		resp, httpRes, err := apiClient.CatalogApi.CatalogCatalogAppParamsDetails(context.Background(), id, v).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogApiService CatalogCatalogAppValue", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		resp, httpRes, err := apiClient.CatalogApi.CatalogCatalogAppValue(context.Background(), v).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogApiService CatalogCatalogAppValueAutocomplete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		resp, httpRes, err := apiClient.CatalogApi.CatalogCatalogAppValueAutocomplete(context.Background(), v).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogApiService CatalogCatalogDropdown", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		resp, httpRes, err := apiClient.CatalogApi.CatalogCatalogDropdown(context.Background(), v).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogApiService CatalogCatalogList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		resp, httpRes, err := apiClient.CatalogApi.CatalogCatalogList(context.Background(), v).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogApiService CatalogCreateCatalog", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		httpRes, err := apiClient.CatalogApi.CatalogCreateCatalog(context.Background(), v).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogApiService CatalogCreateCatalogApp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		httpRes, err := apiClient.CatalogApi.CatalogCreateCatalogApp(context.Background(), v).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogApiService CatalogDeleteCatalog", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32
		var v string

		httpRes, err := apiClient.CatalogApi.CatalogDeleteCatalog(context.Background(), id, v).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogApiService CatalogDeleteCatalogApp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32
		var v string

		httpRes, err := apiClient.CatalogApi.CatalogDeleteCatalogApp(context.Background(), id, v).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogApiService CatalogEditCatalog", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		httpRes, err := apiClient.CatalogApi.CatalogEditCatalog(context.Background(), v).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogApiService CatalogEditCatalogAppParams", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		httpRes, err := apiClient.CatalogApi.CatalogEditCatalogAppParams(context.Background(), v).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogApiService CatalogEditCatalogAppVersion", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		httpRes, err := apiClient.CatalogApi.CatalogEditCatalogAppVersion(context.Background(), v).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogApiService CatalogList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		resp, httpRes, err := apiClient.CatalogApi.CatalogList(context.Background(), v).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogApiService CatalogLockManager", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		httpRes, err := apiClient.CatalogApi.CatalogLockManager(context.Background(), v).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
