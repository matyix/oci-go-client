# \OSSContainersApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOSSContainer**](OSSContainersApi.md#AddOSSContainer) | **Post** /integrations/osscontainer/ | Create an Oracle Storage Cloud Service Container
[**DeleteOSSContainer**](OSSContainersApi.md#DeleteOSSContainer) | **Delete** /integrations/osscontainer/{name} | Delete an Oracle Storage Cloud Service Container
[**DiscoverOSSContainer**](OSSContainersApi.md#DiscoverOSSContainer) | **Get** /integrations/osscontainer/{container} | Retrieve Names of all Oracle Storage Cloud Service Containers and Subcontainers in a Container
[**DiscoverRootOSSContainer**](OSSContainersApi.md#DiscoverRootOSSContainer) | **Get** /integrations/osscontainer/ | Retrieve Names of Containers
[**GetOSSContainer**](OSSContainersApi.md#GetOSSContainer) | **Get** /integrations/osscontainer/{name} | Retrieve Details of an Oracle Storage Cloud Service Container
[**ListOSSContainer**](OSSContainersApi.md#ListOSSContainer) | **Get** /integrations/osscontainer/{container}/ | Retrieve Details of all Oracle Storage Cloud Service Containers in a Container
[**UpdateOSSContainer**](OSSContainersApi.md#UpdateOSSContainer) | **Put** /integrations/osscontainer/{name} | Update an Oracle Storage Cloud Service Container


# **AddOSSContainer**
> OssContainerResponse AddOSSContainer($body, $cookie)

Create an Oracle Storage Cloud Service Container

Creates a container in Oracle Storage Cloud Service.<p>After creating a container in Oracle Storage Cloud Service, you can fetch its details. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=cloud&amp;id=CSSTO-GUID-1A8B51E1-B214-49C6-B506-BC6098739985\">Managing Containers in Oracle Storage Cloud Service</a> in <i>Using Oracle Storage Cloud Service</i>.<p><b>Prerequisite:</b> Ensure that you have selected a replication policy for your Oracle Storage Cloud Service instance. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=CSSTO-GUID-5D53C11F-3D9E-43E4-8D1D-DDBB95DEC715\">Selecting a Replication Policy for Your Service Instance</a> in <em>Using Oracle Storage Cloud Service</em>.</p><p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OssContainerPostRequest**](OssContainerPostRequest.md)|  | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**OssContainerResponse**](OSSContainer-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOSSContainer**
> DeleteOSSContainer($name, $cookie)

Delete an Oracle Storage Cloud Service Container

Deletes the specified Oracle Storage Cloud Service container. When <code>delete_remote</code> attribute of the OSS container object is set to <code>true</code>, the associated Oracle Storage Cloud Service container is deleted along with all its objects when you delete the local OSS container object. When set to <code>false</code>, only the local OSS container object is deleted. To retrieve the value of the delete_remote attribute of the OSS container, use the <a class=\"xref\" href=\"op-integrations-osscontainer-{name}-get.html\">GET /integrations/osscontainer/{name}</a> method. To change the value of the <code>delete_remote</code> attribute of the OSS container object, use the <a class=\"xref\" href=\"op-integrations-osscontainer-{name}-put.html\">PUT /integrations/osscontainer/{name}</a> method. When <code>delete_remote</code> is set to <code>true</code>, all the objects in the remote container are deleted in batches of 200. The remote container is deleted when it is empty, and then the local OSS container object is deleted. When you send this command, the state of the OSS container object changes to <code>deleting</code>. The time taken to delete the object varies depending on the number of objects in the remote container.<p><b>Prerequisite:</b> Ensure that you have selected a replication policy for your Oracle Storage Cloud Service instance. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=CSSTO-GUID-5D53C11F-3D9E-43E4-8D1D-DDBB95DEC715\">Selecting a Replication Policy for Your Service Instance</a> in <em>Using Oracle Storage Cloud Service</em>.</p><p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


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

# **DiscoverOSSContainer**
> OssContainerDiscoverResponse DiscoverOSSContainer($container, $cookie)

Retrieve Names of all Oracle Storage Cloud Service Containers and Subcontainers in a Container

Retrieves names of all Oracle Storage Cloud Service containers. <p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| Specify &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;i&gt;user&lt;/i&gt;/&lt;/code&gt; to retrieve the names of objects that you can access. Specify &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;/code&gt; to retrieve the names of containers that contain objects that you can access. Note: The path parameter &lt;code&gt;container&lt;/code&gt; does not refer to containers in Oracle Storage Cloud Service. It refers to containers in Oracle Compute Cloud Service, such as &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;i&gt;user&lt;/i&gt;/&lt;/code&gt; and &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;/code&gt;. | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**OssContainerDiscoverResponse**](OSSContainer-discover-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+directory+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverRootOSSContainer**
> OssContainerDiscoverResponse DiscoverRootOSSContainer($cookie)

Retrieve Names of Containers

Retrieves the names of containers that contain objects that you can access. You can use this information to construct the multipart name of an object.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**OssContainerDiscoverResponse**](OSSContainer-discover-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+directory+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOSSContainer**
> OssContainerResponse GetOSSContainer($name, $cookie)

Retrieve Details of an Oracle Storage Cloud Service Container

Retrieves details of the specified Oracle Storage Cloud Service container. <p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**OssContainerResponse**](OSSContainer-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOSSContainer**
> OssContainerListResponse ListOSSContainer($container, $name, $cookie)

Retrieve Details of all Oracle Storage Cloud Service Containers in a Container

Retrieves details of the all the Oracle Storage Cloud Service containers that you can access in the specified Oracle Compute Cloud Service container, such as <code>/Compute-<em>identity_domain</em>/<em>user</em></code> or <code>/Compute-<em>identity_domain</em>/</code>.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| &lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;&lt;/code&gt; or &lt;p&gt;&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;&lt;/code&gt; | 
 **name** | **string**| Multipart name of the object. | [optional] 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**OssContainerListResponse**](OSSContainer-list-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOSSContainer**
> OssContainerResponse UpdateOSSContainer($name, $body, $cookie)

Update an Oracle Storage Cloud Service Container

You can modify only the value of the <code>delete_remote</code> attribute of the OSS container. The <code>delete_remote</code> attribute defines whether the remote storage container and its objects will be deleted when the local OSS container object is deleted. Updates the specified local OSS container object in Oracle Compute Cloud Service. This request does not impact the remote Oracle Storage Cloud Service container.<p><b>Prerequisite:</b> Ensure that you have selected a replication policy for your Oracle Storage Cloud Service instance. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=CSSTO-GUID-5D53C11F-3D9E-43E4-8D1D-DDBB95DEC715\">Selecting a Replication Policy for Your Service Instance</a> in <em>Using Oracle Storage Cloud Service</em>.</p><p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | 
 **body** | [**OssContainerPutRequest**](OssContainerPutRequest.md)|  | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**OssContainerResponse**](OSSContainer-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

