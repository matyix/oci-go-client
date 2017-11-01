# \InstancesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteInstance**](InstancesApi.md#DeleteInstance) | **Delete** /instance/{name} | Delete an Instance
[**DiscoverInstance**](InstancesApi.md#DiscoverInstance) | **Get** /instance/{container} | Retrieve Names of all Instances and Subcontainers in a Container
[**DiscoverRootInstance**](InstancesApi.md#DiscoverRootInstance) | **Get** /instance/ | Retrieve Names of Containers
[**GetInstance**](InstancesApi.md#GetInstance) | **Get** /instance/{name} | Retrieve Details of an Instance
[**ListInstance**](InstancesApi.md#ListInstance) | **Get** /instance/{container}/ | Retrieve Details of all Instances in a Container
[**UpdateInstance**](InstancesApi.md#UpdateInstance) | **Put** /instance/{name} | Update an Instance


# **DeleteInstance**
> DeleteInstance($name, $cookie)

Delete an Instance

Shuts down an instance and removes it permanently from the system. No response is returned.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| &lt;p&gt;Multipart name of the object. | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverInstance**
> InstanceDiscoverResponse DiscoverInstance($container, $cookie)

Retrieve Names of all Instances and Subcontainers in a Container

Retrieves the names of objects and subcontainers that you can access in the specified container.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| Specify &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;i&gt;user&lt;/i&gt;/&lt;/code&gt; to retrieve the names of objects that you can access. Specify &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;/code&gt; to retrieve the names of containers that contain objects that you can access. | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**InstanceDiscoverResponse**](Instance-discover-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+directory+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverRootInstance**
> InstanceDiscoverResponse DiscoverRootInstance($cookie)

Retrieve Names of Containers

Retrieves the names of containers that contain objects that you can access. You can use this information to construct the multipart name of an object.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**InstanceDiscoverResponse**](Instance-discover-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+directory+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInstance**
> InstanceResponse GetInstance($name, $cookie)

Retrieve Details of an Instance

Retrieves details of the specified instance.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| &lt;p&gt;Multipart name of the object. | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**InstanceResponse**](Instance-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListInstance**
> InstanceListResponse ListInstance($container, $availabilityDomain, $tags, $cookie)

Retrieve Details of all Instances in a Container

Retrieves details of the instances that are in the specified container and match the specified query criteria. If you don't specify any query criteria, then details of all the instances in the container are displayed. To filter the search results, you can pass one or more of the following query parameters, by appending them to the URI in the following syntax:<p><code>?parameter1=value1&ampparameter2=value2&ampparameterN=valueN</code><p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| &lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;&lt;/code&gt; or &lt;p&gt;&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;&lt;/code&gt; | 
 **availabilityDomain** | **string**| The availability domain the instance is in | [optional] 
 **tags** | [**[]string**](string.md)| Strings used to tag the instance. When you specify tags, only instances tagged with the specified value are displayed. | [optional] 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**InstanceListResponse**](Instance-list-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateInstance**
> InstanceResponse UpdateInstance($name, $body, $cookie)

Update an Instance

You can use this request to shut down and restart individual instances which use a persistent bootable storage volume. <p>* Set the value of the <code>desired_state</code> parameter as <code>shutdown</code> to shut down the instance.<p>* Set the value of the <code>desired_state</code> parameter as <code>running</code> to restart an instance that you had previously shutdown. The instance is restarted without losing any of the instance data or configuration.<p>You can also update the value for the <code>tags</code> parameter. If the instance is already shut down, you can also change the value of the <code>shape</code> parameter.<p>Shutting down an instance is useful when you've created multiple instances in a single orchestration. In this case, stopping the instance orchestration would cause all instances to be deleted. If you want to stop one or more instances, while letting other instances in the same orchestration run, you can shut down the required instances individually.<p>Here's what happens when you shut down an instance:<p>* The instance ID is retained and reused when you restart the instance. So the multipart instance name doesn't change. This is useful in case the instance name is referenced by other objects, such as storage attachments.<p>* For instances created using orchestrations v1, the instance orchestration shows an error. However, even if the HA policy specified is active, the instance isn't automatically re-created.<p>* The resources associated with that instance, such as storage volumes and IP reservations, are freed up and can be used by other instances if required. However, if you attempt to restart an instance, ensure that the required resources are available, otherwise the instance can't restart and will go into an error state.<p>* The private IP address on the shared network is released. If you restart the instance later, it is allotted a private IP address afresh. So the private IP address of the instance on the shared network is likely to change.<p>* Dynamically allocated IP addresses on IP networks are also released. So if you start the instance later, dynamically allocated IP addresses on IP networks are also likely to change. Static private IP addresses that are allocated to interfaces in the instance orchestration won't change.<p>* Any changes that you'd made to the instance in Oracle Compute Cloud Service after the instance was created will be lost. For example, if you added the instance to security lists, attached storage volumes to the instance, or detached and attached an IP reservation, you'll need to make those changes again. The instance will be restarted with the resources that are associated with it in the instance orchestration.<p><b>Note:</b> Changes made to the instance by logging in to the instance won't be lost, however, as these are preserved on the persistent storage volumes attached to the instance. Data on storage volumes isn't affected by shutting down an instance.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Multipart name of the instance that you want to update. | 
 **body** | [**InstancePutRequest**](InstancePutRequest.md)|  | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**InstanceResponse**](Instance-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

