# \PrivateGatewaysApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPrivateGateway**](PrivateGatewaysApi.md#AddPrivateGateway) | **Post** /network/v1/privategateway/ | Create a Private Gateway
[**DeletePrivateGateway**](PrivateGatewaysApi.md#DeletePrivateGateway) | **Delete** /network/v1/privategateway/{name} | Delete a Private Gateway
[**DiscoverPrivateGateway**](PrivateGatewaysApi.md#DiscoverPrivateGateway) | **Get** /network/v1/privategateway/{container} | Retrieve Names of all Private Gateways and Subcontainers in a Container
[**GetPrivateGateway**](PrivateGatewaysApi.md#GetPrivateGateway) | **Get** /network/v1/privategateway/{name} | Retrieve Details of a Private Gateway
[**ListPrivateGateway**](PrivateGatewaysApi.md#ListPrivateGateway) | **Get** /network/v1/privategateway/{container}/ | Retrieve Details of all Private Gateways in a Container
[**UpdatePrivateGateway**](PrivateGatewaysApi.md#UpdatePrivateGateway) | **Put** /network/v1/privategateway/{name} | Update a Private Gateway


# **AddPrivateGateway**
> PrivateGatewayResponse AddPrivateGateway($body, $cookie)

Create a Private Gateway

This endpoint is not available on Oracle Cloud Machine.<p>Create a private gateway when you want to set up a private data connection between subnets in your premises and IP networks in your Oracle Compute Cloud Service account.<p>Create a private gateway object in Oracle Compute Cloud Service and attach your IP networks to this private gateway. Then set up a Oracle Network Cloud Service - FastConnect private peering connection. With this connection, you can access instances on your IP networks using their private IP addresses from your on-premise private networks. You don't need to associate public IP addresses with instances on IP networks. After creating the private gateway object, note down the three-part name of the object (<code>/Compute-<em>identity_domain</em>/<em>user</em>/<em>object</em></code>). You'll need to provide this name when you set up a Oracle Network Cloud Service - FastConnect private peering connection. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=cloud&id=OFCUG-GUID-BB5DFB78-C1F8-44C8-B47A-09701A605411\">Provisioning Oracle Network Cloud Service - FastConnect</a> in <i>Using Oracle Network Cloud Service - FastConnect</i>.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PrivateGatewayPostRequest**](PrivateGatewayPostRequest.md)|  | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**PrivateGatewayResponse**](PrivateGateway-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePrivateGateway**
> DeletePrivateGateway($name, $cookie)

Delete a Private Gateway

This endpoint is not available on Oracle Cloud Machine.<p>Deletes the specified private gateway.<p>Ensure that the private gateway is not being used before deleting it. When you request for Oracle Network Cloud Service - FastConnect private peering, you have to provide the name of the private gateway. Once the peering is created, you cannot change the private gateway. If you delete a private gateway that is being used in private peering,  you can no longer send traffic over this connection. If you want to modify a private gateway, use the <a class=\"xref\" href=\"op-network-v1-privategateway-{name}-put.html\">PUT /network/v1/privategateway/{name}</a> HTTP request to modify the IP networks that are associated with the private gateway. <p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverPrivateGateway**
> PrivateGatewayDiscoverResponse DiscoverPrivateGateway($container, $cookie)

Retrieve Names of all Private Gateways and Subcontainers in a Container

This endpoint is not available on Oracle Cloud Machine.<p>Retrieves the names of private gateways and subcontainers that you can access in the specified container. <p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| Specify &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;i&gt;user&lt;/i&gt;/&lt;/code&gt; to retrieve the names of objects that you can access. Specify &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;/code&gt; to retrieve the names of containers that contain objects that you can access. | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**PrivateGatewayDiscoverResponse**](PrivateGateway-discover-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+directory+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPrivateGateway**
> PrivateGatewayResponse GetPrivateGateway($name, $cookie)

Retrieve Details of a Private Gateway

This endpoint is not available on Oracle Cloud Machine.<p>Retrieves details of the specified private gateway. <p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**PrivateGatewayResponse**](PrivateGateway-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPrivateGateway**
> PrivateGatewayListResponse ListPrivateGateway($container, $cookie)

Retrieve Details of all Private Gateways in a Container

This endpoint is not available on Oracle Cloud Machine.<p>Retrieves details of all the private gateway objects that are available in the specified container. <p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| &lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;&lt;/code&gt; or &lt;p&gt;&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;&lt;/code&gt; | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**PrivateGatewayListResponse**](PrivateGateway-list-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePrivateGateway**
> PrivateGatewayResponse UpdatePrivateGateway($name, $body, $cookie)

Update a Private Gateway

This endpoint is not available on Oracle Cloud Machine.<p>You can update the IP networks that are associated with a private gateway, the description, and tags.<p>If a Oracle Network Cloud Service - FastConnect private peering has already been set up using the private gateway that you are updating, traffic immediately starts flowing to the instances on updated IP networks over the Oracle Network Cloud Service - FastConnect private peering. <p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | 
 **body** | [**PrivateGatewayPutRequest**](PrivateGatewayPutRequest.md)|  | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**PrivateGatewayResponse**](PrivateGateway-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

