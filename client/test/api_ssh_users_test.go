/*
Taikun - WebApi

Testing SshUsersApiService

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

func Test_taikuncore_SshUsersApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SshUsersApiService SshUsersCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		resp, httpRes, err := apiClient.SshUsersApi.SshUsersCreate(context.Background(), v).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SshUsersApiService SshUsersDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		httpRes, err := apiClient.SshUsersApi.SshUsersDelete(context.Background(), v).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SshUsersApiService SshUsersEdit", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		httpRes, err := apiClient.SshUsersApi.SshUsersEdit(context.Background(), v).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SshUsersApiService SshUsersList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accessProfileId int32
		var v string

		resp, httpRes, err := apiClient.SshUsersApi.SshUsersList(context.Background(), accessProfileId, v).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
