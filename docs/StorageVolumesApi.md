# \StorageVolumesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddStorageVolume**](StorageVolumesApi.md#AddStorageVolume) | **Post** /storage/volume/ | Create a Storage Volume
[**DeleteStorageVolume**](StorageVolumesApi.md#DeleteStorageVolume) | **Delete** /storage/volume/{name} | Delete a Storage Volume
[**DiscoverRootStorageVolume**](StorageVolumesApi.md#DiscoverRootStorageVolume) | **Get** /storage/volume/ | Retrieve Names of Containers
[**DiscoverStorageVolume**](StorageVolumesApi.md#DiscoverStorageVolume) | **Get** /storage/volume/{container} | Retrieve Names of all Storage Volumes and Subcontainers in a Container
[**GetStorageVolume**](StorageVolumesApi.md#GetStorageVolume) | **Get** /storage/volume/{name} | Retrieve Details of a Storage Volume
[**ListStorageVolume**](StorageVolumesApi.md#ListStorageVolume) | **Get** /storage/volume/{container}/ | Retrieve Details of all Storage Volumes in a Container
[**UpdateStorageVolume**](StorageVolumesApi.md#UpdateStorageVolume) | **Put** /storage/volume/{name} | Update a Storage Volume


# **AddStorageVolume**
> StorageVolumeResponse AddStorageVolume($body, $cookie)

Create a Storage Volume

Creates a storage volume. You can create a storage volume in one of the following ways:<p>* Create a storage volume to store data and applications. While creating such a storage volume, you must specify values for the following parameters: <code>size</code>, <code>name</code>, and <code>properties</code>.<p>* Create a bootable storage volume by associating a storage volume with a machine image. While creating an instance, you can specify a bootable storage volume as a persistent boot disk for the instance. While creating a bootable storage volume, you must specify values for the following parameters: <code>bootable</code>, <code>name</code>, <code>imagelist</code>, <code>size</code>, and <code>properties</code>.<p>* Create a storage volume from a colocated snapshot of a storage volume. Colocated snapshots are stored in the same physical location as the original storage volume. While creating a storage volume from a colocated snapshot, you must specify values for the following parameters: <code>name</code>, multipart name of the <code>snapshot</code>, <code>size</code>, and <code>properties</code>(Properties name has to be same as the parent storage volume of the snapshot). If you are restoring a bootable storage volume from a snapshot, additionally you must specify the value for <code>bootable</code> as <code>true</code>.<p>* Create a storage volume from a remote snapshot of a storage volume. Remote snapshots aren't stored in the same location as the original storage volume. Instead, they are stored in the associated Oracle Storage Cloud Service instance. While creating a storage volume from a remote snapshot, you must specify values for the following parameters: <code>name</code>, <code>snapshot_id</code>, <code>size</code>, and <code>properties</code>(properties name can be different from the parent storage volume of the snapshot). If you are restoring a bootable storage volume from a snapshot, additionally you must specify the value for <code>bootable</code> as <code>true</code>.<p>After creating storage volumes you can attach them to instances by using the HTTP request, POST /storage/attachment/ <a class=\"xref\" href=\"op-storage-attachment--post.html\">(Create a Storage Attachment)</a>.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**StorageVolumePostRequest**](StorageVolumePostRequest.md)|  | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**StorageVolumeResponse**](StorageVolume-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteStorageVolume**
> DeleteStorageVolume($name, $cookie)

Delete a Storage Volume

Deletes the specified storage volume. No response is returned.<p>Ensure that there are no snapshots of the storage volume that you want to delete. See <a class=\"xref\" href=\"op-storage-snapshot-{container}--get.html\">Retrieve Details of all Storage Volume Snapshots in a Container</a>.


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

# **DiscoverRootStorageVolume**
> StorageVolumeDiscoverResponse DiscoverRootStorageVolume($cookie)

Retrieve Names of Containers

Retrieves the names of containers that contain objects that you can access. You can use this information to construct the multipart name of an object.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**StorageVolumeDiscoverResponse**](StorageVolume-discover-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+directory+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverStorageVolume**
> StorageVolumeDiscoverResponse DiscoverStorageVolume($container, $cookie)

Retrieve Names of all Storage Volumes and Subcontainers in a Container

Retrieves the names of objects and subcontainers that you can access in the specified container.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| Specify &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;i&gt;user&lt;/i&gt;/&lt;/code&gt; to retrieve the names of objects that you can access. Specify &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;/code&gt; to retrieve the names of containers that contain objects that you can access. | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**StorageVolumeDiscoverResponse**](StorageVolume-discover-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+directory+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStorageVolume**
> StorageVolumeResponse GetStorageVolume($name, $cookie)

Retrieve Details of a Storage Volume

Retrieves details about the specified storage volume.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**StorageVolumeResponse**](StorageVolume-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListStorageVolume**
> StorageVolumeListResponse ListStorageVolume($container, $name, $tags, $cookie)

Retrieve Details of all Storage Volumes in a Container

Retrieves details of the storage volumes that are available in the specified container and match the specified query criteria. If you don't specify any query criteria, then details of all the storage volumes in the container are displayed. To filter the search results, you can pass one or more of the following query parameters, by appending them to the URI in the following syntax:<p><code>?parameter1=value1&ampparameter2=value2&ampparameterN=valueN</code><p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| &lt;p&gt;&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;&lt;/code&gt; or &lt;p&gt;&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;&lt;/code&gt; | 
 **name** | **string**| The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | [optional] 
 **tags** | [**[]string**](string.md)| Comma-separated strings that tag the storage volume. | [optional] 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**StorageVolumeListResponse**](StorageVolume-list-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateStorageVolume**
> StorageVolumeResponse UpdateStorageVolume($name, $body, $cookie)

Update a Storage Volume

Updates a storage volume.<p>Although you have to pass values for several parameters, you can only increase the size of the storage volume and modify the values for the tags and description parameters. You can update an existing storage volume to increase the capacity dynamically, even when the volume is attached to an instance. You must specify all the required fields, although these fields won't be updated.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | 
 **body** | [**StorageVolumePutRequest**](StorageVolumePutRequest.md)|  | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**StorageVolumeResponse**](StorageVolume-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

