# \CheckerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckerArtifact**](CheckerApi.md#CheckerArtifact) | **Post** /api/v1/checker/artifact | Check artifact url
[**CheckerAws**](CheckerApi.md#CheckerAws) | **Post** /api/v1/checker/aws | Check aws credential
[**CheckerAzure**](CheckerApi.md#CheckerAzure) | **Post** /api/v1/checker/azure | Check azure credentials
[**CheckerAzureQuota**](CheckerApi.md#CheckerAzureQuota) | **Post** /api/v1/checker/azure/quota/cpu | Check azure cpu quota
[**CheckerCidr**](CheckerApi.md#CheckerCidr) | **Post** /api/v1/checker/cidr | Check valid cidr format
[**CheckerCron**](CheckerApi.md#CheckerCron) | **Post** /api/v1/checker/cron | Check valid cron job format
[**CheckerDns**](CheckerApi.md#CheckerDns) | **Post** /api/v1/checker/dns | Check valid dns format
[**CheckerDuplicateName**](CheckerApi.md#CheckerDuplicateName) | **Post** /api/v1/checker/duplicate | Duplicate name
[**CheckerGoogle**](CheckerApi.md#CheckerGoogle) | **Post** /api/v1/checker/google | 
[**CheckerKeycloak**](CheckerApi.md#CheckerKeycloak) | **Post** /api/v1/checker/keycloak | Check keycloak credential
[**CheckerKubeConfig**](CheckerApi.md#CheckerKubeConfig) | **Post** /api/v1/checker/kube-config | 
[**CheckerNode**](CheckerApi.md#CheckerNode) | **Post** /api/v1/checker/node | Duplicate server name checker
[**CheckerNtp**](CheckerApi.md#CheckerNtp) | **Post** /api/v1/checker/ntp | Check valid ntp format
[**CheckerOpenAi**](CheckerApi.md#CheckerOpenAi) | **Post** /api/v1/checker/openai | Check open-ai token
[**CheckerOpenstack**](CheckerApi.md#CheckerOpenstack) | **Post** /api/v1/checker/openstack | Check openstack credential
[**CheckerOpenstackTaikunImage**](CheckerApi.md#CheckerOpenstackTaikunImage) | **Post** /api/v1/checker/openstack-image/{id} | Check openstack taikun image
[**CheckerOpenstackTaikunLbImage**](CheckerApi.md#CheckerOpenstackTaikunLbImage) | **Post** /api/v1/checker/taikun-lb-image/{id} | Check openstack taikun lb image
[**CheckerOrganization**](CheckerApi.md#CheckerOrganization) | **Post** /api/v1/checker/organization | Check duplicate org name
[**CheckerPrometheus**](CheckerApi.md#CheckerPrometheus) | **Post** /api/v1/checker/prometheus | Check prometheus credentials
[**CheckerProxmox**](CheckerApi.md#CheckerProxmox) | **Post** /api/v1/checker/proxmox | Check proxmox credential
[**CheckerS3**](CheckerApi.md#CheckerS3) | **Post** /api/v1/checker/s3 | Check s3 credential
[**CheckerSsh**](CheckerApi.md#CheckerSsh) | **Post** /api/v1/checker/ssh | Check valid ssh key format
[**CheckerTanzu**](CheckerApi.md#CheckerTanzu) | **Post** /api/v1/checker/tanzu | Check tanzu credential
[**CheckerUser**](CheckerApi.md#CheckerUser) | **Post** /api/v1/checker/user | Check duplicate username
[**CheckerYaml**](CheckerApi.md#CheckerYaml) | **Post** /api/v1/checker/yaml | Check yaml file



## CheckerArtifact

> CheckerArtifact(ctx).ArtifactUrlCheckerCommand(artifactUrlCheckerCommand).Execute()

Check artifact url

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    artifactUrlCheckerCommand := *openapiclient.NewArtifactUrlCheckerCommand() // ArtifactUrlCheckerCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CheckerApi.CheckerArtifact(context.Background()).ArtifactUrlCheckerCommand(artifactUrlCheckerCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckerApi.CheckerArtifact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artifactUrlCheckerCommand** | [**ArtifactUrlCheckerCommand**](ArtifactUrlCheckerCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckerAws

> CheckerAws(ctx).CheckAwsCommand(checkAwsCommand).Execute()

Check aws credential

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    checkAwsCommand := *openapiclient.NewCheckAwsCommand() // CheckAwsCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CheckerApi.CheckerAws(context.Background()).CheckAwsCommand(checkAwsCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckerApi.CheckerAws``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerAwsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkAwsCommand** | [**CheckAwsCommand**](CheckAwsCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckerAzure

> CheckerAzure(ctx).CheckAzureCommand(checkAzureCommand).Execute()

Check azure credentials

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    checkAzureCommand := *openapiclient.NewCheckAzureCommand() // CheckAzureCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CheckerApi.CheckerAzure(context.Background()).CheckAzureCommand(checkAzureCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckerApi.CheckerAzure``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerAzureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkAzureCommand** | [**CheckAzureCommand**](CheckAzureCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckerAzureQuota

> CheckerAzureQuota(ctx).CheckAzureCpuQuotaCommand(checkAzureCpuQuotaCommand).Execute()

Check azure cpu quota

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    checkAzureCpuQuotaCommand := *openapiclient.NewCheckAzureCpuQuotaCommand() // CheckAzureCpuQuotaCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CheckerApi.CheckerAzureQuota(context.Background()).CheckAzureCpuQuotaCommand(checkAzureCpuQuotaCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckerApi.CheckerAzureQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerAzureQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkAzureCpuQuotaCommand** | [**CheckAzureCpuQuotaCommand**](CheckAzureCpuQuotaCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckerCidr

> CheckerCidr(ctx).CidrCommand(cidrCommand).Execute()

Check valid cidr format

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    cidrCommand := *openapiclient.NewCidrCommand() // CidrCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CheckerApi.CheckerCidr(context.Background()).CidrCommand(cidrCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckerApi.CheckerCidr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerCidrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cidrCommand** | [**CidrCommand**](CidrCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckerCron

> CheckerCron(ctx).CronJobCommand(cronJobCommand).Execute()

Check valid cron job format

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    cronJobCommand := *openapiclient.NewCronJobCommand() // CronJobCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CheckerApi.CheckerCron(context.Background()).CronJobCommand(cronJobCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckerApi.CheckerCron``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerCronRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cronJobCommand** | [**CronJobCommand**](CronJobCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckerDns

> CheckerDns(ctx).DnsCommand(dnsCommand).Execute()

Check valid dns format

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    dnsCommand := *openapiclient.NewDnsCommand() // DnsCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CheckerApi.CheckerDns(context.Background()).DnsCommand(dnsCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckerApi.CheckerDns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerDnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnsCommand** | [**DnsCommand**](DnsCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckerDuplicateName

> bool CheckerDuplicateName(ctx).DuplicateNameCheckerCommand(duplicateNameCheckerCommand).Execute()

Duplicate name

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    duplicateNameCheckerCommand := *openapiclient.NewDuplicateNameCheckerCommand() // DuplicateNameCheckerCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckerApi.CheckerDuplicateName(context.Background()).DuplicateNameCheckerCommand(duplicateNameCheckerCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckerApi.CheckerDuplicateName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckerDuplicateName`: bool
    fmt.Fprintf(os.Stdout, "Response from `CheckerApi.CheckerDuplicateName`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerDuplicateNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **duplicateNameCheckerCommand** | [**DuplicateNameCheckerCommand**](DuplicateNameCheckerCommand.md) |  | 

### Return type

**bool**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckerGoogle

> bool CheckerGoogle(ctx).Config(config).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    config := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckerApi.CheckerGoogle(context.Background()).Config(config).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckerApi.CheckerGoogle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckerGoogle`: bool
    fmt.Fprintf(os.Stdout, "Response from `CheckerApi.CheckerGoogle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerGoogleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **config** | ***os.File** |  | 

### Return type

**bool**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckerKeycloak

> CheckerKeycloak(ctx).KeycloakCheckerCommand(keycloakCheckerCommand).Execute()

Check keycloak credential

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    keycloakCheckerCommand := *openapiclient.NewKeycloakCheckerCommand() // KeycloakCheckerCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CheckerApi.CheckerKeycloak(context.Background()).KeycloakCheckerCommand(keycloakCheckerCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckerApi.CheckerKeycloak``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerKeycloakRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keycloakCheckerCommand** | [**KeycloakCheckerCommand**](KeycloakCheckerCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckerKubeConfig

> bool CheckerKubeConfig(ctx).Config(config).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    config := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckerApi.CheckerKubeConfig(context.Background()).Config(config).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckerApi.CheckerKubeConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckerKubeConfig`: bool
    fmt.Fprintf(os.Stdout, "Response from `CheckerApi.CheckerKubeConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerKubeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **config** | ***os.File** |  | 

### Return type

**bool**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckerNode

> CheckerNode(ctx).NodeCommand(nodeCommand).Execute()

Duplicate server name checker

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    nodeCommand := *openapiclient.NewNodeCommand() // NodeCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CheckerApi.CheckerNode(context.Background()).NodeCommand(nodeCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckerApi.CheckerNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodeCommand** | [**NodeCommand**](NodeCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckerNtp

> CheckerNtp(ctx).NtpCommand(ntpCommand).Execute()

Check valid ntp format

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    ntpCommand := *openapiclient.NewNtpCommand() // NtpCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CheckerApi.CheckerNtp(context.Background()).NtpCommand(ntpCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckerApi.CheckerNtp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerNtpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ntpCommand** | [**NtpCommand**](NtpCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckerOpenAi

> CheckerOpenAi(ctx).OpenAiCheckerCommand(openAiCheckerCommand).Execute()

Check open-ai token

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    openAiCheckerCommand := *openapiclient.NewOpenAiCheckerCommand() // OpenAiCheckerCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CheckerApi.CheckerOpenAi(context.Background()).OpenAiCheckerCommand(openAiCheckerCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckerApi.CheckerOpenAi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerOpenAiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **openAiCheckerCommand** | [**OpenAiCheckerCommand**](OpenAiCheckerCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckerOpenstack

> CheckerOpenstack(ctx).CheckOpenstackCommand(checkOpenstackCommand).Execute()

Check openstack credential

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    checkOpenstackCommand := *openapiclient.NewCheckOpenstackCommand() // CheckOpenstackCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CheckerApi.CheckerOpenstack(context.Background()).CheckOpenstackCommand(checkOpenstackCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckerApi.CheckerOpenstack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerOpenstackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkOpenstackCommand** | [**CheckOpenstackCommand**](CheckOpenstackCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckerOpenstackTaikunImage

> CheckerOpenstackTaikunImage(ctx, id).Execute()

Check openstack taikun image

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CheckerApi.CheckerOpenstackTaikunImage(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckerApi.CheckerOpenstackTaikunImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckerOpenstackTaikunImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckerOpenstackTaikunLbImage

> CheckerOpenstackTaikunLbImage(ctx, id).Execute()

Check openstack taikun lb image

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CheckerApi.CheckerOpenstackTaikunLbImage(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckerApi.CheckerOpenstackTaikunLbImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckerOpenstackTaikunLbImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckerOrganization

> CheckerOrganization(ctx).OrganizationNameCheckerCommand(organizationNameCheckerCommand).Execute()

Check duplicate org name

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    organizationNameCheckerCommand := *openapiclient.NewOrganizationNameCheckerCommand() // OrganizationNameCheckerCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CheckerApi.CheckerOrganization(context.Background()).OrganizationNameCheckerCommand(organizationNameCheckerCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckerApi.CheckerOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationNameCheckerCommand** | [**OrganizationNameCheckerCommand**](OrganizationNameCheckerCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckerPrometheus

> CheckerPrometheus(ctx).CheckPrometheusCommand(checkPrometheusCommand).Execute()

Check prometheus credentials

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    checkPrometheusCommand := *openapiclient.NewCheckPrometheusCommand() // CheckPrometheusCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CheckerApi.CheckerPrometheus(context.Background()).CheckPrometheusCommand(checkPrometheusCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckerApi.CheckerPrometheus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerPrometheusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkPrometheusCommand** | [**CheckPrometheusCommand**](CheckPrometheusCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckerProxmox

> CheckerProxmox(ctx).ProxmoxCheckerCommand(proxmoxCheckerCommand).Execute()

Check proxmox credential

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    proxmoxCheckerCommand := *openapiclient.NewProxmoxCheckerCommand() // ProxmoxCheckerCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CheckerApi.CheckerProxmox(context.Background()).ProxmoxCheckerCommand(proxmoxCheckerCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckerApi.CheckerProxmox``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerProxmoxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **proxmoxCheckerCommand** | [**ProxmoxCheckerCommand**](ProxmoxCheckerCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckerS3

> CheckerS3(ctx).CheckS3Command(checkS3Command).Execute()

Check s3 credential

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    checkS3Command := *openapiclient.NewCheckS3Command() // CheckS3Command | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CheckerApi.CheckerS3(context.Background()).CheckS3Command(checkS3Command).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckerApi.CheckerS3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerS3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkS3Command** | [**CheckS3Command**](CheckS3Command.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckerSsh

> CheckerSsh(ctx).SshKeyCommand(sshKeyCommand).Execute()

Check valid ssh key format

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    sshKeyCommand := *openapiclient.NewSshKeyCommand() // SshKeyCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CheckerApi.CheckerSsh(context.Background()).SshKeyCommand(sshKeyCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckerApi.CheckerSsh``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerSshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sshKeyCommand** | [**SshKeyCommand**](SshKeyCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckerTanzu

> CheckerTanzu(ctx).CheckTanzuCommand(checkTanzuCommand).Execute()

Check tanzu credential

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    checkTanzuCommand := *openapiclient.NewCheckTanzuCommand() // CheckTanzuCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CheckerApi.CheckerTanzu(context.Background()).CheckTanzuCommand(checkTanzuCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckerApi.CheckerTanzu``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerTanzuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkTanzuCommand** | [**CheckTanzuCommand**](CheckTanzuCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckerUser

> CheckerUser(ctx).UserExistCommand(userExistCommand).Execute()

Check duplicate username

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    userExistCommand := *openapiclient.NewUserExistCommand() // UserExistCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CheckerApi.CheckerUser(context.Background()).UserExistCommand(userExistCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckerApi.CheckerUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userExistCommand** | [**UserExistCommand**](UserExistCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckerYaml

> CheckerYaml(ctx).YamlValidatorCommand(yamlValidatorCommand).Execute()

Check yaml file

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    yamlValidatorCommand := *openapiclient.NewYamlValidatorCommand() // YamlValidatorCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CheckerApi.CheckerYaml(context.Background()).YamlValidatorCommand(yamlValidatorCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckerApi.CheckerYaml``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerYamlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **yamlValidatorCommand** | [**YamlValidatorCommand**](YamlValidatorCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

