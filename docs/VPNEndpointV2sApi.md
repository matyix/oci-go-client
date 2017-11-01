# \VPNEndpointV2sApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddVPNEndpointV2**](VPNEndpointV2sApi.md#AddVPNEndpointV2) | **Post** /vpnendpoint/v2/ | Create a VPN Endpoint V2
[**DeleteVPNEndpointV2**](VPNEndpointV2sApi.md#DeleteVPNEndpointV2) | **Delete** /vpnendpoint/v2/{name} | Delete a VPN Endpoint V2
[**GetVPNEndpointV2**](VPNEndpointV2sApi.md#GetVPNEndpointV2) | **Get** /vpnendpoint/v2/{name} | Retrieve Details of a VPN Endpoint V2
[**ListVPNEndpointV2**](VPNEndpointV2sApi.md#ListVPNEndpointV2) | **Get** /vpnendpoint/v2/{container}/ | Retrieve Details of all VPN Endpoint V2s in a Container
[**UpdateVPNEndpointV2**](VPNEndpointV2sApi.md#UpdateVPNEndpointV2) | **Put** /vpnendpoint/v2/{name} | Update a VPN Endpoint V2


# **AddVPNEndpointV2**
> VpnConnectionResponse AddVPNEndpointV2($body, $cookie)

Create a VPN Endpoint V2

This endpoint is not available on Oracle Cloud Machine.<p>Creates an IPSec VPN connection from Oracle Cloud to your data centers using VPN as a Service (VPNaaS).<p><b>Note:</b> You can use VPNaaS to set up a tunnel to instances that are on IP networks. However, VPNaaS doesn't support VPN connections to instances that don't have any interface on IP networks. To establish a VPN tunnel to instances that are on the shared network, follow the steps for creating a single-homed Corente Services Gateway instance in <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=STCSG-GUID-07FAD4A9-DDE4-4391-BB9C-EA4784552EDA\">Setting Up VPN</a> in <i>Using Oracle Compute Cloud Service (IaaS)</i>.<p><b>Prerequisites</b><p>Ensure that you complete the following tasks and noted the required information before creating a VPN connection.<p>* Create an IP network or use an existing IP network. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=STCSG-GUID-9A826000-2728-4837-905A-7835FA775F9B\">Creating an IP Network</a> in <i>Using Oracle Compute Cloud Service (IaaS)</i>. Make a note of the name of this IP network.<p>* Configure a supported third-party VPN device at your data center and make a note of the public IP address of this gateway. The third-party VPN device must be ready for the VPN connection to be established. For information about certified third-party VPN device configurations, see <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=STCSG-GUID-D5186CEE-08A9-4673-BD80-91CC7646988A\">About Setting Up VPN</a> in <i>Using Oracle Compute Cloud Service (IaaS)</i>.<p>* Ensure that you have the pre-shared key (PSK) that you want to use for this VPN connection.<p>* Create a vNICset. When you create instances, specify this vNICset for each vNIC that is added to an IP network that will be reachable over the VPN connection. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=cloud&id=STCSG-GUID-565766C2-21AF-4EF2-B7C4-728B7604543A\">Creating a vNICset</a> in <i>Using Oracle Compute Cloud Service (IaaS)</i>.<p>While your VPN connection is being configured, its status is <code>PENDING</code>. It can take around 20 to 30 minutes for your VPN gateway to be created. When the cloud VPN gateway is created, the <code>localGatewayAddress</code> parameter provides its public IP address. To monitor the status of your VPN connection and retrieve the public IP address of the cloud VPN gateway, send the <a class=\"xref\" href=\"op-vpnendpoint-v2-{name}-get.html\">GET /vpnendpoint/v2/{name}</a> request. You'll have to update the third-party VPN device in your data center with the public IP address of your cloud VPN gateway. If the third-party device in your data center is configured and ready, the VPN connection is established. The value of the <code>tunnelStatus</code> parameter changes to <code>UP</code> when the connection is established.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**VpnConnectionPostRequest**](VpnConnectionPostRequest.md)|  | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**VpnConnectionResponse**](VPNConnection-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVPNEndpointV2**
> DeleteVPNEndpointV2($name, $cookie)

Delete a VPN Endpoint V2

This endpoint is not available on Oracle Cloud Machine.<p>You may want to add new vNIC sets or update other parameters of an existing VPN connection. Not all parameters of a VPN connection can be updated using the PUT /vpnendpoint/v2/{name} request. In such cases, you can delete the VPN connection, and then recreate the it with the updated parameters by using the POST /vpnendpoint/v2/ request.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


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

# **GetVPNEndpointV2**
> VpnConnectionResponse GetVPNEndpointV2($name, $cookie)

Retrieve Details of a VPN Endpoint V2

This endpoint is not available on Oracle Cloud Machine.<p>Retrieve Details of the specified VPN connection. You can retrieve details of a VPN connection to track the status of the tunnel. The tunnel can be in one of the following states:<p>* PENDING: indicates that your VPN connection is being set up.<p>* UP: indicates that your VPN connection is established.<p>* DOWN: indicates that your VPN connection is down.<p>* ERROR: indicates that your VPN connection is in the error state.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**VpnConnectionResponse**](VPNConnection-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVPNEndpointV2**
> VpnConnectionListResponse ListVPNEndpointV2($container, $name, $customerVpnGateway, $cookie)

Retrieve Details of all VPN Endpoint V2s in a Container

This endpoint is not available on Oracle Cloud Machine.<p>Retrieves details of the VPN connections in the specified container and match the specified query criteria. If you don't specify any query criteria, then details of all the VPN connections in the container are displayed. To filter the search results, you can pass one or more of the following query parameters, by appending them to the URI in the following syntax:<p><code>?parameter1=value1&ampparameter2=value2&ampparameterN=valueN</code><p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| &lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;&lt;/code&gt; or &lt;p&gt;&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;&lt;/code&gt; | 
 **name** | **string**| Specify the three-part name (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;) of the VPN connection. | [optional] 
 **customerVpnGateway** | **string**| The public IP address of the VPN gateway in your data center through which you want to connect to the Oracle Cloud VPN gateway. | [optional] 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**VpnConnectionListResponse**](VPNConnection-list-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVPNEndpointV2**
> VpnConnectionResponse UpdateVPNEndpointV2($name, $body, $cookie)

Update a VPN Endpoint V2

This endpoint is not available on Oracle Cloud Machine.<p>Updates the values for <code>psk</code> and <code>reachable_routes</code> for the specified VPN connection. If you want to update values for any other parameter, you'll have to delete the VPN connection and then re-create with the new parameters using the POST /vpnendpoint/v2/ request. Although you can only update the values for <code>psk</code> and <code>reachable_routes</code> using this request, you must specify the current values for all the existing parameters in the request body.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | 
 **body** | [**VpnConnectionPutRequest**](VpnConnectionPutRequest.md)|  | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**VpnConnectionResponse**](VPNConnection-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

