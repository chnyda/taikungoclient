/*
Taikun - WebApi

Testing SlackApiService

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

func Test_taikuncore_SlackApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SlackApiService SlackCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlackApi.SlackCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlackApiService SlackDeleteMultiple", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.SlackApi.SlackDeleteMultiple(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlackApiService SlackDropdown", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlackApi.SlackDropdown(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlackApiService SlackList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlackApi.SlackList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlackApiService SlackUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		httpRes, err := apiClient.SlackApi.SlackUpdate(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlackApiService SlackVerify", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.SlackApi.SlackVerify(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
