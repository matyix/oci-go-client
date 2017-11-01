# \SnapshotsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSnapshot**](SnapshotsApi.md#AddSnapshot) | **Post** /snapshot/ | Create an Instance Snapshot
[**DeleteSnapshot**](SnapshotsApi.md#DeleteSnapshot) | **Delete** /snapshot/{name} | Delete an Instance Snapshot
[**DiscoverRootSnapshot**](SnapshotsApi.md#DiscoverRootSnapshot) | **Get** /snapshot/ | Retrieve Names of Containers
[**DiscoverSnapshot**](SnapshotsApi.md#DiscoverSnapshot) | **Get** /snapshot/{container} | Retrieve Names of all Instance Snapshots and Subcontainers in a Container
[**GetSnapshot**](SnapshotsApi.md#GetSnapshot) | **Get** /snapshot/{name} | Retrieve Details of an Instance Snapshot
[**ListSnapshot**](SnapshotsApi.md#ListSnapshot) | **Get** /snapshot/{container}/ | Retrieve Details of all Instance Snapshots in a Container


# **AddSnapshot**
> SnapshotResponse AddSnapshot($body, $cookie)

Create an Instance Snapshot

Creates a snapshot request, which in turn creates a machine image to preserve all the changes made in the instance since launch.<p>There can be only one snapshot request in the <code>active</code> or <code>queued</code> state for an instance at any given time. After the request state changes to <code>error</code> or <code>complete</code>, you can issue another snapshot request for that instance.<p>This command returns a snapshot request identifier, which you can use to check the progress of the asynchronous snapshot request by sending a <code>GET</code> request with the snapshot request identifier. When the state of the request changes to complete, you can send the <code>GET /machineimage/{name}</code> request to verify that the machine image is created correctly.<p>If you use the <code>\"delay\": \"shutdown\"</code> option, the snapshot remains in the <code>active</code> state, until you shutdown the instance by deleting it. When the delete instance request <code>(DELETE /instance/{name})</code> is executed, a snapshot of the instance is taken and a machine image is created, the state of the snapshot changes to <code>complete</code>, and then the instance is stopped and deleted.<p>Add this new machine image to an image list, and then use it to create a new instance.<p><b>Prerequisite:</b> Ensure that you have selected a replication policy for your Oracle Storage Cloud Service instance. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=CSSTO-GUID-5D53C11F-3D9E-43E4-8D1D-DDBB95DEC715\">Selecting a Replication Policy for Your Service Instance</a> in <em>Using Oracle Storage Cloud Service</em>.</p><p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SnapshotPostRequest**](SnapshotPostRequest.md)|  | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**SnapshotResponse**](Snapshot-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSnapshot**
> DeleteSnapshot($name, $cookie)

Delete an Instance Snapshot

Deletes a snapshot request. Deleting the snapshot request does not delete the machine image that was created by it. No response is returned for the delete action.<p>The following restrictions apply to deleting a snapshot request:<ul><li>Requests in the <code>active</code> state cannot be deleted.</li><li>Requests can only be deleted when they are in the <code>error</code> or <code>complete</code> state.</li></ul><p><b>Prerequisite:</b> Ensure that you have selected a replication policy for your Oracle Storage Cloud Service instance. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=CSSTO-GUID-5D53C11F-3D9E-43E4-8D1D-DDBB95DEC715\">Selecting a Replication Policy for Your Service Instance</a> in <em>Using Oracle Storage Cloud Service</em>.</p><p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Multipart name of the instance snapshot request. | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverRootSnapshot**
> SnapshotDiscoverResponse DiscoverRootSnapshot($cookie)

Retrieve Names of Containers

Retrieves the names of containers that contain objects that you can access. You can use this information to construct the multipart name of an object.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**SnapshotDiscoverResponse**](Snapshot-discover-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+directory+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverSnapshot**
> SnapshotDiscoverResponse DiscoverSnapshot($container, $cookie)

Retrieve Names of all Instance Snapshots and Subcontainers in a Container

Retrieves the names of objects and subcontainers that you can access in the specified container.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| Specify &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;i&gt;user&lt;/i&gt;/&lt;/code&gt; to retrieve the names of objects that you can access. Specify &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;/code&gt; to retrieve the names of containers that contain objects that you can access. | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**SnapshotDiscoverResponse**](Snapshot-discover-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+directory+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSnapshot**
> SnapshotResponse GetSnapshot($name, $cookie)

Retrieve Details of an Instance Snapshot

Retrieves details for a specific snapshot request.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Multipart name of the instance snapshot request. | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**SnapshotResponse**](Snapshot-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSnapshot**
> SnapshotListResponse ListSnapshot($container, $account, $instance, $machineimage, $name, $cookie)

Retrieve Details of all Instance Snapshots in a Container

Retrieves details of the instance snapshot requests that are available in the specified container and match the specified query criteria. If you don't specify any query criteria, then details of all the instance snapshot requests in the container are displayed. To filter the search results, you can pass one or more of the following query parameters, by appending them to the URI in the following syntax:<p><code>?parameter1=value1&ampparameter2=value2&ampparameterN=valueN</code><p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| &lt;p&gt;&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;&lt;/code&gt; or &lt;p&gt;&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;&lt;/code&gt; | 
 **account** | **string**| Link to object storage account that will store the image. | [optional] 
 **instance** | **string**| Multipart name of the instance. | [optional] 
 **machineimage** | **string**| Three-part name (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt; ) of the machine image generated from the instance snapshot request. | [optional] 
 **name** | **string**| Multipart name of the instance snapshot request. | [optional] 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**SnapshotListResponse**](Snapshot-list-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

