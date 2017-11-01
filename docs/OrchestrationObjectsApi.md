# \OrchestrationObjectsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOrchestrationObject**](OrchestrationObjectsApi.md#AddOrchestrationObject) | **Post** /platform/v1/object/ | Create an Orchestration Object
[**DeleteOrchestrationObject**](OrchestrationObjectsApi.md#DeleteOrchestrationObject) | **Delete** /platform/v1/object/{name} | Delete an Orchestration Object
[**DiscoverOrchestrationObject**](OrchestrationObjectsApi.md#DiscoverOrchestrationObject) | **Get** /platform/v1/object/{container} | Retrieve Names of all Orchestration Objects and Subcontainers in a Container
[**DiscoverRootOrchestrationObject**](OrchestrationObjectsApi.md#DiscoverRootOrchestrationObject) | **Get** /platform/v1/object/ | Retrieve Names of Containers
[**GetOrchestrationObject**](OrchestrationObjectsApi.md#GetOrchestrationObject) | **Get** /platform/v1/object/{name} | Retrieve Details of an Orchestration Object
[**ListOrchestrationObject**](OrchestrationObjectsApi.md#ListOrchestrationObject) | **Get** /platform/v1/object/{container}/ | Retrieve Details of all Orchestration Objects in a Container
[**UpdateOrchestrationObject**](OrchestrationObjectsApi.md#UpdateOrchestrationObject) | **Put** /platform/v1/object/{name} | Update an Orchestration Object


# **AddOrchestrationObject**
> OrchestrationObjectResponse AddOrchestrationObject($body, $cookie)

Create an Orchestration Object

 Adds an object to the specified orchestration v2.<p>The action of the orchestration object you add is determined by the <code>desired_state</code> of the orchestration it is associated with.<p>* Objects are created when the <code>desired_state</code> of the associated orchestration is set to <code>active</code>.<p>* Objects are deleted when the <code>desired_state</code> of the associated orchestration is set to <code>inactive</code>.<p>* Only non-persistent objects are deleted when the <code>desired_state</code> of the associated orchestration is set to <code>suspend</code>. Peristent objects are not deleted when the associated orchestration is suspended. To make an object persistent, set the persistent attribute to true. When an object is set to persist, it is not deleted when the orchestration is suspended. If the persistent objects are not already in the <code>active</code> state, they are created. If the orchestration is terminated, then all the objects are deleted. If you set the persistent attribute of an object to true, then you must set the persistent attribute of all the dependent objects as well to true. For example, if a persistent instance references a bootable storage volume, the storage volume must also be persistent. For more information, see <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=STCSG-GUID-4F84DE2B-BDC9-4B9B-9B5D-B8129E0F7512\">Object Persistence in Orchestrations v2</a> in <i>Using Oracle Compute Cloud Service (IaaS)</i>.<p>To change the <code>desired_state</code> of an orchestration, send the PUT /platform/v1/orchestration/{name}/?desired_state={state} HTTP request.<p>You can determine the status of an object by viewing its <code>health</code> parameter.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrchestrationObjectPostRequest**](OrchestrationObjectPostRequest.md)|  | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**OrchestrationObjectResponse**](OrchestrationObject-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrchestrationObject**
> DeleteOrchestrationObject($name, $terminate, $cookie)

Delete an Orchestration Object

Delete the specified orchestration object and the underlying object. For example, if you send a request to delete the <code>/Compute-acme/jack.jones@example.com/myOrchestration/mySeclist</code> orchestration object, the underlying <code>/Compute-acme/jack.jones@example.com/mySeclist</code> security list is also deleted.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | 
 **terminate** | **string**| Set terminate &lt;code&gt;true&lt;/code&gt; to delete active orchestration objects. | [optional] 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverOrchestrationObject**
> OrchestrationObjectDiscoverResponse DiscoverOrchestrationObject($container, $cookie)

Retrieve Names of all Orchestration Objects and Subcontainers in a Container

Retrieves the names of orchestration objects and subcontainers that you can access in the specified container. <p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| Specify the three-part name of the orchestration (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;orchestration&lt;/em&gt;&lt;/code&gt;) to retrieve names of all the objects in that orchestration. Specify &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;/code&gt; to retrieve the names of containers that contain objects that you can access. | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**OrchestrationObjectDiscoverResponse**](OrchestrationObject-discover-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+directory+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverRootOrchestrationObject**
> OrchestrationObjectDiscoverResponse DiscoverRootOrchestrationObject($cookie)

Retrieve Names of Containers

Retrieves the names of containers that contain objects that you can access. You can use this information to construct the multipart name of an object.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**OrchestrationObjectDiscoverResponse**](OrchestrationObject-discover-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+directory+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrchestrationObject**
> OrchestrationObjectResponse GetOrchestrationObject($name, $cookie)

Retrieve Details of an Orchestration Object

Retrieves details of the specified orchestration object. <p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The four-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;orchestration&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**OrchestrationObjectResponse**](OrchestrationObject-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrchestrationObject**
> OrchestrationObjectListResponse ListOrchestrationObject($container, $orchestration, $cookie)

Retrieve Details of all Orchestration Objects in a Container

Retrieves details of all the orchestration objects that are available and match the specified query criteria. If you don't specify any query criteria, then details of all the orchestration objects in the container are displayed. To filter the search results, you can pass one or more of the documented query parameters by appending them to the URI in the following syntax: <p><code>?parameter1=value1&ampparameter2=value2&ampparameterN=valueN</code><p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| Specify &lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;&lt;/code&gt; or &lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;&lt;/code&gt;.&lt;p&gt;Specify the three-part name of the orchestration (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;orchestration&lt;/em&gt;&lt;/code&gt;) to retrieve details of all the objects in that orchestration. | 
 **orchestration** | **string**| Return a list of orchestration objects filtered by the orchestration to which this object belongs | [optional] 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**OrchestrationObjectListResponse**](OrchestrationObject-list-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrchestrationObject**
> OrchestrationObjectResponse UpdateOrchestrationObject($name, $body, $cookie)

Update an Orchestration Object

Updates the specified orchestration object.<p>* You can update the following fields of an orchestration object: <code>description</code>, <code>label</code>, <code>persistent</code>, <code>relationships</code>, and <code>template</code>.<p>* You cannot update an object while the orchestration within which it is contained is in a transient state, such as <code>activating</code>, <code>deactivating</code>, and  <code>suspending</code>.<p>* You cannot move an object from one orchestration to another.<p>To update an orchestration object:<p>1.Get the orchestration object by sending the <a class=\"xref\" href=\"op-platform-v1-object-{name}-get.html\">GET /platform/v1/object/{name}</a> request. The response body is in the JSON format. <p>2.Modify the JSON file as per your requirements. <p>3.Provide the modified JSON as the request body for the <a class=\"xref\" href=\"op-platform-v1-object-{name}-put.html\">PUT /platform/v1/object/{name}</a> request. Ensure that you specify the latest <code>version</code> of the orchestration object in the request body. <p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The four-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | 
 **body** | [**OrchestrationObjectPutRequest**](OrchestrationObjectPutRequest.md)|  | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**OrchestrationObjectResponse**](OrchestrationObject-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

