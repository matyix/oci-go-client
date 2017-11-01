# \OrchestrationV2sApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOrchestrationV2**](OrchestrationV2sApi.md#AddOrchestrationV2) | **Post** /platform/v1/orchestration/ | Create an OrchestrationV2
[**DeleteOrchestrationV2**](OrchestrationV2sApi.md#DeleteOrchestrationV2) | **Delete** /platform/v1/orchestration/{name} | Delete an OrchestrationV2
[**DiscoverOrchestrationV2**](OrchestrationV2sApi.md#DiscoverOrchestrationV2) | **Get** /platform/v1/orchestration/{container} | Retrieve Names of all OrchestrationV2 and Subcontainers in a Container
[**DiscoverRootOrchestrationV2**](OrchestrationV2sApi.md#DiscoverRootOrchestrationV2) | **Get** /platform/v1/orchestration/ | Retrieve Names of Containers
[**GetOrchestrationV2**](OrchestrationV2sApi.md#GetOrchestrationV2) | **Get** /platform/v1/orchestration/{name} | Retrieve Details of an OrchestrationV2
[**ListOrchestrationV2**](OrchestrationV2sApi.md#ListOrchestrationV2) | **Get** /platform/v1/orchestration/{container}/ | Retrieve Details of all OrchestrationV2 in a Container
[**UpdateOrchestrationV2**](OrchestrationV2sApi.md#UpdateOrchestrationV2) | **Put** /platform/v1/orchestration/{name} | Update an OrchestrationV2


# **AddOrchestrationV2**
> OrchestrationV2Response AddOrchestrationV2($body, $cookie)

Create an OrchestrationV2

To use an orchestration to control the provisioning and life cycle of objects in Oracle Compute Cloud Service, you must define the orchestration in a JSON-format file and then use this request to add an orchestration to Oracle Compute Cloud Service.<p><b>Prerequisites</b>:<p>* You must have already created the orchestration file that you want to upload. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=STCSG-GUID-130B0A9A-5B21-4BB4-8F4C-FC5FFD0C7586\">Building Your First Orchestration v2</a>.<p>* You should also validate your JSON file. You can do this by using a third-party tool, such as <a target=\"_blank\" href=\"http://jsonlint.com/\">JSONLint</a>, or any other validation tool of your choice. If your JSON isn't valid, then an error occurs when you upload the orchestration. Oracle doesn't support or endorse any third-party validation tool.<p> If you upload an orchestrations v2 file with the <code>desired_state</code> specified as <code>active</code>, the orchestration is activated automatically and the objects defined in it are created.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrchestrationV2PostRequest**](OrchestrationV2PostRequest.md)|  | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**OrchestrationV2Response**](OrchestrationV2-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrchestrationV2**
> DeleteOrchestrationV2($name, $terminate, $cookie)

Delete an OrchestrationV2

After you've terminated an orchestration, if you don't need it any more, you can delete the orchestration. If you want to retain the orchestration but delete the objects created by it, you can terminate the orchestration. See <a class=\"xref\" href=\"op-platform-v1-orchestration-{name}-put.html\">Update an OrchestrationV2</a>.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | 
 **terminate** | **string**| Set terminate to &lt;code&gt;true&lt;/code&gt; to delete an active orchestration. | [optional] 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverOrchestrationV2**
> OrchestrationV2DiscoverResponse DiscoverOrchestrationV2($container, $cookie)

Retrieve Names of all OrchestrationV2 and Subcontainers in a Container

Retrieves the names of orchestrationV2s and subcontainers that you can access in the specified container. <p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| Specify &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;i&gt;user&lt;/i&gt;/&lt;/code&gt; to retrieve the names of objects that you can access. Specify &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;/code&gt; to retrieve the names of containers that contain objects that you can access. | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**OrchestrationV2DiscoverResponse**](OrchestrationV2-discover-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+directory+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverRootOrchestrationV2**
> OrchestrationV2DiscoverResponse DiscoverRootOrchestrationV2($cookie)

Retrieve Names of Containers

Retrieves the names of containers that contain objects that you can access. You can use this information to construct the multipart name of an object.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**OrchestrationV2DiscoverResponse**](OrchestrationV2-discover-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+directory+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrchestrationV2**
> OrchestrationV2Response GetOrchestrationV2($name, $cookie)

Retrieve Details of an OrchestrationV2

Retrieves details of the specified orchestrationV2. <p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**OrchestrationV2Response**](OrchestrationV2-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrchestrationV2**
> OrchestrationV2ListResponse ListOrchestrationV2($container, $tags, $cookie)

Retrieve Details of all OrchestrationV2 in a Container

Retrieves details of the orchestrations that are available in the specified container and match the specified query criteria. If you don't specify any query criteria, then details of all the orchestrations in the container are displayed. To filter the search results, you can pass one or more tags by appending them to the URI in the following syntax:<p><code>?parameter1=value1&ampparameter2=value2&ampparameterN=valueN</code> <p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| &lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;&lt;/code&gt; or &lt;p&gt;&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;&lt;/code&gt; | 
 **tags** | [**[]string**](string.md)| Strings that describe the orchestration and help you identify it. | [optional] 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**OrchestrationV2ListResponse**](OrchestrationV2-list-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrchestrationV2**
> OrchestrationV2Response UpdateOrchestrationV2($name, $body, $desiredState, $cookie)

Update an OrchestrationV2

Orchestration cannot be updated when they are in the activating, suspending, or deactivating state. However, you can terminate or delete an orchestration irrespective of the current state of the orchestration.<p>* To activate an orchestration, use the <code>PUT /platform/v1/orchestration/orchestrationName?desired_state=active</code> method.<p>* To delete all the nonpersistent objects defined in the orchestration, use the <code>PUT /platform/v1/orchestration/orchestrationName?desired_state=suspend</code> method. When you suspend an active orchestration, only the nonpersistent objects are deleted; the persistent objects are not deleted. <p>* To delete all the objects in an orchestration, use the <code>PUT /platform/v1/orchestration/orchestrationName?desired_state=inactive</code> method.<p><b>Note:</b> Instead of using <code>PUT /platform/v1/orchestration/</code> to add, remove, or modify individual objects, use the <a class=\"xref\" href=\"api-OrchestrationObjects.html\">Object API</a>.<p>* To add an object to the orchestration, send the <a class=\"xref\" href=\"op-platform-v1-object--post.html\">POST /platform/v1/object/</a> request.<p>* To modify an existing object in the orchestration, send the <a class=\"xref\" href=\"op-platform-v1-object-{name}-put.html\">PUT /platform/v1/object/{name}</a> request.<p>* To delete an object from the orchestration, send the <a class=\"xref\" href=\"op-platform-v1-object-{name}-delete.html\">DELETE /platform/v1/object/{name}</a> request.<p>If you want to update an orchestration, get the latest version of the orchestration by sending the <a class=\"xref\" href=\"op-platform-v1-orchestration-{name}-get.html\">GET /platform/v1/orchestration/{name}</a> request. Modify the orchestration as per your requirements, and then use the <a class=\"xref\" href=\"op-platform-v1-orchestration-{name}-put.html\">PUT /platform/v1/orchestration/{name}</a> request to update the orchestration. Ensure that you specify the latest <code>version</code> of the orchestration in the request body.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | 
 **body** | [**OrchestrationV2PutRequest**](OrchestrationV2PutRequest.md)|  | 
 **desiredState** | **string**| Manage the state of the orchestration objects by changing the desired state of the orchestration. | [optional] 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**OrchestrationV2Response**](OrchestrationV2-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

