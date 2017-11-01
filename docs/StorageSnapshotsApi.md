# \StorageSnapshotsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddStorageSnapshot**](StorageSnapshotsApi.md#AddStorageSnapshot) | **Post** /storage/snapshot/ | Create a Storage Volume Snapshot
[**DeleteStorageSnapshot**](StorageSnapshotsApi.md#DeleteStorageSnapshot) | **Delete** /storage/snapshot/{name} | Delete a Storage Volume Snapshot
[**DiscoverRootStorageSnapshot**](StorageSnapshotsApi.md#DiscoverRootStorageSnapshot) | **Get** /storage/snapshot/ | Retrieve Names of Containers
[**DiscoverStorageSnapshot**](StorageSnapshotsApi.md#DiscoverStorageSnapshot) | **Get** /storage/snapshot/{container} | Retrieve Names of all Storage Volume Snapshots and Subcontainers in a Container
[**GetStorageSnapshot**](StorageSnapshotsApi.md#GetStorageSnapshot) | **Get** /storage/snapshot/{name} | Retrieve Details of a Storage Volume Snapshot
[**ListStorageSnapshot**](StorageSnapshotsApi.md#ListStorageSnapshot) | **Get** /storage/snapshot/{container}/ | Retrieve Details of all Storage Volume Snapshots in a Container


# **AddStorageSnapshot**
> StorageSnapshotResponse AddStorageSnapshot($body, $cookie)

Create a Storage Volume Snapshot

Creates a storage volume snapshot. Creating a storage volume snapshot enables you to capture the current state of the storage volume. You can retain snapshots as a backup, or use them to create new, identical storage volumes.<p>You can create a snapshot of a storage volume either when it is attached to an instance or after detaching it. If the storage volume is attached to an instance, then only data that has already been written to the storage volume will be captured in the snapshot. Data that is cached by the application or the operating system will be excluded from the snapshot.<p>To create a snapshot of a bootable storage volume that is currently being used by an instance, you should delete the instance before you create the snapshot, to ensure the consistency of data. Deleting the instance doesn't delete any data on the bootable storage volume, because the data is stored on a persistent boot disk. When the instance is deleted, it ensures that all data is written to the storage volume and no further data can be written to the disk while taking the snapshot. Ensure that you stop only the instance orchestration and not the orchestration that creates storage volumes. That way, only your instance is deleted and re-created and storage volumes or other resources defined in other orchestrations are not deleted. You can create the instance again later on, after the snapshot is created.<p>To use this snapshot to create a storage volume, see <a class=\"xref\" href=\"op-storage-volume--post.html\">Create a Storage Volume</a>.<p><b>Prerequisite:</b> Before creating a remote storage snapshot, ensure that you have selected a replication policy for your Oracle Storage Cloud Service instance. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=CSSTO-GUID-5D53C11F-3D9E-43E4-8D1D-DDBB95DEC715\">Selecting a Replication Policy for Your Service Instance</a> in <em>Using Oracle Storage Cloud Service</em>.</p><p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**StorageSnapshotPostRequest**](StorageSnapshotPostRequest.md)|  | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**StorageSnapshotResponse**](StorageSnapshot-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteStorageSnapshot**
> DeleteStorageSnapshot($name, $cookie)

Delete a Storage Volume Snapshot

Deletes the specified storage volume snapshot. No response is returned.<p>Note that you can't delete a snapshot if it has been used to create a new storage volume.<p><b>Prerequisite:</b> Before deleting a remote storage snapshot, ensure that you have selected a replication policy for your Oracle Storage Cloud Service instance. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=CSSTO-GUID-5D53C11F-3D9E-43E4-8D1D-DDBB95DEC715\">Selecting a Replication Policy for Your Service Instance</a> in <em>Using Oracle Storage Cloud Service</em>.</p><p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Multipart name of the storage snapshot. | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverRootStorageSnapshot**
> StorageSnapshotDiscoverResponse DiscoverRootStorageSnapshot($cookie)

Retrieve Names of Containers

Retrieves the names of containers that contain objects that you can access. You can use this information to construct the multipart name of an object.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**StorageSnapshotDiscoverResponse**](StorageSnapshot-discover-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+directory+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverStorageSnapshot**
> StorageSnapshotDiscoverResponse DiscoverStorageSnapshot($container, $cookie)

Retrieve Names of all Storage Volume Snapshots and Subcontainers in a Container

Retrieves the names of objects and subcontainers that you can access in the specified container.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| Specify &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;i&gt;user&lt;/i&gt;/&lt;/code&gt; to retrieve the names of objects that you can access. Specify &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;/code&gt; to retrieve the names of containers that contain objects that you can access. | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**StorageSnapshotDiscoverResponse**](StorageSnapshot-discover-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+directory+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStorageSnapshot**
> StorageSnapshotResponse GetStorageSnapshot($name, $cookie)

Retrieve Details of a Storage Volume Snapshot

Retrieves details about the specified storage volume snapshot.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Multipart name of the snapshot of the storage volume. | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**StorageSnapshotResponse**](StorageSnapshot-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListStorageSnapshot**
> StorageSnapshotListResponse ListStorageSnapshot($container, $since, $until, $tags, $cookie)

Retrieve Details of all Storage Volume Snapshots in a Container

Retrieves details of the storage volume snapshots that are available in the specified container and match the specified query criteria. If you don't specify any query criteria, then details of all the storage volume snapshots in the container are displayed. To filter the search results, you can pass one or more of the following query parameters, by appending them to the URI in the following syntax:<p><code>?parameter1=value1&ampparameter2=value2&ampparameterN=valueN</code><p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| &lt;p&gt;&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;&lt;/code&gt; or &lt;p&gt;&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;&lt;/code&gt; | 
 **since** | **string**| Lists only those snapshots that were generated since the specified time or date. Snapshots that were generated before the specified time are not retrieved. The date you specify is checked against the &lt;code&gt;snapshot_timestamp&lt;/code&gt; field of the snapshots in the specified container.Example values that you can pass for this parameter are: &lt;code&gt;2016&lt;/code&gt;, &lt;code&gt;2016-06-17&lt;/code&gt;, or &lt;code&gt;2016-06-17T07:04:52Z &lt;/code&gt;.&lt;p&gt;To retrieve snapshots that were generated within a specific timeframe, specify values for both &lt;code&gt;since&lt;/code&gt; and &lt;code&gt;until&lt;/code&gt; parameters. | [optional] 
 **until** | **string**| Lists only those snapshots that were generated until the specified time or date. Snapshots that were generated after the specified time are not retrieved. The date you specify is checked against the &lt;code&gt;snapshot_timestamp&lt;/code&gt; field of the snapshots in the specified container. Example values that you can pass for this parameter are: &lt;code&gt;2016&lt;/code&gt;, &lt;code&gt;2016-06-17&lt;/code&gt;, or &lt;code&gt;2016-06-17T07:04:52Z &lt;/code&gt;.&lt;p&gt;To retrieve snapshots that were generated within a specific timeframe, specify values for both &lt;code&gt;since&lt;/code&gt; and &lt;code&gt;until&lt;/code&gt; parameters. | [optional] 
 **tags** | [**[]string**](string.md)| A comma-separated list of strings which helps you to identify storage snapshots. | [optional] 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**StorageSnapshotListResponse**](StorageSnapshot-list-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

