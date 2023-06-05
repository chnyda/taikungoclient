/*
Taikun - WebApi

Testing PartnerApiService

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

func Test_taikuncore_PartnerApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PartnerApiService PartnerAddWhiteListDomain", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		httpRes, err := apiClient.PartnerApi.PartnerAddWhiteListDomain(context.Background(), v).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PartnerApiService PartnerBecomePartner", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		httpRes, err := apiClient.PartnerApi.PartnerBecomePartner(context.Background(), v).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PartnerApiService PartnerBindOrganizations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		httpRes, err := apiClient.PartnerApi.PartnerBindOrganizations(context.Background(), v).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PartnerApiService PartnerContactUs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		httpRes, err := apiClient.PartnerApi.PartnerContactUs(context.Background(), v).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PartnerApiService PartnerCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		httpRes, err := apiClient.PartnerApi.PartnerCreate(context.Background(), v).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PartnerApiService PartnerDeleteWhiteListDomain", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		httpRes, err := apiClient.PartnerApi.PartnerDeleteWhiteListDomain(context.Background(), v).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PartnerApiService PartnerDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		resp, httpRes, err := apiClient.PartnerApi.PartnerDetails(context.Background(), v).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PartnerApiService PartnerList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		resp, httpRes, err := apiClient.PartnerApi.PartnerList(context.Background(), v).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PartnerApiService PartnerPartnerInfoRegistration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		resp, httpRes, err := apiClient.PartnerApi.PartnerPartnerInfoRegistration(context.Background(), v).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PartnerApiService PartnerPartnerList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var v string

		resp, httpRes, err := apiClient.PartnerApi.PartnerPartnerList(context.Background(), v).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PartnerApiService PartnerUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32
		var v string

		httpRes, err := apiClient.PartnerApi.PartnerUpdate(context.Background(), id, v).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
