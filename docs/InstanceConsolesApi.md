# \InstanceConsolesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInstanceConsole**](InstanceConsolesApi.md#GetInstanceConsole) | **Get** /instanceconsole/{name} | Retrieve Details of an Instance Console


# **GetInstanceConsole**
> InstanceConsoleResponse GetInstanceConsole($name, $cookie)

Retrieve Details of an Instance Console

Retrieves the messages that appear when an instance boots. Use these messages to diagnose unresponsive instances and failures in the boot up process.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Name of the instance in one of the following format: &lt;code&gt;/Compute-&lt;em&gt;identityDomain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;id&lt;/em&gt;&lt;/code&gt; or &lt;code&gt;/Compute-&lt;em&gt;identityDomain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;provided_name&lt;/em&gt;/&lt;em&gt;id&lt;/em&gt;&lt;/code&gt;, where id is an autogenerated ID. | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**InstanceConsoleResponse**](InstanceConsole-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

