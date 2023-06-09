/*
Taikun - WebApi

Testing UserGroupApiService

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

func Test_taikuncore_UserGroupApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UserGroupApiService UsergroupsBindProjectsGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.UserGroupApi.UsergroupsBindProjectsGroup(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserGroupApiService UsergroupsCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserGroupApi.UsergroupsCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserGroupApiService UsergroupsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.UserGroupApi.UsergroupsDelete(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserGroupApiService UsergroupsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserGroupApi.UsergroupsList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserGroupApiService UsergroupsListByProjectGroupId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectGroupId int32

		resp, httpRes, err := apiClient.UserGroupApi.UsergroupsListByProjectGroupId(context.Background(), projectGroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserGroupApiService UsergroupsUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userGroupId int32

		httpRes, err := apiClient.UserGroupApi.UsergroupsUpdate(context.Background(), userGroupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserGroupApiService UsergroupsUserGroupUsers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userGroupId int32

		resp, httpRes, err := apiClient.UserGroupApi.UsergroupsUserGroupUsers(context.Background(), userGroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
