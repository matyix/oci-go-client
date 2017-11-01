# \MachineImagesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMachineImage**](MachineImagesApi.md#AddMachineImage) | **Post** /machineimage/ | Create a Machine Image
[**DeleteMachineImage**](MachineImagesApi.md#DeleteMachineImage) | **Delete** /machineimage/{name} | Delete a Machine Image
[**DiscoverMachineImage**](MachineImagesApi.md#DiscoverMachineImage) | **Get** /machineimage/{container} | Retrieve Names of all Machine Images and Subcontainers in a Container
[**DiscoverRootMachineImage**](MachineImagesApi.md#DiscoverRootMachineImage) | **Get** /machineimage/ | Retrieve Names of Containers
[**GetMachineImage**](MachineImagesApi.md#GetMachineImage) | **Get** /machineimage/{name} | Retrieve Details of a Machine Image
[**ListMachineImage**](MachineImagesApi.md#ListMachineImage) | **Get** /machineimage/{container}/ | Retrieve Details of all Machine Images in a Container


# **AddMachineImage**
> MachineImageResponse AddMachineImage($body, $cookie)

Create a Machine Image

Adds a machine image to Oracle Compute Cloud Service.<p>After adding a machine image, you can use the machine image to create a bootable storage volume. See <a class=\"xref\" href=\"op-storage-volume--post.html\">Create a Storage Volume</a>.<p>You can also add the machine image to an image list. See <a class=\"xref\" href=\"Adding a Machine Image.html\">Adding a Machine Image</a>.<p><b>Note:</b> Before performing this task, you must upload your machine image file to Oracle Storage Cloud Service. Make a note of the name of the <code>.tar.gz</code> file that you have uploaded. You'll need to provide this name while adding a machine image to Oracle Compute Cloud Service.<p>For more information, see <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=STCSG-GUID-799D6F6D-BDED-4DDE-9B3D-BE23BE5F687F\">Uploading Machine Image Files to Oracle Storage Cloud Service</a> in <em>Using Oracle Compute Cloud Service (IaaS)</em>.<p><b>Prerequisite:</b> Ensure that you have selected a replication policy for your Oracle Storage Cloud Service instance. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=CSSTO-GUID-5D53C11F-3D9E-43E4-8D1D-DDBB95DEC715\">Selecting a Replication Policy for Your Service Instance</a> in <em>Using Oracle Storage Cloud Service</em>.</p><p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**MachineImagePostRequest**](MachineImagePostRequest.md)|  | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**MachineImageResponse**](MachineImage-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMachineImage**
> DeleteMachineImage($name, $cookie)

Delete a Machine Image

Deletes the specified machine image. No response is returned. When you no longer need a machine image that you've registered in Oracle Compute Cloud Service, you can delete the image. <p>You can't delete system-provided machine images that are available in the <code>/oracle/public</code> container.<p><b>Caution:</b><p> If you delete a machine image that's used in an orchestration, then when that orchestration is stopped and re-started, the instances won't be created.<p>Deleting an image removes the image from your Oracle Compute Cloud Service account. However, the image file in your Oracle Storage Cloud Service account is <b>not</b> deleted. You can register this image in your Oracle Compute Cloud Service account again later, if required. You'll only need to remember the name of the <code>.tar.gz</code> image file that is available in your Oracle Storage Cloud Service account. For instructions to permanently remove a machine image file from Oracle Storage Cloud Service, see the <a target=\"_blank\" href=\"https://apexapps.oracle.com/pls/apex/f?p=44785:112:::::P112_CONTENT_ID:11959]\">Deleting Machine Image Files from Oracle Storage Cloud Service tutorial</a><p><b>Prerequisites:</b><ul><li>Ensure that the machine image you want to delete isn't used in any orchestration.</li><li>Ensure that you have selected a replication policy for your Oracle Storage Cloud Service instance. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=CSSTO-GUID-5D53C11F-3D9E-43E4-8D1D-DDBB95DEC715\">Selecting a Replication Policy for Your Service Instance</a> in <em>Using Oracle Storage Cloud Service</em>.</li></ul></p><p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| &lt;p&gt;The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverMachineImage**
> MachineImageDiscoverResponse DiscoverMachineImage($container, $cookie)

Retrieve Names of all Machine Images and Subcontainers in a Container

Retrieves the names of objects and subcontainers that you can access in the specified container.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| Specify &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;i&gt;user&lt;/i&gt;/&lt;/code&gt; to retrieve the names of objects that you can access. Specify &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;/code&gt; to retrieve the names of containers that contain objects that you can access. Specify &lt;code&gt;/oracle/public&lt;/code&gt; to retrieve the names of system-provided objects. | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**MachineImageDiscoverResponse**](MachineImage-discover-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+directory+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverRootMachineImage**
> MachineImageDiscoverResponse DiscoverRootMachineImage($cookie)

Retrieve Names of Containers

Retrieves the names of containers that contain objects that you can access. You can use this information to construct the multipart name of an object.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**MachineImageDiscoverResponse**](MachineImage-discover-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+directory+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMachineImage**
> MachineImageResponse GetMachineImage($name, $cookie)

Retrieve Details of a Machine Image

<b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| &lt;p&gt;The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt; or &lt;code&gt;/oracle/public/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;.) | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**MachineImageResponse**](MachineImage-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMachineImage**
> MachineImageListResponse ListMachineImage($container, $cookie)

Retrieve Details of all Machine Images in a Container

<b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| &lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;&lt;/code&gt; and &lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;/code&gt;for user-defined machine images and &lt;code&gt;/oracle/public&lt;/code&gt; for the Oracle-provided machine images. | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**MachineImageListResponse**](MachineImage-list-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

