# \IPAddressReservationsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddIPAddressReservation**](IPAddressReservationsApi.md#AddIPAddressReservation) | **Post** /network/v1/ipreservation/ | Create an IP Address Reservation for IP Networks
[**DeleteIPAddressReservation**](IPAddressReservationsApi.md#DeleteIPAddressReservation) | **Delete** /network/v1/ipreservation/{name} | Delete an IP Address Reservation Used in IP Networks
[**GetIPAddressReservation**](IPAddressReservationsApi.md#GetIPAddressReservation) | **Get** /network/v1/ipreservation/{name} | Retrieve Details of an IP Address Reservation Used in IP Networks
[**ListIPAddressReservation**](IPAddressReservationsApi.md#ListIPAddressReservation) | **Get** /network/v1/ipreservation/{container}/ | Retrieve Details of all IP Address Reservations in a Container
[**UpdateIPAddressReservation**](IPAddressReservationsApi.md#UpdateIPAddressReservation) | **Put** /network/v1/ipreservation/{name} | Update an IP Address Reservation Used in IP Networks


# **AddIPAddressReservation**
> IpAddressReservationResponse AddIPAddressReservation($body, $cookie)

Create an IP Address Reservation for IP Networks

Reserves a NAT IPv4 address, which you can associate with one or more virtual NICs for routing traffic outside an IP network or an IP network exchange using NAT.<p>To reserve an IP address for an instance that you have created in the flat network, see Create an IP Reservation.<p>Reserve an IP address from one of the following IP pools:<p>* A pool of public IP addresses. An IP address from this pool is accessible over the public Internet.<p>* A pool of cloud IP addresses. An IP address from this pool is accessible to other IP networks in the Oracle cloud. You can use these IP addresses to communicate with other Oracle services.<p>When you reserve an IP address from a specified IP pool, an IPv4 address is allocated for your use.<p>A public IP address or a cloud IP address can be associated with only one vNIC at a time. However, a single vNIC can have a maximum of two NAT IP addresses, one from each IP pool.<p>After reserving an IP address, you can and associate the IP address with a vNIC on your instance. See <a class=\"xref\" href=\"op-network-v1-ipassociation--post.html\">Create an IP Address Association</a>.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IpAddressReservationPostRequest**](IpAddressReservationPostRequest.md)|  | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**IpAddressReservationResponse**](IpAddressReservation-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIPAddressReservation**
> DeleteIPAddressReservation($name, $cookie)

Delete an IP Address Reservation Used in IP Networks

Deletes the specified IP address reservation.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


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

# **GetIPAddressReservation**
> IpAddressReservationResponse GetIPAddressReservation($name, $cookie)

Retrieve Details of an IP Address Reservation Used in IP Networks

Retrieves details of the specified IP address reservation.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**IpAddressReservationResponse**](IpAddressReservation-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIPAddressReservation**
> IpAddressReservationListResponse ListIPAddressReservation($container, $cookie)

Retrieve Details of all IP Address Reservations in a Container

Retrieves details of the IP address reservations that are available in the specified container. This request is for IP address reservations used in IP networks.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| &lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;&lt;/code&gt; or &lt;p&gt;&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;&lt;/code&gt; | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**IpAddressReservationListResponse**](IpAddressReservation-list-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateIPAddressReservation**
> IpAddressReservationResponse UpdateIPAddressReservation($name, $body, $cookie)

Update an IP Address Reservation Used in IP Networks

Updates the specified IP address reservation.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | 
 **body** | [**IpAddressReservationPutRequest**](IpAddressReservationPutRequest.md)|  | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**IpAddressReservationResponse**](IpAddressReservation-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

